package skeleton

func (v1 *v11) func1(v3 pkg6.v4, v8 *v9, v0 *v6) type0 {
	var racyVar0 type0
	var v5 sync.WaitGroup
	go func() {
		defer v5.Done()
	}()
	go func() {
		defer v5.Done()
		racyVar0 = v1.func5(v3, v8.v12, v8.v7.v10)
	}()
	go func() {
		defer v5.Done()
		racyVar0 = v1.func7(v3, v8.v12, v8.v7.v10)
	}()
	go func() {
		defer v5.Done()
	}()
	go func() {
		defer v5.Done()
	}()
	go func() {
		defer v5.Done()
	}()
	go func() {
		defer v5.Done()
	}()
	Wrapper1.Wait()
	return v2
}