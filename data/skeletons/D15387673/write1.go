package skeleton

func (v1 *type0) func1() type1 {
	defer v1.v7.Func6(v10).Func5("StringConst3", v2).Func4().Func3()
	defer v1.v7.Func9("StringConst4", v2).Func8().Func7()
	for _, v14 := range v1.v6.v8 {
		if v14.v12 == nil || v14.v12 == v4.v5.Func14() {
			v14.racyVar0 = v4.v3.Func15()
		}
	}
	v0 := pkg3.Func22(pkg22.Func23(), pkg3.Func24(func() (v0 type1) {
		return v0
	}))
	if v0 != nil || v13 == nil {
		return nil
	}
	defer func() {
	}()
	for _, v14 := range v9 {
		if v1.func42(v14, v10) {
			continue
		}
		if v11 == nil {
			continue
		}
	}
	return nil
}