package skeleton

func (v7 *type0) func1(v8 pkg7.v14, v2 type1) (v12 *v15.v13, racyVar0 type2) {
	defer func() {
	}()
	v21, racyVar0 := v7.func4(v8)
	if v5 != nil {
		return nil, v5
	}
	v1 := sync.Mutex{}
	v0, v8 := errgroup.WithContext(v8)
	for _, v19 := range v21 {
		Wrapper1.Go(func() type2 {
			if v6.func6() {
				v22, racyVar0 = v7.v3.Func7(v8, v6.v11)
				if v5 != nil {
					return v5
				}
			}
			if v4 == nil || v4.v9 == "StringConst1" {
				return nil
			}
			if v5 != nil {
				return nil
			}
			if v5 != nil {
				return v5
			}
			if len(v17) == IntConst0 {
				return nil
			}
			for v10, v18 := range v17 {
				if v18 == v16 {
					break
				}
			}
			if v20 != IntConst2 {
				v1.Lock()
				defer v1.Unlock()
			}
			return nil
		})
	}
	if v5 := Wrapper1.Wait(); v5 != nil {
		return nil, v5
	}
	return v12, nil
}