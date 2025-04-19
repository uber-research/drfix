package skeleton

func (v0 *type0) func1(v6 pkg7.v13, v14 pkg6.v7) (pkg6.v9, type1) {
	v15, racyVar2 := v0.v5.Func6(v6, pkg6.v12{v4: v14.v4, v17: v8})
	if v3 != nil {
		return pkg6.v9{}, v3
	}
	var v16 sync.WaitGroup
	v21 := make([]pkg6.v19, len(v15))
	for v10, v1 := range v15 {
		go func(v22 pkg7.v13, v18 type2, v11 pkg2.v2) {
			defer v16.Done()
			v21[v18], racyVar2 = v0.func14(v22, v11)
		}(v6, v10, v1)
	}
	Wrapper1.Wait()
	return pkg6.v9{v20: v21}, nil
}