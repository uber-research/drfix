func FetchAllUsers(
	ctx context.Context,
	userCountPerRequest int,
	startEmpNo int,
	endEmpNo int,
	batchCount int,
) ([]entity.User, error) {
	logger := activity.GetLogger(ctx)
	users := make([]entity.User, 0)
	sendRequestChan := make(chan struct{})
	stopRequestChan := make(chan struct{})
	doneChan := make(chan struct{})
	errChan := make(chan error)

	responseChan := make(chan []entity.User)
	minBatchResponseLength := int32(userCountPerRequest)
	wg := new(sync.WaitGroup)
	waitRWMutex := sync.RWMutex{}

	var err error

	go func() {
		for {
			select {
			case err = <-errChan:
				logger.Error("fetching employees errored out",
					zap.Error(err),
					zap.Any("startEmpNo", startEmpNo),
					zap.Any("endEmpNo", startEmpNo+userCountPerRequest-1),
				)
			case <-stopRequestChan:
				close(sendRequestChan)
				close(responseChan)
				close(doneChan)
				return
			case response := <-responseChan:
				if startEmpNo > endEmpNo {
					atomic.StoreInt32(&minBatchResponseLength, 0)
				}
				users = append(users, response...)
			case <-sendRequestChan:
				for i := 0; i < batchCount; i++ {
					waitRWMutex.Lock()
					wg.Add(1)
					waitRWMutex.Unlock()
					go getUsersFromWorkday(
						ctx, wg, responseChan, errChan, startEmpNo, startEmpNo+userCountPerRequest-1, endEmpNo)
					startEmpNo += userCountPerRequest
				}
			}
		}
	}()

	go func() {
		for {
			if atomic.LoadInt32(&minBatchResponseLength) == 0 {
				close(stopRequestChan)
				return
			}
			sendRequestChan <- struct{}{}
			waitRWMutex.RLock()
			wg.Wait()
			waitRWMutex.RUnlock()
		}
	}()
	<-doneChan
	defer close(errChan)
	if err != nil {
		return users, err
	}

	return users, nil
}
