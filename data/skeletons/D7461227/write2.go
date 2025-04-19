package skeleton

func (v20 *v2) func1(v10 pkg9.v17, v6 []*pkg6.v5, v28 pkg6.v27, v4 chan *pkg6.v15) {
	var racyVar1 type0
	defer func() {
		if v9 != nil {
			for range v6 {
				v4 <- v21
			}
		}
	}()
	if len(v6) == IntConst0 {
		return
	}
	switch v28 {
	case pkg6.v11, pkg6.v3:
		if v22 != nil {
			return
		}
	case pkg6.v25, pkg6.v18:
		if v9 != nil {
			return
		}
	default:
		return
	}
	for _, v7 := range v6 {
		Wrapper1.Go(func(v7 *pkg6.v5, v24 pkg1.v29, v12 pkg1.v1) func() {
			return func() {
				v24.v13 = &v12
				v24.v13.v30, racyVar1 = func35(v7)
				v24.v13.v16 = v24.v13.v23.v0[v24.v13.v23.v19[IntConst7]]
				v20.func36(v10, v24, v28, v7, v14)
			}
		}(v7, v24, *v24.v13))
	}
	for range v6 {
		if v26.v8 != nil {
			v4 <- v21
			continue
		}
		v4 <- v26.v15
	}
}