package skeleton

func (v9 *type0) func1(v3 *pkg16.v4, v2 []*pkg16.v4, v8 pkg9.v6) (pkg9.v6, type1) {
	var racyVar0 type1
	Wrapper1.Go(func() type1 {
		v5, racyVar0 = v9.Func2(v3, v2, v8)
		return v1
	})
	Wrapper1.Go(func() type1 {
		v7, racyVar0 = pkg0.Func3(v3, v9.v10, v9.v0, v8)
		return v1
	})
	v1 = Wrapper1.Wait()
	return v5, v1
}