package skeleton

func func1(v9 pkg9.v13) (*pkg3.v2, type0) {
	if v0 != nil {
		return nil, v0
	}
	if v1 == nil {
		return nil, pkg7.v10
	}
	for _, v6 := range *v1 {
		for racyVar0, v5 := range v14 {
			v16.Func10(func() type0 {
				v12, v7 := pkg4.Func11(v9, v17.Func12(v5))
				if v7 == nil {
					v11[racyVar0] = v12
				}
				return v7
			}, func() {
			}, func(v7 type0) {
			})
		}
		if v7 := Wrapper1.Wait(); v7 != nil {
			continue
		}
	}
	return &pkg3.v2{v8: v15, v3: v4}, nil
}