package skeleton

func (v17 *type0) func1(v8 pkg8.v14, v6 *pkg6.v12, v1 *pkg6.v0, v18 type1) (v19 []type2, v13 type3, racyVar0 type4) {
	Wrapper1.Go(func() type4 {
		v5, racyVar0 = v17.func2(v8, v21, v18)
		return v4
	})
	Wrapper1.Go(func() type4 {
		v10, racyVar0 = v17.func4(v8, v11, v18)
		return v4
	})
	v4 = Wrapper1.Wait()
	if v4 != nil {
		return
	}
	for v9, v16 := range v5.v20 {
		if v21.v22[v9] != nil && v21.v15 != nil && v21.v7 > IntConst1 {
			Wrapper2.Go(func() {
			})
		}
	}
	return v2, v3, nil
}