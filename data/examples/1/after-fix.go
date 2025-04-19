type ActionScanningGC[ROW any] struct {
	spConfig *scConfig[ROW]
	shardCtrl     sharding.Controller
	config        ASGCConfig
	sRouter spannercr.sRouter
	scope         tally.Scope
	logger        *zap.Logger
	client   fl.client

	lockMap         sync.Map
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
		lockMap:         sync.Map{},
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

	t.lockMap.Range(func(key, value interface{}) bool {
		shardKey := key.(sharding.ShardKey)
		if _, ok := newShards[shardKey]; !ok {
			t.lockMap.Delete(shardKey)
		}
		return true
	})

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
			val, _ := t.lockMap.LoadOrStore(shardKey, shardLock{
				m:       &sync.Mutex{},
				running: false,
			})
			shardLock := val.(shardLock)
			go t.runGC(t.ctx, shardKey, shardLock)
		case <-t.isStopped:
			return
		}
	}
}

func (t *ActionScanningGC[ROW]) stop(ctx context.Context) error {
	close(t.isStopped)
	<-t.hasStopped

	t.lockMap.Range(func(key, value interface{}) bool {
		t.lockMap.Delete(key)
		return true
	})

	t.cancelFunc()
	t.wg.Wait()
	return ctx.Err()
}