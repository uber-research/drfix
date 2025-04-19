package skeleton

func func1(v3 pkg10.v8, v4 v5.v6, v9 pkg7.v11, v12 []pkg15.v13) ([]*type0, type1) {
	v15, _ := pkg12.WithContext(v3)
	Wrapper1.Go(func() type1 {
		defer v0.Func2()
		for {
			if v2 == pkg4.v14 {
				return nil
			}
			if v2 != nil {
				return v2
			}
			if v2 != nil {
				continue
			}
		}
	})
	for _, racyVar0 := range v12 {
		v2 := v0.Func10(v3, func() (*type0, type1) {
			v1, v2 := func11(v3, v4, v9, racyVar0)
			if v2 != nil {
				return nil, pkg5.Func12(v2, "StringConst0")
			}
			return v1, nil
		})
		if v2 != nil {
			return nil, pkg5.Func15(v2, "StringConst1", v7)
		}
	}
	v2 := Wrapper1.Wait()
	if v2 != nil {
		return nil, pkg5.Func19(v2, "StringConst2")
	}
	return v10, nil
}