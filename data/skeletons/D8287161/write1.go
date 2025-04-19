package skeleton

func (v6 *type0) func1(v5 *pkg4.v0) (v12 []*pkg4.v11, racyVar0 type1) {
	Wrapper1.Go(func() type1 {
		v7, racyVar0 = v6.v4.Func4().Func3(v6.v2, v8)
		return v1
	})
	Wrapper1.Go(func() type1 {
		if v10 != nil {
			v3, racyVar0 = v6.v4.Func6().Func5(v6.v2, v10)
			return v1
		}
		return nil
	})
	v9 := Wrapper1.Wait()
	if v9 != nil {
		return []*pkg4.v11{}, v1
	}
	return append(v7, v3...), nil
}