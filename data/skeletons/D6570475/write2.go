package skeleton

func func1(v5 v6, v12 pkg8.v8, v11 pkg8.v0, v7 pkg8.v20, v10 pkg2.v9, v18 pkg9.v16) type0 {
	if v12 == nil {
		return pkg4.Func5("StringConst3")
	}
	if v11 == nil {
		return pkg4.Func7("StringConst5")
	}
	if v10 == nil {
		return pkg4.Func9("StringConst7")
	}
	v4, racyVar0 := v7.Func12(v15.Func13(pkg7.v1*IntConst0*IntConst1*-IntConst2), v15.Func14(pkg7.v1*-IntConst3))
	if v3 != nil {
		return v3
	}
	for _, v19 := range v4 {
		v2 <- struct{}{}
		go func(v13 *pkg1.v14, v21 type1, v12 pkg8.v8, v11 pkg8.v0, v18 pkg9.v16) {
			defer func() {
			}()
			if racyVar0 = func21(v13, v5.v17, v21, v12, v11, v18); racyVar0 != nil {
			}
		}(v13, v19, v12, v11, v18)
	}
	return nil
}