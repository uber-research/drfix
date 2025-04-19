package skeleton

func (v11 *type0) func1(v5 pkg8.v9, v4 v14.v16, v0 pkg11.v6) ([]*v1.v13, pkg4) {
	var racyVar0 pkg4
	v2 = pkg0.Func1(v5, pkg0.Func2(func(v5 pkg8.v9) pkg4 {
		v12, racyVar0 = v11.v10.Func3(v5, v4, v0.v15)
		if v2 != nil {
			return pkg2.Func4("StringConst0", v2, map[type1]interface{}{"StringConst1": v4.Func5()})
		}
		return nil
	}), pkg0.Func6(func(v5 pkg8.v9) pkg4 {
		if !v0.v8.v7 {
			return nil
		}
		v3, racyVar0 = v11.func7(v5, v4, nil)
		if v2 != nil {
			return v2
		}
		return nil
	}))
	if v2 != nil {
		return nil, v2
	}
	return v11.func8(v12, v3), nil
}