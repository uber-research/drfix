package skeleton

func func1(v0 *pkg5.v1) {
	v2, racyVar0 := pkg8.Func2()
	v0.Func6("StringConst0", func(v0 *pkg5.v1) {
		racyVar0 = v3.Func12(v4, v2)
	})
	v0.Func14("StringConst1", func(v0 *pkg5.v1) {
		racyVar0 = v3.Func21(v4, v2)
	})
}