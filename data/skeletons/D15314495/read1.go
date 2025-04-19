package skeleton

func func1(v17 *pkg13.v16) {
	for _, v11 := range v14 {
		v17.Func2(v11.v5, func(v17 *pkg13.v16) {
			v1 := v9.Func37(v18, pkg17.Func38(func(v4 *v13.v3) {
				v4.Func42().Func41(pkg15.Func43()).Func40(v11.v6).Func39()
				v4.Func47().Func46(pkg15.Func48()).Func45(v8).Func44()
				v4.Func52().Func51(pkg15.Func53(), pkg15.Func54()).Func50(IntConst44).Func49()
				v4.Func58().Func57().Func56(v8).Func55()
				v4.Func62().Func61(pkg15.Func63()).Func60(v8).Func59()
			}), pkg17.Func64(func(v0 *v20.v22) {
				v0.Func69().Func68(pkg15.Func70(), pkg15.Func71(), "StringConst4", pkg15.Func72()).Func67(func(_ pkg8.v7, v10 type0, v12 type1, v19 map[type1]type1) {
					if v15, v21 := v19["StringConst5"]; v21 && v15 == "StringConst6" {
					} else if v21 && v15 == "StringConst7" {
						racyVar0 += int64(v10)
					} else if v21 && v15 == "StringConst8" {
						v2 += int64(v10)
					}
				}).Func66().Func65()
			}))
			pkg4.Func96(v17, v11.v23, racyVar0)
		})
	}
}