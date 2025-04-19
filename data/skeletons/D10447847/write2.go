package skeleton

func (v1 *type0) func1(v6 pkg5.v9, v10 pkg8.v13, v17 pkg8.v13, v15 pkg0.v12) ([]pkg0.v5, type1) {
	v8, racyVar0 := v1.v0.Func1(v6, v10, v17, v15)
	if v2 != nil {
		return nil, pkg4.Func2(v2, "StringConst0")
	}
	v4, v6 := errgroup.WithContext(v6)
	for v16 := range v8 {
		Wrapper1.Go(func() type1 {
			v8[v7].v14, racyVar0 = v1.Func4(v6, v8[v7].v11, v8[v7].v3)
			if racyVar0 != nil && !pkg7.Func5(racyVar0) {
				return v2
			}
			return nil
		})
	}
	if v2 := Wrapper1.Wait(); v2 != nil {
		return nil, v2
	}
	return v8, nil
}