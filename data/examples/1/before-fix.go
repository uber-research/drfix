type ActionScanningGC[ROW any] struct {
	spConfig *scConfig[ROW]
	shardCtrl     sharding.Controller
	config        ASGCConfig
	sRouter spannercr.sRouter
	scope         tally.Scope
	logger        *zap.Logger
	client   fl.client

	lockMap         map[sharding.ShardKey]shardLock
	tickDuration    time.Duration
	groupsGCBuckets tally.ValueBuckets
	isStopped       chan struct{}
	hasStopped      chan struct{}
	ctx             context.Context
	cancelFunc      context.CancelFunc
	wg              sync.WaitGroup
}

func newActionScanningGC[ROW any](
	spConfig *scConfig[ROW],
	shardCtrl sharding.Controller,
	config ASGCConfig,
	sRouter spannercr.sRouter,
	scope tally.Scope,
	logger *zap.Logger,
	client fl.client,
) scanningGC {
	ctx, cancelFunc := context.WithCancel(context.Background())
	return &ActionScanningGC[ROW]{
		spConfig:   spConfig,
		shardCtrl:       shardCtrl,
		config:          config,
		sRouter:   sRouter,
		scope:           scope,
		logger:          logger,
		client:     client,
		lockMap:         make(map[sharding.ShardKey]shardLock),
		tickDuration:    config.ScheduleTick,
		groupsGCBuckets: tally.MustMakeLinearValueBuckets(0, float64(config.DeleteBatch/16), 16),
		isStopped:       make(chan struct{}),
		hasStopped:      make(chan struct{}),
		ctx:             ctx,
		cancelFunc:      cancelFunc,
	}
}

func (t *ActionScanningGC[ROW]) runShards() {
	newShards := t.shardCtrl.MyShards()

	for removedShard := range t.lockMap {
		if _, ok := newShards[removedShard]; !ok {
			delete(t.lockMap, removedShard)
		}
	}

	enabled, _ := t.client.GetBoolValue(
		t.ctx, "enabled", fl.Constraints{fl.Constraints2(map[string]interface{}{})}, true)
	if !enabled {
		return
	}

	var timer *time.Ticker
	if len(newShards) > 0 {
		timer = time.NewTicker(t.tickDuration / time.Duration(len(newShards)))
	}

	for shardKey := range newShards {
		select {
		case <-timer.C:
			if _, ok := t.lockMap[shardKey]; !ok {
				t.lockMap[shardKey] = shardLock{
					m:       &sync.Mutex{},
					running: false,
				}
			}
			shardLock := t.lockMap[shardKey]
			go t.runGC(t.ctx, shardKey, shardLock)
		case <-t.isStopped:
			return
		}
	}
}

func (t *ActionScanningGC[ROW]) stop(ctx context.Context) error {
	close(t.isStopped)
	<-t.hasStopped

	for k := range t.lockMap {
		delete(t.lockMap, k)
	}

	t.cancelFunc()
	t.wg.Wait()
	return ctx.Err()
}