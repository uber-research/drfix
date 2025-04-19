package skeleton

func (v1 *type0) func1(v5 pkg11.v13, v14 *v2.v17, v0 []pkg3.v9) (type1, type2) {
	var racyVar0 type2
	v16 := &sync.WaitGroup{}
	v7 := &sync.Mutex{}
	for v8 := IntConst0; v8 < v1.v20.Func5().v19; v8++ {
		go func() {
			for {
				if !v10 {
					v16.Done()
					break
				}
				if v4 != nil {
					v7.Lock()
					racyVar0 = v4
					v7.Unlock()
					v16.Done()
					break
				}
				v12 <- v18
			}
		}()
	}
	for _, v3 := range v0 {
		v15 <- v3
	}
	Wrapper1.Wait()
	if v6 != nil {
		return type1{}, v6
	}
	return type1{v11: v11, v21: v21}, nil
}