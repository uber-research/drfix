package skeleton

func (v2 v32) func1(v8 pkg19.v19, v21 *pkg23.v20) (v0 *v26.v10, racyVar0 type0) {
	defer v2.v27.Func2(v8, v2.v23, v12).Func1(func(v8 pkg19.v19, v11 *pkg10.v16) {
		v11.v30 = pkg10.v34{"StringConst0": v21.v4.Func4().Func3(), "StringConst1": v21.v4.Func7().Func6().Func5(), "StringConst2": v21.v4.Func8(), "StringConst3": v21.v4.Func9(), "StringConst4": v21.v4.Func10(), "StringConst5": v21.v4.Func11(), "StringConst6": v21.v29.Func12()}
		v11.v6 = v7
		v11.v25 = v17
	})
	if !v2.func13(v8, v21.v4) {
		return v15, nil
	}
	v5, v1 := errgroup.WithContext(v8)
	Wrapper1.Go(func() type0 {
		v24, racyVar0 = v2.v37.Func14(v1, v3)
		return v7
	})
	Wrapper1.Go(func() type0 {
		v36, racyVar0 = v2.v13.Func15(v8, v21.v29.Func16())
		return v7
	})
	Wrapper1.Go(func() type0 {
		return nil
	})
	if v7 := Wrapper1.Wait(); v7 != nil {
		return v15, v7
	}
	if !v9 {
		return v15, nil
	}
	if len(v24.v35) > IntConst0 {
		return v15, nil
	}
	if len(v24.v14) == IntConst1 {
		return v15, nil
	}
	if v7 != nil {
		return v15, v7
	}
	if len(v28) == IntConst3 {
		return v15, v7
	}
	if v7 != nil || v22 {
		return v15, v7
	}
	if v7 != nil {
		return v15, v7
	}
	if !v31 {
		return v15, nil
	}
	v33, v38 := v2.func46(v8, v21.v4, pkg6, v28, v18)
	if v7 != nil {
		return v15, v7
	}
	return &v33, nil
}