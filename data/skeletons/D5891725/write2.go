package skeleton

func (v16 type0) func1(v4 pkg5.v9, v10 v5.v11) (v13 *v5.v0, racyVar0 type1) {
	defer func() {
	}()
	defer v19.Func16()
	v12 := &sync.Mutex{}
	for v6 := IntConst0; v6 < v1; v6++ {
		v19.Func18(func() {
			if v14 != nil {
				racyVar0 = v14
			} else {
				v12.Lock()
				v18 = append(v18, v2.v17...)
				v12.Unlock()
			}
			return
		})
	}
	if v3 := Wrapper1.Wait(); v3 != nil {
		return nil, v15.Func23(v15.v7, v3)
	} else if v8 != nil {
		return nil, v8
	}
	return v13, nil
}