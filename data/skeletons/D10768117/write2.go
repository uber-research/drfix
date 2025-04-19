package skeleton

func (v1 *type0) func1(v5 pkg7.v11, v14 *pkg1.v2) (racyVar0 type1) {
	defer func() {
	}()
	if v14 == nil || v14.v10 == nil {
		return
	}
	v12, racyVar0 := v1.v6.Func6(v5, v4.v13{*v9})
	if v3 != nil {
		return
	}
	if !v15 {
		return nil
	}
	v7, cancel := pkg7.Func14(pkg7.Func15(), pkg6.Func17().Func16(v8))
	Wrapper1.Go(func() {
		defer cancel()
		racyVar0 = v1.v0.Func18(v7, v14, v6)
	})
	return nil
}