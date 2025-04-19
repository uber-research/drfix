package skeleton

func (v2 *type0) func1(v3 pkg9.v7, v10 *pkg0.v8) (v1 type1) {
	defer func() {
	}()
	v11, v3 := errgroup.WithContext(v3)
	racyVar0 := v0
	Wrapper1.Go(func() type1 {
		for !racyVar0 {
		}
		return nil
	})
	Wrapper1.Go(func() type1 {
		racyVar0 = v6
		if v1 != nil {
			return pkg5.Func36(pkg5.v5, v1).Func35("StringConst12", v9).Func34("StringConst13", v4)
		}
		return nil
	})
	v1 = Wrapper1.Wait()
	if v1 != nil {
		return v1
	}
	return nil
}