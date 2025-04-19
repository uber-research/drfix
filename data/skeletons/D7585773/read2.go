package skeleton

func func1[type0 v17](v4 pkg0.v7, v0 []type0, v18 []type1[type0], v5 type2) (v11 []type0, racyVar0 type3) {
	if len(v0) == IntConst0 {
		return nil, nil
	}
	if v5 <= IntConst1 {
		return nil, pkg1.Func2("StringConst0")
	}
	v4, cancel := pkg0.Func5(v4)
	defer cancel()
	v9, v4 := pkg3.WithContext(v4)
	for _, v15 := range v18 {
		Wrapper1.Go(func() type3 {
			for {
				if v4.Func6() != nil {
					return v4.Func7()
				}
				if len(v13) == IntConst2 {
					return nil
				}
				for _, v13 := range v10 {
					v2 <- v13
				}
				if v12 != nil && !pkg1.Func12(v12, pkg0.v16) {
					return v12
				}
			}
		})
	}
	Wrapper2.Go(func() {
		racyVar0 = Wrapper1.Wait()
		for v14 := range v8.v6 {
			for _, cancel := range v8.v6[v14].v1 {
				cancel()
			}
		}
	})
	for v13 := range v2 {
		if func15(v11) >= v5 {
			Wrapper2.Go(func() {
			})
			return v11, nil
		}
	}
	return v11, v3
}