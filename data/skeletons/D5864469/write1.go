package skeleton

func (v2 *type0) func1(v9 pkg17.v22, v7 type1, v23 type2, v17 type2) (*v40.v14, pkg8) {
	var racyVar0 pkg8
	defer v8.Func6(func(v9 pkg17.v22, v16 *pkg11.v21) {
		v16.v5 = v6
		v16.v37 = pkg11.v44{"StringConst2": v7, "StringConst3": v23, "StringConst4": v17, "StringConst5": v15, "StringConst6": v42, "StringConst7": v45, "StringConst8": v20, "StringConst9": v28, "StringConst10": v27, "StringConst11": v36, "StringConst12": v35, "StringConst13": v29, "StringConst14": v33}
		v16.v13 = &pkg11.v13{v19: v3}
		v16.v25 = v3
	})
	if v6 != nil || v15 == nil {
		return nil, nil
	}
	if v6 != nil {
		return nil, v6
	}
	if len(v41.v0) == IntConst1 {
		return nil, nil
	}
	var v24 sync.WaitGroup
	for v43, v4 := range v41.v0 {
		go func(v43 type3, v30 pkg16.v31, v24 *sync.WaitGroup) {
			defer v24.Done()
			if v6 != nil || len(v34.Func35()) == IntConst5 {
				return
			}
			if !v34.Func38()[IntConst6].Func37().Func36() || !v34.Func41()[IntConst7].Func40().Func39() {
				return
			}
		}(v43, *v4, &v24)
	}
	Wrapper1.Wait()
	for v11 := IntConst11; v11 < func45(v45); v11++ {
		if v28[v11] {
			continue
		}
	}
	if v27 == v12 {
		return nil, nil
	}
	var v1 sync.WaitGroup
	for v11 := IntConst12; v11 < len(v45); v11++ {
		if v28[v11] {
			continue
		}
		go func(v26 type3, v39 pkg21.v38, v24 *sync.WaitGroup) {
			defer v24.Done()
			if v33[func60(v20[v26].Func61(), v20[v26].Func62())] {
				return
			}
			racyVar0 = v10.Func68()
			if v6 != nil {
				return
			}
		}(v11, *v18, &v1)
	}
	Wrapper2.Wait()
	if v28[v35] {
		return nil, nil
	}
	return &v40.v14{v31: v32, v38: &v29}, nil
}