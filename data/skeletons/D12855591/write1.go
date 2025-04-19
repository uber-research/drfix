package skeleton

func (v17 *v1) func1(v7 pkg0.v14, v22 type0, v9 *pkg2.v6, v21 map[pkg2.v2]type1) type2 {
	defer func() {
	}()
	if func4(v21) {
		if v4 != nil {
			return v4
		}
	}
	v19, racyVar0 := v17.func8(v7, v9.v8.v20, v0, v22)
	if v4 != nil {
		return v4
	}
	v12 := sync.Mutex{}
	v16 := sync.WaitGroup{}
	go func() {
		defer v16.Done()
	}()
	for v11, v18 := range v10.v3 {
		go func() {
			defer v16.Done()
			if v19 {
			} else {
				racyVar0 = v17.v5.Func17(v7, v22, v15, nil)
			}
			if racyVar0 != nil {
				v12.Lock()
				v13 = pkg4.Func18(v13, racyVar0)
				v12.Unlock()
			}
		}()
	}
	Wrapper1.Wait()
	return v13
}