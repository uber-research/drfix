package skeleton

func (v4 *type0) func1(v12 pkg46.v23, v44 pkg12.v11) (pkg12.v35, type1) {
	if v7 != nil {
		return pkg12.v35{}, v7
	}
	if !pkg5.Func17() {
		return pkg12.v35{}, pkg32.Func18("StringConst7")
	}
	if v7 := v4.v14.Func19(v12, pkg12.v17{v40: v44.v40, v10: v44.v42, v16: v44.v16, v6: pkg12.v36}); v7 != nil {
		return pkg12.v35{}, pkg32.Func20(pkg37.Func21("StringConst8", v7.Func22()))
	}
	if v7 != nil {
		return pkg12.v35{}, v7
	}
	if !v13.v3 {
		return pkg12.v35{}, pkg32.Func24("StringConst9")
	}
	if v44.v1.Func25(pkg12.v31) && v13.v37 && len(pkg5.v2.Func27()) == IntConst1 {
		return pkg12.v35{}, pkg32.Func28("StringConst10")
	}
	if v48, _ := v4.v47.Func29(v12, pkg45.v8, pkg7.Func30(v44.v40.Func31()), v5); v48 && v44.v43 != nil {
		if v7 != nil {
			if v19 != nil || v15.v41.v38 == "StringConst11" {
				return pkg12.v35{}, v7
			}
			return pkg12.v35{}, v7
		}
		if len(v15.v38) == IntConst2 {
			return pkg12.v35{}, pkg32.Func41("StringConst15")
		}
		if v15.v41.v26 <= IntConst3 || v15.v41.v26 > IntConst4 || v15.v41.v25.v38 != pkg52.v9.Func42() {
			return pkg12.v35{}, pkg32.Func46("StringConst16")
		}
	}
	if v7 != nil {
		return pkg12.v35{}, v7
	}
	if v7 != nil {
		return pkg12.v35{}, v7
	}
	if v7 != nil {
		return pkg12.v35{}, v7
	}
	racyVar0 := sync.WaitGroup{}
	go func() {
		racyVar0.Add(IntConst6)
		defer v24.Done()
	}()
	if v44.v42 != nil && *v44.v42 != "StringConst24" {
		go func() {
			defer v24.Done()
		}()
	}
	for _, v28 := range v20.v39 {
		go func(v21 pkg12.v32) {
			defer v24.Done()
			v22, cancel := pkg46.Func84(v34, IntConst9*pkg39.v29)
			defer cancel()
		}(v30)
	}
	if v44.v1.Func90(pkg12.v31) {
		go func() {
			defer v24.Done()
		}()
		go func() {
			defer v24.Done()
		}()
		go func() {
			defer v24.Done()
		}()
		go func() {
			defer v24.Done()
		}()
		go func() {
			defer v24.Done()
		}()
		if len(v20.v46) > IntConst15 && v20.v46[IntConst16].v45 != pkg65.v27 {
			go func() {
				defer v24.Done()
			}()
		}
		if pkg17.v33 == nil || pkg17.v18 == nil {
			go func() {
				defer v24.Done()
			}()
		}
		go func() {
			defer v24.Done()
		}()
		go func() {
			defer v24.Done()
		}()
		go func() {
			defer v24.Done()
		}()
		if v13 != nil && v13.v37 {
			go func() {
				defer v24.Done()
			}()
			go func() {
				defer v24.Done()
			}()
		}
	}
	go func() {
		defer v24.Done()
	}()
	if pkg5.Func130() {
		go func() {
			defer v24.Done()
		}()
	}
	Wrapper1.Wait()
	return v0, nil
}