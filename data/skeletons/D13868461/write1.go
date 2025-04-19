package skeleton

func (v3 *type0) func1(v8 pkg11.v13, v20 pkg2.v14, pkg5 pkg9.v9) (v1 *pkg2.v2, racyVar0 type1) {
	v5, racyVar0 := v3.v21.Func1(v8, pkg5, v20.v19)
	if v4 != nil {
		return nil, v4
	}
	v6, v0 := errgroup.WithContext(v8)
	Wrapper1.Go(func() type1 {
		v17, racyVar0 = v3.func3(v0, v20.v19, pkg5, v12)
		return v4
	})
	Wrapper1.Go(func() type1 {
		v7, racyVar0 = v3.func4(v0, v20.v19, pkg5, v12)
		return v4
	})
	Wrapper1.Go(func() type1 {
		return nil
	})
	if v4 := Wrapper1.Wait(); v4 != nil {
		return nil, v4
	}
	return &pkg2.v2{v16: v17, v18: v7, v15: nil, v11: v10}, nil
}