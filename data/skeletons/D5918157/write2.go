package skeleton

func (v2 *v31) func1(v6 pkg8.v14, v9 *pkg3.v8) (*v15.v23, type0) {
	racyVar0 := make(map[pkg3.v26]type0)
	var v16 sync.WaitGroup
	Wrapper1.Go(func() {
		defer v16.Done()
		if pkg0.Func3(v9.v32, pkg3.v34) {
			if v5 != nil {
				racyVar0[v36.v37] = v5
				return
			}
		}
	})
	Wrapper1.Go(func() {
		defer v16.Done()
		if pkg0.Func9(v9.v32, pkg3.v27) || pkg0.Func10(v9.v29, []pkg3.v26{pkg3.v13, pkg3.v30, pkg3.v22}) {
			if v5 != nil {
				racyVar0[v36.v37] = v5
				return
			}
		}
	})
	Wrapper1.Go(func() {
		defer v16.Done()
		if pkg0.Func16(v9.v32, pkg3.v24) {
			if !pkg0.Func21(v3, v25) {
				return
			}
			if v5 != nil {
				return
			}
		}
	})
	Wrapper1.Go(func() {
		defer v16.Done()
		if pkg0.Func24(v9.v32, pkg3.v33) {
			if v5 != nil {
				return
			}
		}
	})
	Wrapper1.Go(func() {
		defer v16.Done()
		if pkg0.Func26(v9.v32, []pkg3.v26{pkg3.v1, pkg3.v18, pkg3.v11}) || pkg0.Func27(v9.v29, []pkg3.v26{pkg3.v35, pkg3.v0}) {
			if v5 != nil {
				return
			}
			if !pkg0.Func33([]type1{"StringConst4", "StringConst5", "StringConst6", "StringConst7", "StringConst8", "StringConst9"}, v9.v10) && v20.v4 != v9.v10 {
				return
			}
		}
	})
	Wrapper2.Wait()
	if len(v21) == IntConst1 {
		var v28 sync.WaitGroup
		Wrapper1.Go(func() {
			defer v28.Done()
			if v9.v10 == "StringConst11" && pkg0.Func37(v9.v29, []pkg3.v26{pkg3.v35, pkg3.v13, pkg3.v30}) {
				if v17.v19 == nil {
					if v5 != nil {
						return
					}
				}
				if v7 < IntConst3 {
					return
				}
				if v5 != nil {
					return
				}
			}
		})
		Wrapper1.Go(func() {
			defer v28.Done()
			if pkg0.Func49([]type1{"StringConst15", "StringConst16", "StringConst17"}, v9.v10) && pkg0.Func50(v9.v29, []pkg3.v26{pkg3.v12}) {
				if v17.v19 == nil {
					if v5 != nil {
						return
					}
				}
				if v7 < IntConst5 {
					return
				}
				if v5 != nil {
					return
				}
			}
		})
		Wrapper3.Wait()
	}
	return v17, pkg6.Func65(func66(v21)...)
}