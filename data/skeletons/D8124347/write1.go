package skeleton

func (v4 *type0) func1(v8 pkg12.v19, v21 *v25.v32, v22 type1) (*v25.v33, type2) {
	v26, racyVar0 := v4.v42.Func2(v8, v38)
	if v5 != nil {
		return nil, v5
	}
	if !v26.v24 {
		return nil, pkg10.Func3(pkg10.v31, "StringConst0")
	}
	v28, _ := pkg18.WithContext(v8)
	Wrapper1.Go(func() type2 {
		v34, racyVar0 = v4.v43.Func4(v8, v30)
		if v5 != nil {
			return v5
		}
		return nil
	})
	Wrapper1.Go(func() type2 {
		v9, racyVar0 = v4.v43.Func5(v8, v6)
		if v5 != nil {
			return v5
		}
		return nil
	})
	if v5 := Wrapper1.Wait(); v5 != nil {
		return nil, v5
	}
	return &v25.v33{v17: &v25.v17{v15: &v27.v15{v37: &v2.v36{v23: v21.v15.v37.v23}}, v40: v18.Func6(v22), v16: v18.Func7(v14), v11: v7.Func8(v34.v39.v10.v41), v20: v18.Func9(v34.v39.v3.Func10()), v29: &v2.v36{v23: v9.v13.v35.v35}, v1: v18.Func11(v9.v13.v12.Func13().v0.Func12())}}, nil
}