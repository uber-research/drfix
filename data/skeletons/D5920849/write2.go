package skeleton

func (v1 *type0) func1(racyVar0 pkg12.v13, v19 pkg5.v22, v21 type1) (v0 interface{}, v6 type2) {
	if !v20 {
		return func18(v10, v1.v7)
	}
	defer func() {
		for v11 := IntConst1; v11 < v4; v11++ {
			v5 <- type3{v0: v0, v6: v6}
		}
	}()
	defer func() {
	}()
	for _, v12 := range v3.v17 {
		if !v18 {
			continue
		}
		if pkg1.Func31() >= v2 {
			continue
		}
		if v16 {
			racyVar0 = v14.Func41(racyVar0, v9)
			Wrapper1.Go(func() {
				v1.func42(racyVar0, v8, v5)
			})
			continue
		}
	}
	if !v15 {
		return func50(v10, v1.v7)
	}
	return
}