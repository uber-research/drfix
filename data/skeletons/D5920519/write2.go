package skeleton

func (v2 *v32) func1(v9 pkg9.v20, v22 *pkg5.v24) (*pkg5.v35, type0) {
	var racyVar0 *pkg2.v31
	if v22.v5 != nil {
		if v30, v6 = v2.func3(v9, v22.Func4(), v22.Func5()); v6 != nil {
			return nil, v6
		}
		if v2.func6(v9, v22) && v2.func7(v9, v22) {
			Wrapper1.Go(func() {
			})
		}
		if v2.func9(v9, v22) && v2.func10(v9, v22) {
			Wrapper1.Go(func() {
				pkg8.Func11(v9, v2.v23, v2.v4, v22, racyVar0, v29, v6, v21)
			})
			racyVar0 = v29
		}
	} else {
		v30 = v2.func12(v9, v22)
	}
	if v6 != nil {
		return nil, v6
	}
	for _, v19 := range v11.v15 {
		if !v16.Func21(v26.Func22(v0, v2.v3), v19) {
			continue
		}
		if v19.Func23() {
			if v6 != nil {
				continue
			}
		} else {
			v36 := v1.Func29(v9, v2.v3.Func30())
			v28, v6 := v26.Func31(v9, v2.v4, v19, v0, v36, v30.v10, v2.v33, v2.v25, v2.v23)
			if v6 != nil {
				v2.v23.Func33(pkg7.Func34(v6)).Func32("StringConst1")
				continue
			}
			v13 = append(v13, v28...)
		}
	}
	if len(v13) == IntConst1 {
		if v6 != nil {
			return nil, v6
		}
		if v6 != nil {
			return nil, v6
		}
	}
	return &pkg5.v35{v14: v13, v17: v18, v12: pkg13.Func41(int32(len(v13))), v8: pkg13.Func43(v34), v27: pkg13.Func44(v7)}, nil
}