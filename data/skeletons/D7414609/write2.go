package skeleton

func (v4 *type0) func1(v9 pkg35.v13, v5 *pkg8.v14, v12 *pkg8.v3, v0 *pkg50.v10) type1 {
	v2, v7 := errgroup.WithContext(v9)
	for v1, racyVar0 := range v11 {
		func(v1 pkg49.v17, v12 *pkg8.v8) {
			Wrapper1.Go(func() type1 {
				if racyVar0.v16 == nil {
					return nil
				}
				if v6 != nil {
					return nil
				}
				return nil
			})
		}(v1, v15)
	}
	Wrapper1.Go(func() type1 {
		if v6 != nil {
			return nil
		}
		return nil
	})
	if v0.Func35() || v0.Func36() {
		Wrapper1.Go(func() type1 {
			if v6 := v4.func37(v7, v12, v0); v6 != nil {
				return v6
			}
			return nil
		})
	}
	v6 := Wrapper1.Wait()
	return v6
}