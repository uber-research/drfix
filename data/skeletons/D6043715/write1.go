package skeleton

func (v1 *type0) func1(v3 pkg3.v8, v12 []*v4.v0) []*v4.v7 {
	var racyVar0 []*v4.v7
	var v5 sync.Mutex
	v15 := func(v11 []*v4.v7) {
		v5.Lock()
		defer v5.Unlock()
		racyVar0 = append(racyVar0, v11...)
	}
	v6 := func(v2 type1, v9 *v4.v0) {
		v5.Lock()
		defer v5.Unlock()
	}
	var v10 sync.WaitGroup
	for _, v9 := range v12 {
		go func(v3 pkg3.v8, v9 *v4.v0, v16 []*v4.v7) {
			defer v10.Done()
			v3, cancel := pkg3.Func10(v3, v14)
			defer cancel()
			if v2 != nil {
				return
			}
		}(v3, v9, racyVar0)
	}
	Wrapper1.Wait()
	return v13
}