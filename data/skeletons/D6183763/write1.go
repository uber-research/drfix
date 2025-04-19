package skeleton

func (v1 *type0) func1(v5 pkg8.v7) type1 {
	if v4 != nil {
		return v4
	}
	for _, racyVar0 := range v11 {
		go func() {
			defer func() {
			}()
			if v4 := v1.v0.Func5(v5, v12, v3, *racyVar0.v15); v4 != nil {
			}
		}()
		if v4 != nil {
			continue
		}
		if v4 != nil {
			continue
		}
		if v14 && pkg7.Func13().Func12().Func11(v16.v17.Func14()) > pkg7.v2*IntConst0 {
			continue
		}
		if v6 == "StringConst1" || func15(v6, *v16.v15) {
			if v4 := v1.pkg3.Func16(v5, v8); v4 != nil {
				continue
			}
		}
		if v4 := v1.pkg3.Func18(v5, v13, v9); v4 != nil {
			continue
		}
	}
	return v10
}