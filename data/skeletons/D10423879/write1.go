package skeleton

func (v3 *v20) func1(v6 pkg13.v11, v21 pkg8.v22, v0 v10) ([]pkg8.v2, type0) {
	v23, racyVar0 := v3.func1(v6, v21)
	if v4 != nil {
		return nil, pkg6.Func2(v4, "StringConst0")
	}
	v24, v7 := pkg14.WithContext(v6)
	if len(v0.v17) == IntConst1 || len(v18.v17) > IntConst2 {
		Wrapper1.Go(func() type0 {
			v8, racyVar0 = v3.v15.Func10(v7, v21.v14, v23, v18)
			if v4 != nil && pkg6.Func11(v4) != pkg6.v19 {
				return pkg6.Func12(v4, "StringConst1")
			}
			return nil
		})
	}
	if len(v0.v17) == IntConst3 || len(v5.v17) > IntConst4 {
		Wrapper1.Go(func() type0 {
			v16, racyVar0 = v3.v9.Func15(v7, v21.v14, v23, v5)
			if v4 != nil && pkg6.Func16(v4) != pkg6.v19 {
				return pkg6.Func17(v4, "StringConst2")
			}
			return nil
		})
	}
	if v4 := Wrapper1.Wait(); v4 != nil {
		return nil, v4
	}
	if len(v1) == IntConst5 {
		return nil, pkg6.Func20("StringConst3")
	}
	if v0.v13 {
		if v4 != nil {
			return nil, v4
		}
	}
	if v0.v12 == nil {
		return v1, nil
	}
	return v3.func22(v1, v0.v12)
}