package skeleton

func (v1 *v16) func1(v6 pkg1.v8) type0 {
	if v3 != nil {
		return v3
	}
	var racyVar0 *type1
	for v10 := IntConst0; v10 < v7 && (v4 == nil || *v4 >= v10); {
		var v12 sync.WaitGroup
		for v5 := IntConst1; v5 < v14; v5++ {
			if (v4 != nil && *v4 < v10) || v10 > v7 {
				break
			}
			Wrapper1.Go(func(v15 type1) func() {
				return func() {
					defer v12.Done()
					v0, v9, v3 := v1.func10(v6, v13, v15)
					if v9 && (racyVar0 == nil || *racyVar0 > v15) {
						racyVar0 = &v15
					}
					if v3 != nil {
						v1.v2.Func12("StringConst7").Func11(IntConst3)
						v1.v11.Func13("StringConst8", pkg7.Func14(v3))
						return
					}
					if v0 != nil {
						v3 = v1.func15(v6, []type2{*v0})
					}
					if v3 != nil {
						v1.v2.Func17("StringConst9").Func16(IntConst4)
						v1.v11.Func18("StringConst10", pkg7.Func19(v3))
						return
					}
					v1.v2.Func21("StringConst11").Func20(IntConst5)
				}
			}(v10))
		}
		Wrapper2.Wait()
	}
	return nil
}