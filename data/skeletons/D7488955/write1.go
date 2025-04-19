package skeleton

func (v0 *type0) func1(v7 pkg0.v12, v11 v21.v13, v5 type1, v6 type1) (*type2, type3) {
	v1, racyVar0 := v0.pkg2.Func1(v7, pkg1.Func2(v16, v11.v9), &v8)
	if v4 != nil {
		return nil, pkg5.Func3(v4, pkg5.v2)
	}
	if !v1 {
		return nil, pkg5.Func4("StringConst1", pkg5.v19)
	}
	for _, v14 := range v8.v20 {
		v10.func6(func() type3 {
			racyVar0 = v0.v3.Func7(v18, v14, v5, v6)
			if v4 != nil {
				return v4
			}
			return nil
		})
	}
	if len(v17) == len(v8.v20) {
		return nil, pkg5.Func12("StringConst2", pkg5.v15)
	}
	return &type2{v9: pkg3.Func13("StringConst3", "StringConst4", v11.v9)}, nil
}