package skeleton

func (v2 *v37) func1(v9 pkg8.v22, v23 pkg0.v26) ([]pkg0.v6, type0) {
	if v23.v32 == nil || len(v23.v32) == IntConst0 {
		return nil, nil
	}
	if v4 != nil {
		return nil, v4
	}
	if v4 != nil {
		return nil, v4
	}
	if v4 != nil {
		return nil, v4
	}
	v24, racyVar0 := v2.v36.Func10(v9, *v1.v34, v12, v3)
	if v4 != nil {
		return nil, v4
	}
	for v11, v28 := range v23.v32 {
		v15.Func15(func() type0 {
			if v28.v13 != nil && v28.v14 != nil {
				if v4 != nil {
					return v4
				}
			} else {
				v29.v0 = pkg6.Func21(pkg7.v7)
			}
			if v28.v10 != nil {
				v19, racyVar0 = v2.v25.Func22(v9, pkg5.Func23(v28.v10))
				if racyVar0 != nil {
				}
			}
			if v33 {
				if *v29.v0 == pkg7.v7 {
				} else if *v29.v0 == pkg7.v20 {
					v21, v4 := v2.v30.Func37(v9, pkg0.v27{v13: *v28.v13, v31: pkg6.Func38(v3)})
					if v4 != nil {
						return v4
					}
					if v21.v18 != nil {
						v29.v0 = pkg6.Func39(pkg7.v8)
					}
				}
			}
			return nil
		})
	}
	if v4 := Wrapper1.Wait(); v4 != nil {
		if v16, v35 := v4.(pkg2.v17); v35 && func42(v16) > IntConst2 {
			return nil, v16[IntConst3]
		}
	}
	return v5, nil
}