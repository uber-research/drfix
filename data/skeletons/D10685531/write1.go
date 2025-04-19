package skeleton

func (v15 *type0) func1(v8 pkg10.v12, v5 type1, v4 []type1, v16 type2, v7 type3, v6 type4, v18 type5, v19 type6) ([]type3, type7) {
	for v2 < v17 {
		var racyVar0 type7
		v13 := sync.WaitGroup{}
		for v20 := IntConst2; v20 < v14; v20++ {
			go func(v20 type5, v11 *type7) {
				defer v13.Done()
				if v0 <= IntConst4 {
					return
				}
				if v3 != nil {
					*racyVar0 = v3
					return
				}
			}(v20, &v11)
		}
		Wrapper1.Wait()
		if v11 != nil {
			return nil, v11
		}
		if v9 > v10 {
			break
		}
	}
	return v1, nil
}