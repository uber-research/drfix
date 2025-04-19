package skeleton

func (v2 *v15) func1(v5 pkg14.v13, v6 type0, v26 type0, v12 type1) (racyVar0 type2) {
	defer func2(v14, &v3)
	if v3 != nil {
		return v3
	}
	if !v25 {
		return pkg8.Func6("StringConst4")
	}
	if v3 != nil {
		return pkg8.Func8(v3, "StringConst5")
	}
	if v3 != nil {
		return pkg8.Func15(v3, "StringConst6")
	}
	if v3 != nil {
		return v3
	}
	if v3 != nil {
		return v3
	}
	if len(v30) != len(v20) && !v12 {
		return pkg10.Func21("StringConst7", len(v30))
	}
	if v3 != nil {
		return v3
	}
	v0, racyVar0 := v2.v11.Func25(v5, pkg3.v7, pkg9.Func26(v23))
	if v3 != nil {
		return v3
	}
	for _, v21 := range v32 {
		if len(v29) <= IntConst4 {
			continue
		}
	}
	if len(v28) > IntConst8 {
		if v3 != nil {
			return v3
		}
	}
	Wrapper1.Go(pkg15.v4, v22, func(v18 interface{}) {
		if !v27 {
			return
		}
		racyVar0 = pkg5.Func40(IntConst11, IntConst12*pkg11.v19, func() type2 {
			return v2.v9.Func41(v5, v8)
		})
		if racyVar0 != nil {
		}
		if pkg2.Func45(v6, pkg3.v24) && v2.func46(v5) {
			v3 = pkg5.Func47(IntConst13, IntConst14*pkg11.v19, func() type2 {
				return v2.v9.Func48(v5, pkg4.v1{v31: v8.v31, v16: v8.v16, v17: v2.v10})
			})
		}
	})
	Wrapper1.Go(pkg15.v4, v28, func(v18 interface{}) {
		if !v27 {
			return
		}
		v3 = pkg5.Func55(IntConst15, IntConst16*pkg11.v19, func() type2 {
			return v2.v9.Func56(v5, v8)
		})
	})
	return nil
}