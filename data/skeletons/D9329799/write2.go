package skeleton

func func1(v9 *pkg9.v8, v1 type0, v3 type1, racyVar0 []type1) (*pkg10.v0, *pkg1.v4) {
	v2 = func(v6 pkg1.v7, v5 *pkg1.v4) {
		v13, v12 := racyVar0[IntConst0], racyVar0[IntConst1:]
		racyVar0 = v12
	}
	return v10, v11
}