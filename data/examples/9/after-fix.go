func TestCollectAll(t *testing.T) {
	scenarioSet := []struct {
		name      string
		calls     [][]*bool
		batchSize int
		maxWait   time.Duration
		expectErr []error
	}{
		{
			name: "Send 1 batch of 1 and wait for success response",
			calls: [][]*bool{
				{&dataVal},
			},
			batchSize: 3,
			maxWait:   100 * time.Millisecond,
			expectErr: []error{nil},
		},
		{
			name: "Send 2 batches of 3 with differing responses",
			calls: [][]*bool{
				{&dataVal, &dataVal, &dataVal},
				{&dataVal, nil, &dataVal},
			},
			batchSize: 3,
			maxWait:   5 * time.Second,
			expectErr: []error{nil, errTrial},
		},
		{
			name: "Send 4 entries with 4 batch size and confirm they are batched together",
			calls: [][]*bool{
				{nil},
				{&dataVal},
				{&dataVal},
				{&dataVal},
			},
			batchSize: 4,
			maxWait:   5 * time.Second,
			expectErr: []error{errTrial, errTrial, errTrial, errTrial},
		},
		{
			name: "send unfilled batch because incoming request overfills buffer",
			calls: [][]*bool{
				{&dataVal},
				{nil, &dataVal, &dataVal},
			},
			batchSize: 3,
			maxWait:   500 * time.Millisecond,
			expectErr: []error{nil, errTrial},
		},
		{
			name: "send nil and empty calls, these will not error",
			calls: [][]*bool{
				nil,
				{},
				{nil},
			},
			batchSize: 3,
			maxWait:   100 * time.Millisecond,
			expectErr: []error{nil, nil, errTrial},
		},
		{
			name: "send batch size which exceeds limit",
			calls: [][]*bool{
				{&dataVal, &dataVal, &dataVal},
			},
			batchSize: 2,
			maxWait:   5 * time.Second,
			expectErr: []error{fmt.Errorf(_exceedErrFmt, 3, 2)},
		},
	}

	for _, tt := range scenarioSet {
		t.Run(tt.name, func(t *testing.T) {
			var wg sync.WaitGroup
			aggregator := NewScheduler[bool](
				tally.NoopScope,
				zap.NewNop(),
				tt.batchSize,
				tt.maxWait,
				boolProcessFunc,
			)

			readyChan := make(chan struct{})
			go func() {
				aggregator.Start(context.Background())
				close(readyChan)
			}()
			<-readyChan

			wg.Add(len(tt.calls))
			for i := 0; i < len(tt.calls); i++ {
				go func(idx int) {
					defer wg.Done()
					err := aggregator.CollectAll(tt.calls[idx])
					assert.Equal(t, tt.expectErr[idx], err)
				}(i)
			}
			wg.Wait()
			aggregator.Stop(context.Background())
		})
	}

	for _, tt := range scenarioSet {
		t.Run(tt.name, func(t *testing.T) {
			var wg sync.WaitGroup
			aggregator := NewScheduler[bool](
				tally.NoopScope,
				zap.NewNop(),
				tt.batchSize,
				tt.maxWait,
				boolProcessFunc,
			)

			go func() {
				aggregator.Start(context.Background())
			}()

			wg.Add(2 * len(tt.calls))
			for i := 0; i < len(tt.calls); i++ {
				errChan := make(chan error)

				go func(idx int) {
					defer wg.Done()
					aggregator.CollectAllAsync(tt.calls[idx], errChan)
				}(i)

				go func(idx int) {
					defer wg.Done()
					defer close(errChan)
					assert.Equal(t, tt.expectErr[idx], <-errChan)
				}(i)
			}

			wg.Wait()
			aggregator.Stop(context.Background())
		})
	}
}
