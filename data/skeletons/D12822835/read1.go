package skeleton

func func1(v0 *pkg2.v1) {
	for v2, v3 := range v4 {
		v0.Func2(v2, func(v0 *pkg2.v1) {
			v5.v6.Wait()
			pkg5.Func5(v0, IntConst2, v5.racyVar0)
		})
	}
}