package skeleton

func (v22 *v16) func1(v12 pkg13.v18, v25 pkg2.v19, v10 *v1, v33 type0) (pkg2.v19, type1) {
	var racyVar0 type2
	defer v22.v31.v28.Func2("StringConst0").Func1(func(pkg13 *pkg6.v18) {
		var v9 map[type2]type2
		if v10.v29 != nil {
			v9 = v10.v29.Func3()
		}
		pkg13.v32 = pkg6.v34{"StringConst1": v9, "StringConst2": v10}
		pkg13.v6 = v7
		pkg13.v23 = v17
	})
	if v7 != nil {
		return nil, v7
	}
	for v24, v35 := range v5 {
		Wrapper1.Go(func() type1 {
			if v33 {
			} else {
				racyVar0 = pkg15.v14
			}
			v12, cancel := pkg13.Func12(v12, pkg0.Func13(v22.v31.v20, v10.v27, v13, &v15.v11{}, v10.v29, int(v22.v4.v30.v0.Func14())))
			defer cancel()
			return v8
		})
	}
	Wrapper1.Wait()
	for _, v21 := range v3 {
		if v21 == nil {
			continue
		}
	}
	return func26(v22.v31, v25, v2.v26, nil), nil
}