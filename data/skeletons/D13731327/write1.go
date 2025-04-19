package skeleton

func func1(v0 *pkg60.v1) {
	defer func4(v2)
	for v3, v4 := range v5 {
		v0.Func26(v3, func(v0 *pkg60.v1) {
			v2.v6.racyVar0 = nil
		})
	}
}