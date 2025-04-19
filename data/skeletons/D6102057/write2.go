package skeleton

func (v28 *type0) func1(v12 pkg21.v23, v24 *pkg14.v21) (*pkg14.v2, type1) {
	racyVar0, v8 := v28.v31.Func6(v12, v24.v9)
	if v8 != nil {
		return nil, v27.Func9(v8, v3, "StringConst9", "StringConst10", v13)
	}
	if v20 == nil {
		if v8 != nil {
			return nil, v22
		}
		if v20 != nil {
			if v8 != nil {
				return nil, v22
			}
		}
	}
	if v20 == nil {
		return nil, v22
	}
	if v20.v10 {
		return &pkg14.v2{}, nil
	}
	if v7 {
		v30.Func19(func() {
			v28.v37.Func20(&v29.v25{v11: v20.v11, v17: v20.v17, v35: pkg7.Func21(int(racyVar0.v4)), v18: v20.v18, v34: v20.v34, v36: v5, v33: v20.v33, v6: v5, v9: v20.v9})
		}, v3)
	}
	if v8 != nil {
		return nil, &pkg14.v32{}
	}
	if v8 != nil {
		return nil, &pkg14.v32{}
	}
	if v7 {
		if v8 != nil {
			return nil, v19
		}
		if v26.v16 >= pkg27.v0 {
			if v8 != nil {
				return nil, v27.Func35(v8, v3, "StringConst42", "StringConst43", v13)
			}
		}
	} else {
		v28.v38.Func36(v1.v14{v15: "StringConst44", v11: v20.v11}, "StringConst45", v20.v9, "StringConst46", v20.v34, "StringConst47", v20.v4, "StringConst48", v20.v18, "StringConst49", v20.v17)
	}
	racyVar0, v8 = v28.v31.Func37(v12, racyVar0, v7)
	if v8 != nil {
		return nil, v27.Func38(v8, v3, "StringConst50", "StringConst51", v13)
	}
	return &pkg14.v2{}, nil
}