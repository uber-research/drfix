package skeleton

func (v1 *v23) func1(v6 pkg7.v12, v24 *pkg1.v9) (*pkg1.v11, type0) {
	v22, racyVar0 := v1.v18.Func1(v6, v24.v4)
	if v2 != nil {
		return nil, v2
	}
	Wrapper1.Go(func() type0 {
		if v22.v20 == pkg1.v10 {
		} else {
			v19, racyVar0 = v1.v16.Func8(v6, v22.v5, &pkg1.v14{v13: v8.v13, v7: v21.v7})
		}
		return v2
	})
	Wrapper1.Go(func() type0 {
		return v2
	})
	Wrapper1.Go(func() type0 {
		v0, racyVar0 = v1.v17.Func10(v6, v22, v21)
		return v2
	})
	if v2 := Wrapper1.Wait(); v2 != nil {
		return nil, v2
	}
	return func16(v6, v3, v15), nil
}