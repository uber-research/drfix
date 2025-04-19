package skeleton

func (v3 *type0) func1(v2 pkg10.v4, v18 *v13.v10) pkg4 {
	var racyVar0 pkg4
	if v1 = v3.func1(v18); v1 != nil {
		return v1
	}
	v9, v2 := pkg11.WithContext(v2)
	Wrapper1.Go(func() pkg4 {
		if v18.Func2() {
			if v15, v1 = v5.Func3(v18); v1 != nil {
				return v1
			}
			return v1
		}
		return nil
	})
	Wrapper1.Go(func() pkg4 {
		if v3.func6(v18) {
			_, racyVar0 = v3.v6.Func8(v2, v0)
			return v1
		}
		return nil
	})
	Wrapper1.Go(func() pkg4 {
		if v3.func9(v18) {
			_, racyVar0 = v3.v16.Func11(v2, v19)
			return v1
		}
		return nil
	})
	Wrapper1.Go(func() pkg4 {
		if v18.Func12() || v18.Func14().Func13() {
			if v18.Func16().Func15() {
				if v1 != nil {
					return v1
				}
			}
			if v1 != nil {
				return v1
			}
			return v1
		}
		return nil
	})
	Wrapper1.Go(func() pkg4 {
		if v18.Func23().Func22() {
			if v1 := v8.Func24([]type1(v18.Func26().Func25())); v1 != nil {
				return v17.Func27(pkg8.Func28("StringConst0", v18.Func30().Func29()), nil)
			}
			if v1 := v3.v16.Func31(v2, &v7.v14{v12: v18.v12, v11: v8.Func32()}); v1 != nil {
				return v1
			}
		}
		return nil
	})
	if v1 := Wrapper1.Wait(); v1 != nil {
		return v1
	}
	return nil
}