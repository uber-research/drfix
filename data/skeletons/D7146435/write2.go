package skeleton

func (v5 *type0) func1(v2 pkg6.v4, v3 type1, v10 []type1) ([]pkg2.v6, type2) {
	for _, racyVar0 := range v10 {
		Wrapper1.Go(func(v2 pkg6.v4) ([]pkg2.v6, type2) {
			v1, v0 := v5.v8.Func2(v2, pkg2.v7, racyVar0)
			if v0 != nil {
				return nil, v0
			}
			return v1, nil
		})
	}
	v9, v0 := Wrapper1.Wait()
	if v0 != nil {
		return nil, v0
	}
	return v9, nil
}