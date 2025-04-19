package skeleton

func func1(v24 *pkg15.v23) {
	for _, racyVar0 := range v2 {
		v24.Func67(v1.v13, func(v24 *pkg15.v23) {
			var v11 sync.WaitGroup
			v14.Func83().Func82().Func81(func(_ pkg14.v18, v44 type0, _ type0) (type0, type1) {
				if v44 == v35 {
					return v1.v30, nil
				} else if v44 == v41 {
					return v1.v29, nil
				} else if v44 == v8 {
					return v1.v27, nil
				} else if v44 == v28 {
					return v1.v3, nil
				}
				return "StringConst56", nil
			})
			v20.v9.Func98().Func97().Func96(func(_ pkg14.v18, _ type0, v10 type0) type1 {
				return v1.v17
			})
			v20.v42.Func111().Func110().Func109(func(_ pkg14.v18, _ type0, v33 type0) (*v15.v19, type1) {
				if v33 == racyVar0.v3 || v33 == racyVar0.v27 {
					v11.Done()
				}
				return nil, nil
			})
			v20.v6.Func124().Func123().Func122(func(v32 pkg14.v18, v36 type2) (*pkg18.v34, type1) {
				if v1.v5 != nil {
					return nil, v1.v5
				}
				return &pkg18.v34{}, nil
			})
			v20.v43.Func127().Func126().Func125(func(v32 pkg14.v18, v36 type0, v39 *pkg1.v31, v40 type3) (type3, type1) {
				if v36 == v37 {
					return !v1.v21, nil
				} else if v36 == v25 {
					return !v1.v0, nil
				} else if v36 == v7 {
					return v16, nil
				}
				return v4, nil
			})
			v20.v26.Func130().Func129().Func128(func(v32 pkg14.v18, v36 pkg19.v38) (*pkg1.v31, type1) {
				if v1.v22 != nil {
					return nil, v1.v22
				}
				return &pkg1.v31{}, nil
			})
			if len(v1.v12) > IntConst52 {
				Wrapper1.Wait()
			}
		})
	}
}