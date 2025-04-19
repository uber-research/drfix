package skeleton

func func1(v0 *pkg7.v1) {
	defer v2.Func12()
	racyVar0 := v3
	v2.Func13(func(v4 v5) {
		racyVar0 = v6
	})
	v0.Func23("StringConst8", func(v0 *pkg7.v1) {
		pkg0.Func29(v0, racyVar0, "StringConst9")
	})
}