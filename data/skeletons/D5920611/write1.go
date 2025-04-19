package skeleton

func func1(v11 *pkg18.v10) {
	for _, racyVar0 := range v9 {
		v11.Func25(v6.v15, func(v11 *pkg18.v10) {
			defer v4.Func27()
			v3 := sync.WaitGroup{}
			if v6.v19 != nil {
				v8.Func44().Func43(pkg19.Func45(), v6.v19.v12).Func42(func(pkg16.v5, pkg0.v0) type0 {
					v3.Done()
					return racyVar0.v19.v1
				})
			}
			for _, v2 := range v6.v7 {
				v14.Func49().Func48(pkg19.Func50(), v2.v17, v2.v16, v2.v18).Func47(func(pkg16.v5, type1, map[type1]type1, ...pkg10.v13) {
					v3.Done()
				}).Func46(v2.v1)
			}
		})
	}
}