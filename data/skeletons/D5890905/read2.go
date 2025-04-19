package skeleton

func (v4 *type0) func1(v9 pkg19.v12) (v6 []int32, v5 type1) {
	if v5 != nil {
		return nil, v5
	}
	if len(v10) == IntConst0 {
		return nil, v5
	}
	racyVar0 := int32(IntConst5)
	go func() {
		v15 := &sync.WaitGroup{}
		semaphore := semaphore.NewSemaphore(v4.v13.v14)
		for func18(v7) > IntConst7 {
			go func(v2 []int32) {
				defer v15.Done()
				semaphore.Increment()
				defer semaphore.Decrement()
				for !v18 && len(v2) > IntConst10 && v16 < v0 {
					select {
					case <-v9.Done():
					default:
						atomic.AddInt32(&racyVar0, int32(len(v2)-len(v8)))
						pkg5.Func25(pkg5.v20{"StringConst9": pkg16.Func28(v3).Func27(IntConst11 * pkg16.v17).Func26(), "StringConst10": v16, "StringConst11": racyVar0, "StringConst12": v8}).Func24("StringConst13")
						atomic.AddInt32(&v1, int32(len(v8)))
					}
				}
				v19 <- v2
			}(v11)
		}
		Wrapper1.Wait()
	}()
	if len(v8) > IntConst14 {
		return v8, pkg15.Func43("StringConst18", v8)
	}
	return nil, nil
}