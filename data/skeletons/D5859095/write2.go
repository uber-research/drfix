package skeleton

func (v13 type0) func1(v3 pkg4.v9, v10 pkg5.v4) (*pkg5.v2, type1) {
	defer v18.Func4()
	var racyVar0 type1
	if v10.v15 {
		v18.Func5(func() {
			return
		})
	}
	if v10.v16 {
		v18.Func7(func() {
			if v1 != nil {
				racyVar0 = v1
			} else {
				v11.v8 = v0.v8
			}
			return
		})
	}
	if v10.v14 {
		v18.Func9(func() {
			if v1 != nil {
				racyVar0 = v1
			} else {
				v11.v17 = v6.v17
			}
			return
		})
	}
	if v1 := Wrapper1.Wait(); v1 != nil {
		return nil, v12.Func11(v12.v5, v1)
	} else if v7 != nil {
		return nil, v7
	}
	return v11, nil
}