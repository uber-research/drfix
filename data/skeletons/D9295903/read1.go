package skeleton

func (v8 *v5) func1(v10 map[pkg3.v12]*pkg2.v0, v2 map[pkg3.v12][]*pkg2.v15) (v4, v6 []*pkg2.v0) {
	defer func() {
	}()
	for racyVar0, racyVar1 := range v10 {
		if v13 {
			for _, v14 := range v11 {
				if v3 := v1.Func3(v14); v3 != nil {
					continue
				}
			}
		}
		if !v7 && v1.v9 > IntConst0 {
			Wrapper1.Go(func() {
				v8.func5(racyVar0, racyVar1)
			})
			continue
		}
	}
	return
}