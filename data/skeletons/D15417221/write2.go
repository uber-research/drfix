package skeleton

func func1(v17 *pkg13.v16) {
	for _, v11 := range v14 {
		v17.Func2(v11.v4, func(v17 *pkg13.v16) {
			v1 := v9.Func39(v18, pkg17.Func40(func(v3 *v13.v2) {
				v3.Func44().Func43(pkg15.Func45()).Func42(v11.v5).Func41()
				v3.Func49().Func48(pkg15.Func50()).Func47(v7).Func46()
				v3.Func54().Func53(pkg15.Func55(), pkg15.Func56()).Func52(IntConst44).Func51()
				v3.Func60().Func59().Func58(v7).Func57()
				v3.Func64().Func63(pkg15.Func65()).Func62(v7).Func61()
			}), pkg17.Func66(func(v0 *v20.v23) {
				v0.Func87().Func86(pkg15.Func88(), pkg15.Func89(), "StringConst18", pkg15.Func90()).Func85(func(_ pkg8.v6, v10 type0, v12 type1, v19 map[type1]type1) {
					if v15, v21 := v19["StringConst19"]; v21 && v15 == "StringConst20" {
						racyVar0 += int64(v10)
					} else if v21 && v15 == "StringConst21" {
						v22 += int64(v10)
					}
				}).Func84().Func83()
			}))
			pkg4.Func101(v17, v11.v8, racyVar0)
		})
	}
}