package skeleton

func (v8 *v4) func1(v2 pkg4.v6, v5 []*pkg1.v1) ([]*pkg1.v9, type0) {
	v7, v11 := errgroup.WithContext(v2)
	for _, racyVar0 := range v5 {
		Wrapper1.Go(func() type0 {
			v10, v0 := v8.Func2(v11, racyVar0)
			if v0 != nil {
				return v0
			}
			for _, v12 := range v10 {
				v3 <- v12
			}
			return nil
		})
	}
	v0 := Wrapper1.Wait()
	if v0 != nil {
		return nil, v0
	}
	return v10, nil
}