package skeleton

func (v13 *type0) func1(v12 pkg39.v19, v24 *v15.v50) (v1 *v15.v28, racyVar0 type1) {
	defer v13.v32.Func4(v12, pkg13.Func6(v13.v21).Func5(), "StringConst0").Func3(func(pkg39 pkg39.v19, v35 *pkg26.v16) {
		if v24 != nil {
			v35.v39 = pkg26.v49{"StringConst1": v24.Func7(), "StringConst2": v24.Func8(), "StringConst3": v24.Func9()}
			if v35.v30 != nil {
				v10 = v35.v30
			}
			v35.v9 = v10
			v35.v23 = v18
			v5 := len(v1.Func11()) == IntConst0
			v35.v45 = map[type2]type2{"StringConst4": v24.Func13().Func12(), "StringConst5": pkg14.Func14(v5), "StringConst6": v33}
			v35.v11 = &pkg26.v11{v14: pkg11.v36, v38: map[type2]type2{"StringConst7": v24.Func16().Func15(), "StringConst8": "StringConst9"}}
		}
	})
	if v24 == nil || v24.v46 == nil || len(v24.v8) == IntConst1 || v24.v0 == nil {
		return nil, pkg18.Func18("StringConst10", nil)
	}
	if v40 = v13.v25.Func30(v12, v7, v4); v40 != nil {
		return nil, v40
	}
	if v27 {
		Wrapper1.Go(func() type1 {
			if v31 != nil && v13.v20.v47 != nil && v24.v41 != nil {
				v29, racyVar0 = v13.v20.v47.Func42(v17, *v24.v41)
			}
			return nil
		})
	}
	Wrapper1.Go(func() type1 {
		defer v13.v32.Func51(v12, pkg13.Func53(v13.v21).Func52(), "StringConst17").Func50(func(v12 pkg39.v19, v35 *pkg26.v16) {
			v35.v11 = &pkg26.v11{v14: pkg11.v36, v38: map[type2]type2{"StringConst18": "StringConst19"}}
		})
		return nil
	})
	if !v34 && v24.v3 != nil && (type2)(*v24.v3) != "StringConst22" {
		Wrapper1.Go(func() type1 {
			v48, v26, racyVar0 = v13.func76(v12, v22, v6)
			return nil
		})
	}
	if v13.func77(v12, v4) {
		Wrapper1.Go(func() type1 {
			defer v13.v32.Func79(v12, pkg13.Func81(v13.v21).Func80(), "StringConst23").Func78(func(v12 pkg39.v19, v35 *pkg26.v16) {
				v35.v11 = &pkg26.v11{v14: pkg11.v36, v38: map[type2]type2{"StringConst24": "StringConst25"}}
			})
			return nil
		})
	}
	if func84(v24) && v13.func85(v12, v4) {
		Wrapper1.Go(func() type1 {
			defer v13.v32.Func87(v12, pkg13.Func89(v13.v21).Func88(), "StringConst26").Func86(func(v12 pkg39.v19, v35 *pkg26.v16) {
				v35.v11 = &pkg26.v11{v14: pkg11.v36, v38: map[type2]type2{"StringConst27": "StringConst28"}}
			})
			return nil
		})
	}
	if v13.func94(v12, v24) {
		Wrapper1.Go(func() type1 {
			defer v13.v32.Func96(v12, pkg13.Func98(v13.v21).Func97(), "StringConst29").Func95(func(v12 pkg39.v19, v35 *pkg26.v16) {
				v35.v11 = &pkg26.v11{v14: pkg11.v36, v38: map[type2]type2{"StringConst30": "StringConst31"}}
			})
			return nil
		})
	}
	if v44 && len(v24.Func111()) == IntConst6 {
		Wrapper1.Go(func() type1 {
			if v37 == nil || len(v37.v2) == IntConst7 {
				return nil
			}
			return nil
		})
	}
	if v10 := Wrapper1.Wait(); v10 != nil {
		return nil, v10
	}
	if v42 != nil {
		return nil, v42
	}
	return &v43, nil
}