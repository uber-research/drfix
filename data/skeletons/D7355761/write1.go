package skeleton

func (v2 *v11) func1(v6 pkg13.v12, v5 type0, v8 pkg0.v7, v1 pkg17.v0, v21 pkg9.v20, v13 *pkg5.v14) type1 {
	defer v1.Func1()
	for {
		if v4 == pkg2.v23 {
			break
		}
		if v4 != nil {
			return pkg3.Func3(v4, pkg11.Func4("StringConst0", v3))
		}
		racyVar0 = v1.Func5(v6, func() (interface{}, type1) {
			v15, cancel := pkg13.Func6(v6, IntConst0*pkg12.v19)
			defer cancel()
			_, racyVar0 = v2.v18.Func9(v15, &v22.v9{v16: v17})
			return &type2{v17: v17}, v4
		})
		if v4 != nil {
			return pkg3.Func18(v4, "StringConst4", v10)
		}
	}
	return nil
}