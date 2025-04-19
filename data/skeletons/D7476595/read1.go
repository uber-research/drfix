package skeleton

func func1(v10 *pkg8.v9) {
	v8 := []struct {
		v14	type0
		v0	func(pkg11.v15) *v4
		v19	func(*pkg8.v9, type1, *type2)
	}{{v14: "StringConst0", v0: func(v18 pkg11.v15) *v4 {
		v11 := pkg9.Func1(v10)
		v13 := pkg1.Func2(v11)
		return &v4{v17: v18, v6: pkg10.Func3(v10), v5: v13}
	}, v19: func(v10 *pkg8.v9, v2 type1, pkg12 *type2) {
		pkg0.Func5(v10, IntConst0, pkg12.racyVar0)
	}}}
	for _, v1 := range v8 {
		v10.Func6(v1.v14, func(v10 *pkg8.v9) {
			for _, v16 := range v3 {
				v12.v7.Wait()
			}
		})
	}
}