package skeleton

func func1(v0 func(pkg0.v7, pkg3.v15) (pkg2.v14, type0), v9 func(pkg0.v7, pkg3.v15) (pkg2.v12, type0)) (type1, type0) {
	racyVar0 := make(map[pkg3.v15]func() (pkg2.v14, type0))
	v2 := func(v3 pkg0.v7, v11 pkg3.v15) (pkg2.v14, type0) {
		if v13, v16 := racyVar0[v11]; v16 {
			return func3()
		}
		racyVar0[v11] = sync.OnceValues(func() (pkg2.v14, type0) {
			return func4(v3, v11)
		})
		return v4[v11]()
	}
	v6 := func(v3 pkg0.v7, v5 pkg3.v15) (pkg2.v12, type0) {
		if v10, v16 := v8[v5]; v16 {
			return func5()
		}
		v8[v5] = sync.OnceValues(func() (pkg2.v12, type0) {
			return func6(v3, v5)
		})
		return v8[v5]()
	}
	return &type2{v0: v2, v9: v6, v4: v4, v8: v8, v17: make(map[pkg3.v15]struct{}), v1: make(map[pkg3.v15]struct{})}, nil
}