package skeleton

func (v19 *type0) func1(v7 pkg11.v15, v2 []type1, v12 map[type2]type2) (type3, type4) {
	defer func() {
	}()
	var racyVar0 type4
	v30, cancel := pkg11.Func19(v7, pkg10.Func20(v28)*pkg10.v1)
	defer cancel()
	switch v23 {
	case v8:
		go func(v9 pkg11.v15, v29 type5, v6 pkg8.v11, v31 []pkg15.v20, v26 map[type2]type2, v24 pkg9.v21, v16 *pkg7.v18) {
			defer func() {
			}()
		}(v7, v28, v10, v25, v14, v22, v17)
	case v3:
		go func() {
			defer func() {
			}()
			v30, cancel := pkg11.Func70(pkg11.Func71(), pkg10.Func72(v28)*pkg10.v1)
			defer cancel()
			_, racyVar0 = v19.v27.Func74(v30, &v10, v0)
			if v5 != nil {
				return
			}
		}()
	}
	if racyVar0 != nil {
		return v4, v5
	}
	return v13, nil
}