package skeleton

func (v4 *v15) func1(v12 *pkg99.v19, v2 type0, v11 v25.v32, v10 *pkg162.v0) (v28 pkg44.v23, v27 *pkg44.v30, v1 *v20.v31, racyVar0 type1) {
	defer func() {
	}()
	v8, v9 := pkg113.WithContext(v12.v26)
	Wrapper1.Go(func() type1 {
		defer close(v29)
		if v5 := v22.Func10(); v5 != nil {
			return v5
		}
		v17, racyVar0 = v4.v3.Func11(v12, pkg68.v24{v7: v2.v14, v18: v16}, v22.v33, v10)
		if v21 != nil {
			if v13 == nil {
				return v21
			}
			return nil
		}
		if v5 := func28(v6, v2); v5 != nil {
			return v5
		}
		select {
		case v29 <- v6:
		case <-v9.Done():
			return v9.Func29()
		}
		return nil
	})
	Wrapper1.Go(func() type1 {
		v27, racyVar0 = v4.func31(v12, v28, v2, v10)
		return v5
	})
	v5 = Wrapper1.Wait()
	return v28, v27, v1, v5
}