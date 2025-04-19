package skeleton

func (v4 *v6) func1(v5 pkg0.v8, v13 *v9) (*v1, type0) {
	v3 := pkg1.Func3(v5, pkg1.Func4(v2), pkg1.Func5(func(v14 type1, v15 interface{}) type0 {
		if v3 != nil {
			return v3
		}
		if len(v13.v12) > IntConst0 {
			v0.racyVar0 = v7
		}
		return nil
	}, v13.v12))
	if v3 != nil {
		return nil, v3
	}
	return &v1{v11: v10}, nil
}