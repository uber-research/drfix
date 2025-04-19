package skeleton

func (v14 *type0) func1(v4 *pkg3.v2, v17 *pkg3.v18, v8 *pkg3.v7) (type1, type2) {
	if v4 == nil || v4.v6 == nil {
		return v3, pkg1.Func1("StringConst0")
	}
	if v8 == nil {
		return v3, pkg1.Func2("StringConst1", v4.v9)
	}
	if v17 == nil {
		return v3, pkg1.Func3("StringConst2", v4.v9)
	}
	if v17.v26 == nil || v17.v26.v12 == nil {
		return v3, pkg1.Func4("StringConst3", v4.v9)
	}
	v14.v0.Do(func() {
		if v8.v21 == nil {
			return
		}
		v14.racyVar0 = v8.v21.Func5(v22.v25{v9: v14.v1.v9, v20: pkg4.Func6(v14.v1.v10), v19: v14.v1.v19, v16: v14.v1.v16, v11: v14.v1.v11})
	})
	if v8.v21 == nil {
		return v13, nil
	}
	if v5 != nil {
		return v3, v5
	}
	if v23.Func9() {
		if v5 != nil {
			return v3, v5
		}
		return v24, nil
	}
	if v23 == v14.v15 {
		return v13, nil
	}
	return v3, nil
}