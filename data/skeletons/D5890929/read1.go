package skeleton

func (v2 *type0) func1(v7 pkg0.v10, v8 *pkg5.v18) type1 {
	if v5 != nil {
		return v0.Func2(v5, "StringConst0")
	}
	var v12 sync.WaitGroup
	go func(v7 pkg0.v10, v17 *pkg7.v19) {
		defer v12.Done()
		if v5 != nil {
			v15 <- v5
			return
		}
	}(v7, v8.v16)
	var racyVar0 []*pkg5.v22
	for _, v21 := range v11 {
		go func(v7 pkg0.v10, v3 pkg5.v20) {
			defer v12.Done()
			if v5 != nil {
				v15 <- v5
				return
			}
			racyVar0 = append(racyVar0, pkg4.Func10(v3.v6, v14))
		}(v7, v21)
	}
	Wrapper1.Wait()
	if len(v15) > IntConst1 {
		return v0.Func13(<-v15, "StringConst1"+"StringConst2", v8.v16, v8.v4, v8.v1, v8.v9)
	}
	return v2.v13.Func15(v7, v8)
}