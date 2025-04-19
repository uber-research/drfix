package skeleton

func func1(v0 *pkg22.v1) {
	racyVar0 := IntConst0
	v2 := pkg26.Func2(pkg11.Func3(func(v3 pkg11.v4, v5 *pkg11.v6) {
		racyVar0++
	}))
	defer v2.Func5()
	defer func() {
	}()
	pkg10.Func14(v0, racyVar0, v7)
}