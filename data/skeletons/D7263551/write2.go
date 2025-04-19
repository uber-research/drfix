package skeleton

func (v1 *type0) func1(v5 pkg9.v7) type1 {
	if v4 != nil {
		return v4
	}
	var racyVar0 type1
	for _, v16 := range v11 {
		go func() {
			defer func() {
			}()
			if v4 := v1.v0.Func6(v5, v12, v3, *v16.v15); v4 != nil {
				racyVar0 = pkg6.Func7(racyVar0, v4)
			}
		}()
		if v4 != nil {
			continue
		}
		if v4 != nil {
			continue
		}
		if v14 && pkg8.Func14().Func13().Func12(v16.v17.Func15()) > pkg8.v2*IntConst0 {
			continue
		}
		if v6 == "StringConst1" || func16(v6, *v16.v15) {
			if v4 := v1.v18.Func17(v5, v8); v4 != nil {
				continue
			}
		}
		if v4 := v1.v18.Func19(v5, v13, v9); v4 != nil {
			continue
		}
	}
	if v4 != nil {
		return v4
	}
	return v10
}