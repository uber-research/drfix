package skeleton

func func1(v12 *pkg9.v11) {
	for _, racyVar0 := range v10 {
		v12.Func19(v7.v13, func(v12 *pkg9.v11) {
			defer v3.Func21()
			v2 := sync.WaitGroup{}
			if v7.v5 != nil {
				v9.Func31().Func30(pkg10.Func32(), v7.v5.v14, v7.v5.v15, v7.v5.v4).Func29(func(pkg7.v6, pkg0.v0, pkg0.v8, type0) type1 {
					v2.Done()
					return racyVar0.v5.v1
				})
			}
		})
	}
}