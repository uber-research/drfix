package skeleton

func (v3 *v35) func1(v10 pkg11.v18, pkg11 *v20.v23, v34 *v20.v30) (v1 *v20.v36, racyVar0 type0) {
	if v27.v6 != nil && v27.v6.v0 != nil {
		if v8 = v9.Func1(v10, v10.Func2(), v27.v6.Func3(), v3.v24); v8 != nil {
			return
		}
	}
	if v8 != nil {
		return
	}
	if v8 != nil {
		return
	}
	if v8 = v3.func30(v17, v29.Func31()); v8 != nil {
		return
	}
	if v8 != nil {
		return
	}
	if pkg11.v16.v6 != nil && pkg11.v16.v6.v31 != nil {
		if v8 != nil {
			return
		}
	}
	v7, racyVar0 := v3.v22.Func49(v10, v29, pkg11.v16.v33)
	if v8 != nil {
		return
	}
	v3.v25.Func50(v10, v3.v19.Func51(), "StringConst6", func(v14 *pkg4.v13) type0 {
		return nil
	})
	v15 := v3.v25.Func58(v10, v3.v19.Func59(), "StringConst7", func(v14 *pkg4.v13) type0 {
		return v8
	})
	v11 := v3.v25.Func63(v26, v3.v19.Func64(), "StringConst9", func(v14 *pkg4.v13) type0 {
		v5, racyVar0 = v3.v12.Func65(v26, v29.Func66())
		return v8
	})
	if racyVar0 = <-v15; racyVar0 != nil {
		return
	}
	if v8 = <-v11; v8 != nil {
		return
	}
	if v29.Func82().Func81() {
		if v5.Func83() != v21.Func85().Func84() {
			if v8 != nil {
				return nil, v8
			}
		}
	}
	if v21.v6 == nil || v21.v6.v28 == nil || string(*v21.v6.v28) == "StringConst14" {
		return nil, pkg8.Func93("StringConst18", nil)
	}
	if pkg11.v16.Func94() && pkg11.v16.v6.Func95() {
		v3.v25.Func99(v10, v3.v19.Func100(), "StringConst20", func(v14 *pkg4.v13) type0 {
			return v3.v4.Func101(v10, v32)
		})
	}
	if v2 {
		v8 = v3.v25.Func109(v10, v3.v19.Func110(), "StringConst23", func(v14 *pkg4.v13) type0 {
			return v8
		})
	}
	return v1, nil
}