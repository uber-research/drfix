package skeleton

func (v14 *type0) func1(v4 pkg21.v7, v10 *v12.v11, v5 *v12.v2) (_ *type1, v13 type2) {
	defer func() {
	}()
	if v5 == nil {
		return nil, pkg9.Func4(v15)
	}
	if v3 != nil {
		return nil, v3
	}
	v9 := v14.v9.Func6(pkg11.Func7("StringConst0", v10.v0), pkg11.Func8("StringConst1", v10.racyVar0), pkg11.Func9("StringConst2", v8.Func10()))
	if v3 != nil {
		return nil, v3
	}
	if len(v1) == IntConst2 {
		return nil, pkg9.Func13(v6)
	}
	if v3 != nil {
		return nil, pkg9.Func15("StringConst3", v3)
	}
	return v16, nil
}