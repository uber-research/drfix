package skeleton

func (v1 *v15) func1(v4 pkg15.v8, v11 *pkg3.v12) (pkg3.v9, type0) {
	if v10, v2 := v1.v14.Func1(v4, v3, v0, v5); v2 != nil {
		return nil, v2
	} else if v10 {
		return nil, nil
	}
	for _, racyVar0 := range v13 {
		if v6 != nil {
			Wrapper1.Go(func() {
			})
		} else {
			Wrapper1.Go(func() {
				pkg16.Func18(v4).Func17(pkg10.Func19("StringConst6", racyVar0)).Func16(pkg13.Func20("StringConst7"))
			})
		}
	}
	if len(v7) != IntConst0 {
		return nil, pkg12.Func22(pkg1.Func23(v7, "StringConst8"))
	}
	return nil, nil
}