package skeleton

func (v11 *type0) func1(v3 pkg7.v8, v15 *v10.v6) (v0 *v10.v12, v2 type1) {
	if v2 != nil {
		return nil, v2
	}
	if v2 != nil {
		return nil, v2
	}
	if v2 != nil {
		return nil, v2
	}
	v17 := &sync.Mutex{}
	if v13 != nil && len(v13.v14) > IntConst0 {
		for v16, racyVar0 := range v9 {
			if v16 != v1 {
				Wrapper1.Go(func() type1 {
					v5 := v11.func18(v4, racyVar0[IntConst1])
					if v2 != nil {
						return v2
					}
					v17.Lock()
					v17.Unlock()
					return nil
				})
			}
		}
	}
	if v2 := Wrapper1.Wait(); v2 != nil {
		return nil, v2
	}
	if v2 != nil {
		return nil, v2
	}
	return v7, v2
}