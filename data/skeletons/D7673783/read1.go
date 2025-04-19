package skeleton

func func1(v8 *pkg3.v7, v0 type0, v1 type1) *pkg2.v3 {
	racyVar0 := &v0
	v2 := pkg2.Func1(pkg1.Func2(func(v9 pkg1.v6, v5 *pkg1.v4) {
		*racyVar0--
	}))
	return v2
}