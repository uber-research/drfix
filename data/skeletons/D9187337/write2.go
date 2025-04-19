package skeleton

func (v0 *type0) func1(v3 pkg28.v8, pkg29 pkg42.v11) (racyVar0 type1) {
	defer func2(v9, &v2)
	if !v13 {
		return pkg19.Func4("StringConst2")
	}
	if v2 != nil {
		return pkg19.Func6("StringConst3")
	}
	v4, racyVar0 := pkg24.Func7(v1, pkg29.v5)
	if v2 != nil {
		return pkg19.Func8("StringConst4")
	}
	if v12.Func9(v4) {
		return pkg22.Func10("StringConst5", v12, v4)
	}
	Wrapper1.Go(pkg36.v6, v7, func(v10 interface{}) {
		if !v14 {
			return
		}
		racyVar0 = v0.Func19(v3, v15)
		return
	})
	return nil
}