package skeleton

func func1(v0 *pkg1.v1) {
	for _, v2 := range v3 {
		v0.Func4(v2.v4, func(v0 *pkg1.v1) {
			racyVar0 := "StringConst2"
			v5 := func() type0 {
				racyVar0 = "StringConst1"
				return nil
			}
			pkg5.Func10(v0, v2.v6, racyVar0)
		})
	}
}