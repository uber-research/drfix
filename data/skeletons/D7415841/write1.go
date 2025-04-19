package skeleton

func (v9 *type0) func1(v1 pkg9.v7, v18 type1, v20 type1) (_ []*pkg0.v19, racyVar0 type2) {
	defer func2(v8, &v0)
	v12, racyVar0 := v9.func8(v5)
	if v0 != nil {
		return nil, v0
	}
	Wrapper1.Go(pkg12.v4, v12, func(v10 interface{}) {
		if !v17 {
			return
		}
		v0 := pkg1.Func14(IntConst0, IntConst1*pkg7.v11, func() type2 {
			v21, racyVar0 = v9.Func15(v1, v13.v14, v13.v15)
			return v0
		})
		if v0 != nil {
			v6 <- v0
			return
		}
		v16 <- v3
	})
	if v0 != nil {
		return nil, v0
	}
	return v2, nil
}