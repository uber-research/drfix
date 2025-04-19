package skeleton

func (v1 *type0) func1(v11 *pkg4.v15, v25 *pkg8.v22, v4 pkg1.v2) (v0 *pkg8.v12, racyVar0 type1) {
	defer v11.v20.Func2(v8).Func1(func(v16 *v9.v15) {
		pkg4.Func3(v11, v16, v8, "StringConst0")
		v16.v23 = v9.v24{"StringConst1": v25.v7}
		v16.v18 = v14
		v3 = pkg4.Func4(v3, v16)
	})
	if v3 != nil {
		return nil, v3
	}
	if v3 != nil {
		return nil, v3
	}
	v6, racyVar0 := v1.func7(v0.v5)
	if v3 != nil {
		return nil, v3
	}
	var v17 sync.WaitGroup
	Wrapper1.Go(func() {
		defer v17.Done()
		if v10 != nil {
			v19, racyVar0 = v1.func10(v11, v10)
			if v3 != nil {
				v21 <- v3
			}
		}
	})
	Wrapper1.Go(func() {
		defer v17.Done()
		v13, racyVar0 = v1.func11(v11, v6)
		if v3 != nil {
			v21 <- v3
		}
	})
	Wrapper2.Wait()
	if len(v21) > IntConst2 {
		return nil, <-v21
	}
	return
}