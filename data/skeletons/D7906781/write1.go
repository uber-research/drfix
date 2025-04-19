package skeleton

func (v2 *v14) func1(v7 pkg12.v8, v10 v4.v9) (v4.v1, type0) {
	if v5 != nil {
		return v0, v5
	}
	v3, racyVar0 := v2.v12.Func7(v7)
	if v5 != nil {
		return v0, v5
	}
	Wrapper1.Go(func() {
		v7, cancel := pkg12.Func12(pkg12.Func13(), pkg8.v11*IntConst0)
		defer cancel()
		v2.func14(v7, v13, v10, v3, racyVar0)
	})
	if v5 != nil {
		return v0, v5
	}
	if v2.func15(v7) {
		if racyVar0 = v2.func16(v7, v10.v6); racyVar0 != nil {
		}
	}
	return v0, nil
}