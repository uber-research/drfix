package skeleton

func (v0 *type0) func1(v1 pkg0.v2, v3 []*pkg3.v4, v5 type1) type2 {
	for _, racyVar0 := range v3 {
		v6 <- struct{}{}
		Wrapper1.Go(func() {
			v0.v7.Func2(v1, racyVar0)
		})
	}
	return nil
}