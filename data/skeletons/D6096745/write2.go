package skeleton

func (v5 *type0) func1(v4 v10.v16, v6 *v12.v3) (v19 *v12.v8, v2 type1) {
	racyVar0 := map[type2]type2{"StringConst0": pkg32.v25.Func1()}
	defer v5.v18.Func3(v4, v5.v22.Func4(), "StringConst1").Func2(func(pkg25 pkg25.v16, v14 *pkg15.v13) {
		v14.v23 = v24
		v14.v1 = v2
		v14.v17 = v15
	})
	Wrapper1.Go(func() (v2 type1) {
		if v6 != nil {
			for _, v0 := range v6.v21 {
				if v0 == nil {
					continue
				}
				racyVar0[v0.Func6().Func5()] = "StringConst2"
			}
		}
		return nil
	})
	Wrapper1.Go(func() (v2 type1) {
		if v20 == nil && v11.v7 != nil {
			racyVar0["StringConst3"] = pkg32.v9.Func9()
		}
		return nil
	})
	if v2 != nil {
		return nil, v2
	}
	if v2 != nil {
		return nil, v2
	}
	Wrapper1.Wait()
	if v19 == nil {
		return &v12.v8{}, nil
	}
	return v19, nil
}