package skeleton

func (v4 *type0) func1(v5 pkg9.v11, v14 *pkg0.v12) (v0 *pkg0.v1, v3 type1) {
	defer func() {
	}()
	var v7 sync.Mutex
	v15, v5 := errgroup.WithContext(v5)
	racyVar0 := v2
	Wrapper1.Go(func() type1 {
		v8 := racyVar0
		for !v8 {
			v7.Lock()
			v7.Unlock()
		}
		return nil
	})
	Wrapper1.Go(func() type1 {
		v7.Lock()
		racyVar0 = v10
		v7.Unlock()
		if v3 != nil {
			return pkg5.Func36(pkg5.v9, v3).Func35("StringConst12", v13).Func34("StringConst13", v6)
		}
		return nil
	})
	v3 = Wrapper1.Wait()
	if v3 != nil {
		return nil, v3
	}
	return &pkg0.v1{}, nil
}