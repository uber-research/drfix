package skeleton

func func1(v0 map[type0]interface{}, v1 pkg3.v2) type1 {
	for v3, racyVar0 := range v0 {
		if _, v4 := v5[v3]; !v4 {
			Wrapper1.Go(func() {
				v1.func2().Func3(v7.v8{
					"StringConst1": "StringConst2",
					"StringConst3": racyVar0,
				}).Func4("StringConst4")
			})
		}
	}
	return v6
}