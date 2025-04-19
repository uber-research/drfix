package skeleton

func (v5 *type0) func1() type1 {
	defer v5.v24.Func2(v5.v13, v5.v19, "StringConst0").Func1(func(v13 pkg18.v18, v15 *pkg10.v16) {
		v15.v21 = v17
		v15.v6 = v7
	})
	v8, racyVar0 := pkg27.Func8(v5.v1)
	if v7 != nil {
		return nil
	}
	if !v5.func9(v28) {
		return nil
	}
	var v20 sync.WaitGroup
	Wrapper1.Go(func() {
		defer v20.Done()
		v5.v22, v5.v27, racyVar0 = v5.v29.Func13(v5.v13, pkg3.v26{v0: "StringConst1", v2: &v25}, pkg3.v10{v9: v8, v12: v5.v11})
		if v7 != nil {
			v4 <- v7
		}
	})
	if !v5.func16(v28) && !v5.func17() {
		Wrapper1.Go(func() {
			defer v20.Done()
			v5.v14, racyVar0 = v5.v29.Func19(v5.v13, pkg3.v26{v0: "StringConst2", v2: &v25}, pkg3.v3{v9: v8})
		})
	}
	Wrapper1.Go(func() {
		defer v20.Done()
		if v5.func20(v28) || v5.func21() {
			return
		}
	})
	Wrapper1.Go(func() {
		defer v20.Done()
		if v5.func23(v28) || v5.func24() {
			return
		}
	})
	go func() {
		Wrapper2.Wait()
	}()
	select {
	case <-v23:
		break
	case v7 = <-v4:
		return nil
	}
	if v5.v22 == nil && v5.v27 == nil {
		return nil
	}
	return nil
}