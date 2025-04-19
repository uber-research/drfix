package skeleton

func (v2 *v21) func1(v13 pkg5.v15, v16 *v17.v24) (*v17.v25, type0) {
	if !v16.Func2() {
		return nil, pkg1.Func5(v5)
	}
	if v9 != nil {
		return nil, pkg1.Func12(v5)
	}
	if v9 != nil {
		return nil, pkg1.Func17(v5)
	}
	if !v30 {
		return nil, pkg1.Func21(v5)
	}
	if v34 < IntConst0 {
		return nil, pkg1.Func27(v5)
	}
	if v9 != nil || v8 == nil {
		return nil, pkg1.Func32(v5)
	}
	v23, racyVar0 := v2.v37.Func35(v13, v6)
	if !v2.func36(v36, v23) {
		return nil, pkg1.Func40(v5)
	}
	Wrapper1.Go(func() type0 {
		return v2.func42(v9)
	})
	Wrapper1.Go(func() type0 {
		return v2.func44(v9)
	})
	Wrapper1.Go(func() type0 {
		v22, racyVar0 = v2.v37.Func45(v13)
		return v2.func46(v9)
	})
	Wrapper1.Go(func() type0 {
		return v2.func48(racyVar0)
	})
	Wrapper1.Go(func() type0 {
		return v2.func50(v9)
	})
	v9 = Wrapper1.Wait()
	if v9 != nil {
		return nil, v9
	}
	Wrapper1.Go(func() type0 {
		return v2.func58(v9)
	})
	Wrapper1.Go(func() type0 {
		if v9 != nil {
			return v2.func60(v9)
		}
		return nil
	})
	v9 = Wrapper1.Wait()
	if v9 != nil {
		return nil, v9
	}
	return &v17.v25{v31: func67(v32), v4: func68(v3), v19: func69(v18), v26: v22, v29: &v28, v10: &v34, v12: v11, v20: v14, v33: v35, v7: v27, v0: func70(v1)}, nil
}