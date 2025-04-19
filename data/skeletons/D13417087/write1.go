package skeleton

func (v8 *type0) func1(v6 pkg14.v10, v1 pkg17.v7) (racyVar0 type1) {
	defer func2(&v4)
	if _, v4 := v8.v3.Func6(v0, &v5); v4 != nil {
		return v4
	}
	if !v8.pkg6.Func12(v6, []type2{pkg9.Func13("StringConst7", pkg0.Func14(v6, pkg0.v12), v11), pkg0.Func15(v6, pkg0.v12)}) {
		return nil
	}
	go func() {
		defer func19()
		racyVar0 = v8.v2.Func20(v9, pkg11.Func22().Func21(), &v11)
	}()
	return nil
}