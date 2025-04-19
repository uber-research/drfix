package skeleton

func (v8 *type0) func1(v3 pkg7.v5, v1 *type1, v10 *type2, v4 func(v3 pkg7.v5)) {
	for _, racyVar0 := range v7 {
		go pkg1.Func2(func() {
			v2 := pkg9.Func3(v3)
			v2 = pkg0.Func4(v2, racyVar0)
			v2, v0 := pkg7.Func5(v2, IntConst0*pkg5.v9)
			func6(v2)
			func7()
		}, v8.v6)
	}
}