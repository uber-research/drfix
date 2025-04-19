package skeleton

func (v13 *type0) func1(v4 pkg14.v9, v3 type1, v14 type1) ([]v11, type2) {
	if v2 != nil {
		return nil, pkg9.Func2(v2, "StringConst0")
	}
	var v6 sync.Mutex
	racyVar0 := v3
	for _, v0 := range v12 {
		if v21[v10] != "StringConst1" {
			if v2 != nil {
				return nil, pkg9.Func6(v2, "StringConst2")
			}
		}
		if v17 > v14 {
			break
		}
		if v17 >= v3 {
			Wrapper1.Go(func() type2 {
				if v2 != nil {
					return pkg9.Func9(v2, "StringConst3")
				}
				if v2 != nil {
					return pkg9.Func11(v2, "StringConst4")
				}
				v6.Lock()
				for _, v18 := range v5.v16 {
					if v18.v20 != racyVar0 {
						racyVar0 = v18.v20
					}
				}
				v6.Unlock()
				return nil
			})
		}
	}
	v2 = Wrapper1.Wait()
	if v2 != nil {
		return nil, pkg9.Func14(v2, "StringConst6")
	}
	pkg18.Func15(v1, func(v7, v8 type3) type4 {
		return v1[v7].v15 < v1[v8].v15
	})
	return v19, nil
}