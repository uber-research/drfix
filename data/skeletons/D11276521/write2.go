package skeleton

func (v1 *pkg21) func1(v5 pkg13.v14, v12 []pkg2.v28) (v26 []pkg2.v4, v2 type0) {
	racyVar0 := v13.v25
	var v15 sync.WaitGroup
	for _, v3 := range v12 {
		go func(v7 pkg2.v28) {
			defer v15.Done()
			if v7.v9 != nil && pkg2.v27 == *v7.v9 {
				racyVar0 = v13.v24
			}
			if v2 != nil {
				v23 <- v2
			} else {
				v20, v2 := pkg0.Func21(v7.v22)
				if v2 != nil {
					v23 <- v2
					return
				}
				v0 := func22(v11, v7.v9, v20)
				v10 <- v0
			}
		}(v3)
	}
	Wrapper1.Wait()
	if v16 > IntConst1 {
		if len(v10) > IntConst2 {
		} else {
			v21[v17] = "StringConst9"
			v21[v19] = "StringConst10"
			v18.Func41(v21).Func40(v6).Func39(IntConst3)
			return []pkg2.v4{}, pkg6.Func42(pkg6.v8, v2, "StringConst11")
		}
	}
	return v26, v2
}