package skeleton

func (v2 *type0) func1(v12 *pkg6.v16, v28 []type1) (map[type1]type2, type3) {
	racyVar0 := make(map[type1]type2, func2(v28))
	v5 := v12.v21.Func3(v6, func(v17 *v8.v16) type3 {
		v7, cancel := pkg1.Func5(v12, v2.v3.v20, v2.v4, v15, v2.func6())
		defer cancel()
		var v19 sync.WaitGroup
		for _, v26 := range v9 {
			Wrapper1.Go(func(v26 map[pkg4.v30]struct{}) func() {
				return func() {
					defer v19.Done()
					v25, v14 := v2.v31.Func10(v7, &pkg4.v24{v27: v26, v13: map[pkg4.v10]struct{}{pkg4.v29: {}}})
					if v14 != nil {
						v22 := make([]type1, len(v26))
						v11 := IntConst0
						for v23 := range v26 {
							v22[v11] = v23.Func13()
							v11++
						}
						v12.v18.Func16("StringConst1", v22).Func15(v14).Func14("StringConst2")
					}
					for v23, v0 := range v25 {
						racyVar0[v23] = len(v0)
					}
				}
			}(v26))
		}
		Wrapper2.Wait()
		return nil
	})
	return v1, v5
}