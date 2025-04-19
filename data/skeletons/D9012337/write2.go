package skeleton

func (v11 *type0) func1(v2 pkg21.v3, v4 *pkg11.v5, v6 *v8.v7) ([]type1, type2) {
	if v6 == nil || len(v6.v0) == IntConst0 {
		return nil, pkg9.Func2(v12)
	}
	if func3(v6.v9) == IntConst1 {
		v6.racyVar0 = "StringConst0"
	}
	if pkg1.Func4(v6.v0, "StringConst1") {
		return func5(v2, v4, v11.v10, v6)
	}
	return func6(v2, v4, v11.v1, v6)
}