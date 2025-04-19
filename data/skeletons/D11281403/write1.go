package skeleton

func (v23 *type0) func1() {
	v19 := []type1{{v9: "StringConst0", v29: &v21.v18{}, v22: IntConst0, v35: func(v0 *v21.v18, v7 type2, v10 type3) (*v21.v18, type2) {
		pkg7.Func1(pkg7.Func2(v10) * pkg7.v1)
		return v0, v7
	}}, {v9: "StringConst1", v32: pkg6.Func3("StringConst2"), v22: IntConst1, v35: func(v0 *v21.v18, v7 type2, v10 type3) (*v21.v18, type2) {
		pkg7.Func4(pkg7.Func5(v10) * pkg7.v1)
		return v0, v7
	}}, {v9: "StringConst3", v29: &v21.v18{v6: []v21.v27{{v13: v21.v13{v20: v21.v31}}}}, v15: v12, v22: IntConst2, v35: func(v0 *v21.v18, v7 type2, v10 type3) (*v21.v18, type2) {
		pkg7.Func6(pkg7.Func7(v10) * pkg7.v1)
		return v0, v7
	}}, {v9: "StringConst4", v14: pkg6.Func8("StringConst5"), v22: IntConst3, v35: func(v0 *v21.v18, v7 type2, v10 type3) (*v21.v18, type2) {
		pkg7.Func9(pkg7.Func10(v10) * pkg7.v1)
		return v0, v7
	}}, {v9: "StringConst6", v22: IntConst4, v26: IntConst5, v35: func(v0 *v21.v18, v7 type2, v10 type3) (*v21.v18, type2) {
		pkg7.Func11(pkg7.Func12(v10) * pkg7.v1)
		return v0, v7
	}}, {v9: "StringConst7", v15: v2, v22: IntConst6, v35: func(v0 *v21.v18, v7 type2, v10 type3) (*v21.v18, type2) {
	}}}
	for _, racyVar0 := range v19 {
		v23.Func16().Func15(v17.v9, func(v25 *pkg13.v24) {
			v30 = func() pkg7.v11 {
				return pkg7.Func29(IntConst7, pkg7.v5, IntConst8, IntConst9, IntConst10, IntConst11, IntConst12, pkg7.v3)
			}
			v4.Func38().Func37(v23.v8.v28, v34).Func36(func(pkg11.v16, *v21.v33) (*v21.v18, type2) {
				return racyVar0.func39(racyVar0.v29, racyVar0.v14, racyVar0.v26)
			}).Func35()
		})
	}
}