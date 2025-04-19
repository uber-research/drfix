package skeleton

func (v3 *type0) func1(v4 pkg5.v10, v1 pkg11.v15, v12 pkg1.v11) (*pkg1.v6, type1) {
	v18, v8 := errgroup.WithContext(v4)
	v17 := sync.Mutex{}
	for _, racyVar0 := range v14 {
		Wrapper1.Go(func() type1 {
			v7, v2 := v3.v9.Func3(v8, racyVar0)
			if v2 != nil {
				return v2
			}
			v17.Lock()
			defer v17.Unlock()
			for _, v13 := range v7.Func6().Func5() {
				if v0 != "StringConst0" && v0 != v13.Func8().v16.Func7() {
					return pkg3.Func9("StringConst1")
				}
			}
			return nil
		})
	}
	if v2 := Wrapper1.Wait(); v2 != nil {
		return &v5, v2
	}
	return &v5, nil
}