func getItemBundles(ctx context.Context, dataClient servicepb.DataAPIYARPCClient, logger common.BasicLogger, itemUUIDs []uuid.UUID) ([]*itemBundle, error) {
	bundles := make([]*itemBundle, 0, 0)

	workQueue := async.NewWorkQueue[*itemBundle](_flagsVar.Concurrency)
	group, _ := safegroup.WithContext(ctx)

	group.Go(func() error {
		defer workQueue.Close()
		for {
			result, err := workQueue.Next(ctx)
			if err == io.EOF {
				return nil
			}
			if err != nil {
				return err
			}

			err = result.Error()
			if err != nil {
				logger.Error(err.Error())
				bundles = append(bundles, nil)
				continue
			}
			bundle := result.Value()
			bundles = append(bundles, bundle)
		}
	})

	for _, itemUUID := range itemUUIDs {
		itemUUID := itemUUID
		err := workQueue.Enqueue(ctx, func() (*itemBundle, error) {
			bundle, err := getItemBundle(ctx, dataClient, logger, itemUUID)
			if err != nil {
				return nil, someErrs.WrapMessage(err, "unable to get item bundles")
			}
			return bundle, nil
		})
		if err != nil {
			code := someErrs.Code(err)
			if errors.Is(err, async.ErrChClosed) {
				code = someErrs.CodeInternal
			}
			return nil, someErrs.Wrap(err, "failed to enqueue delete entity request", code)
		}
		time.Sleep(time.Duration(_flagsVar.Delay) * time.Millisecond)
	}

	workQueue.Close()
	err := group.Wait()
	if err != nil {
		return nil, someErrs.WrapMessage(err, "failed waiting for the group to finish")
	}

	return bundles, nil
}
