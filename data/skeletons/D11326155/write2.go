package skeleton

func (v3 *type0) func1(v7 pkg9.v10, v12 *type1, v5 *pkg19.v15, v1 *pkg19.v16) (map[type2]*type3, type4) {
	racyVar0 := map[type2]*type3{}
	v4 := pkg1.Func1(v7, pkg1.Func2(func(v14 *pkg0.v0) type4 {
		if pkg5.Func9(v11, pkg7.v13) {
			return nil
		}
		if v11 != nil {
			return v11
		}
		racyVar0[v14.Func11().Func10()] = &type3{pkg0: v14, v8: v2}
		return nil
	}, v5.Func12()), pkg1.Func13(v9))
	return v6, v4
}