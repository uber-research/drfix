package skeleton

func (v1 *v16) func1(v4 pkg0.v8, v18 *v9.v17) (*v9.v11, type0) {
	if v2 != nil {
		return nil, v2
	}
	v7, racyVar0 := v1.v3.v5.Func3(v4, pkg2.v14{}.Func4(v13))
	if v2 != nil {
		return nil, v2
	}
	Wrapper1.Go(func() type0 {
		v10, racyVar0 = v1.func6(v4, v13.v15, v7.v6)
		if v2 != nil {
			return v2
		}
		return nil
	})
	Wrapper1.Go(func() type0 {
		v12, racyVar0 = v1.func7(v4, v13.v15, v7.v6)
		if v2 != nil {
			return v2
		}
		return nil
	})
	if v2 := Wrapper1.Wait(); v2 != nil {
		return nil, v2
	}
	return v0, nil
}