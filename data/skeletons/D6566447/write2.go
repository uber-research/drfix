package skeleton

func (v0 *type0) func1(v3 pkg9.v6, v8 *pkg16.v7) (*pkg15.v4, type1) {
	v9 := new(sync.WaitGroup)
	var racyVar0 type1
	go func() {
		defer v9.Done()
		v10, racyVar0 = v0.v11.Func4(v3, v8)
		if v1 != nil {
			v12 <- v1
		}
	}()
	go func() {
		defer v9.Done()
		v5, racyVar0 = v0.v11.Func11(v3, v8)
		if v1 != nil {
			v12 <- v1
		}
	}()
	Wrapper1.Wait()
	if len(v12) > IntConst0 {
		select {
		case v1 := <-v12:
			if v1 != nil {
				return nil, v1
			}
		}
	}
	return v2.Func21(v10, int64(pkg2.Func22(float64(v5)/float64(v8.v13))))
}