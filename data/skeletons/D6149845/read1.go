package skeleton

func func1(v11 *pkg18.v10) {
	v11.Func6("StringConst5", func(v11 *pkg18.v10) {
		defer v2.Func8()
	})
	for _, v7 := range v0 {
		v11.Func31(v7.v3, func(v11 *pkg18.v10) {
			defer v2.Func33()
			racyVar0 := "IntConst1"
			v14.Func65().Func64(pkg21.Func66(), pkg21.Func67()).Func63(func(v1 pkg16.v6, v4 pkg8.v9) (pkg10.v8, type0) {
				racyVar0++
				if v5 == IntConst9 && v7.v13 {
					return nil, nil
				}
				return nil, v12
			}).Func62(v15)
		})
	}
	v11.Func82("StringConst19", func(v11 *pkg18.v10) {
		defer v2.Func84()
	})
}