package skeleton

func (v0 *type0) func1(v4 pkg15.v8, v1 pkg25.v13) (*pkg5.v5, *v7.v14, type1) {
	if v2 != nil {
		return nil, nil, v2
	}
	if !v9.Func3(v9.Func4()) {
		return nil, nil, v15.v12
	}
	v6, racyVar0 := pkg20.Func5(v9)
	if v2 != nil {
		return nil, nil, v2
	}
	v3, v4 := errgroup.WithContext(v4)
	Wrapper1.Go(func() type1 {
		v6.v10, racyVar0 = v0.v15.Func6(v4, v1.Func7())
		return v2
	})
	Wrapper1.Go(func() type1 {
		return v2
	})
	Wrapper1.Go(func() type1 {
		return v2
	})
	Wrapper1.Go(func() type1 {
		return v2
	})
	Wrapper1.Go(func() type1 {
		v6.v11, racyVar0 = v0.v15.Func14(v4, v1.Func15())
		return v2
	})
	if v2 := Wrapper1.Wait(); v2 != nil {
		return nil, nil, v2
	}
	return v9, v6, nil
}