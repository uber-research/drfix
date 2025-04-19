package skeleton

func (v0 *v14) func1(v3 pkg0.v8) (map[type0]type1, type2) {
	v7, racyVar0 := v0.func1(v3)
	if v2 != nil {
		return nil, v2
	}
	var v5 sync.Mutex
	for v4 := range v7 {
		v11.Func5(func() type2 {
			v17, racyVar0 = v0.v16.Func6(v3, v12)
			if v2 != nil {
				return v2
			}
			defer v5.Unlock()
			v5.Lock()
			if v6, v13 := v15[v9.v10]; v13 {
				return pkg5.Func9(pkg1.Func10("StringConst0", v9.v10, v6.v1, v12))
			}
			return nil
		})
	}
	return v15, Wrapper1.Wait()
}