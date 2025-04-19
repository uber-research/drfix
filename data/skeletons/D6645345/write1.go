package skeleton

func (v1 *v27) func1(v8 pkg0.v18, v19 *v9.v12) (*v9.v5, type0) {
	defer v1.v25.Func2(v8, v1.v20, v14.v2).Func1(func(v8 pkg0.v18, v10 *pkg1.v16) {
		v10.v26[v14.v13] = v19.Func3()
		if v19.Func4() != nil {
			v10.v26[v14.v15] = v19.Func6().Func5()
			v10.v26[v14.v3] = v19.Func8().Func7()
		}
		v10.v23 = v17
		v10.v6 = v7
	})
	if v7 != nil {
		return nil, v7
	}
	racyVar0 := map[type1]*v9.v11{}
	v21 := sync.WaitGroup{}
	for _, v28 := range v24 {
		go func(v28 *v9.v4) {
			defer v21.Done()
			racyVar0[v28.Func20()] = v22
		}(v28)
	}
	Wrapper1.Wait()
	return v0, nil
}