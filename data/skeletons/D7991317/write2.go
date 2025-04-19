package skeleton

func (v0 *type0) func1(v5 pkg30.v8, v4 type1, v6 []type2) type3 {
	defer func2(v9, nil)
	v1, racyVar0 := func3(v5, v4, v6)
	if v2 != nil {
		return v2
	}
	if len(v1) == IntConst0 {
		return nil
	}
	Wrapper1.Go(pkg41.v3, v1, func(v10 interface{}) {
		if !v12 {
			return
		}
		racyVar0 = pkg9.Func10(IntConst1, IntConst2*pkg25.v11, func() type3 {
			return v0.pkg18.Func11(v5, *v7)
		})
		if v2 != nil {
			return
		}
	})
	return nil
}