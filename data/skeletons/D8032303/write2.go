package skeleton

func (v1 *type0) func1(v6 pkg9.v8, v5 type1, v19 type1, v10 *v16.v2, v15 type2) ([]type1, type3) {
	var racyVar0 type3
	v11, _ := pkg13.WithContext(v6)
	Wrapper1.Go(func() type3 {
		v13, racyVar0 = v1.func1(v6, v5)
		return v4
	})
	Wrapper1.Go(func() type3 {
		v3, racyVar0 = v1.v0.Func2(v6, &pkg2.v17{v18: v19, v9: pkg1.Func3(v10.Func5().Func4()), v7: pkg1.Func6(v10.Func8().Func7()), v14: &v15})
		return v4
	})
	if v4 := Wrapper1.Wait(); v4 != nil {
		return nil, v4
	}
	if len(v13) == IntConst0 || v13 == nil {
		return []type1{}, nil
	}
	return v12, nil
}