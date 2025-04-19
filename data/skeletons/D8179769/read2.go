package skeleton

func (v1 *type0) func1(v2 pkg9.v5, v3 type1) ([]type1, pkg2) {
	if v0 != nil {
		return nil, v0
	}
	if v7 == nil || len(v7) == IntConst0 {
		return nil, nil
	}
	for _, racyVar0 := range v7 {
		v4.Func4(func() pkg2 {
			_, v0 := v1.Func5(v2, racyVar0.Func7().Func6(), []type1{v3})
			if v0 != nil {
				return v0
			}
			return nil
		})
	}
	if v0 := Wrapper1.Wait(); v0 != nil {
		return v8, v6.Func11(v0, "StringConst0")
	}
	return v8, nil
}