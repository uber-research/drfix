package skeleton

func (v2 *type0) func1(v5 pkg6.v8, v4 pkg11.v7) (map[type1][]*v11.v1, type2) {
	v13, v5 := pkg8.WithContext(v5)
	racyVar0 := make(map[type1][]*v11.v1)
	for _, v9 := range v2.v6.v10 {
		func(v9 type1) {
			Wrapper1.Go(func() type2 {
				if v3 != nil {
					return v3
				}
				racyVar0[v9] = v0
				return nil
			})
		}(v9)
	}
	if v3 := Wrapper1.Wait(); v3 != nil {
		return nil, v3
	}
	return v12, nil
}