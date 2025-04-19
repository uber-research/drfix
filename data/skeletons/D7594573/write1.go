package skeleton

func (v10 *type0) func1(v4 pkg11.v9, v2 *type1, v0 type2) (racyVar0 type3) {
	v13, v4 := pkg16.WithContext(v4)
	Wrapper1.Go(func() type3 {
		v7, racyVar0 = v10.v5.v6.Func2(v4, v12, pkg5.v8)
		return v1
	})
	Wrapper1.Go(func() type3 {
		v3, racyVar0 = v10.v5.v6.Func3(v4, v12, pkg5.v11)
		return v1
	})
	if v1 := Wrapper1.Wait(); v1 != nil {
		return nil
	}
	return nil
}