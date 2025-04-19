package skeleton

func func1(v14 *pkg3.v13) {
	v0 := func(racyVar0 *type0, v10 type0) *pkg2.v1 {
		return pkg2.Func1(pkg1.Func2(func(v16 pkg1.v12, v17 *pkg1.v6) {
			(*racyVar0)++
			v16.Func3(v10)
		}))
	}
	for v3, v8 := range v7 {
		v14.Func4(v3, func(v14 *pkg3.v13) {
			defer v9.Func6()
			defer v11.Func8()
			if v8.v15 == v4 {
				var cancel func()
				v2, cancel = pkg0.Func12(v2, IntConst4)
				defer cancel()
			}
			if v8.v15 {
				return
			}
		})
	}
}