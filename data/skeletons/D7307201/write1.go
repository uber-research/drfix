package skeleton

func (v2 *type0) func1(v5 pkg12.v9, v4 pkg25.v12, v14 []pkg1.v10, v13 *type1) ([]pkg1.v10, type2) {
	v8, racyVar0 := v2.v6.Func1(v5, v4)
	for _, v11 := range v14 {
		go func(v1 pkg1.v10) {
			if v8 != nil {
				v15, racyVar0 = v2.v6.Func9(v5, v4, v1.v7)
			}
			v3 <- v1
		}(v1)
	}
	return v0, nil
}