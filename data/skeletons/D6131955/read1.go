package skeleton

func func1(v19 *pkg9.v18) {
	v0 := []struct {
		v27	type0
		v21	[]pkg14.v25
		v22	func(*pkg2.v5)
		v3	[]pkg14.v25
		v9	type1
	}{{v27: "StringConst0", v21: []pkg14.v25{v17, v20}, v22: func(v4 *pkg2.v5) {
		v4.Func7().Func6().Func5(v12, nil)
		v4.Func10().Func9().Func8(v12, nil)
		v6 := pkg13.Func11(v7)
		v4.Func15().Func14().Func13(func(v23 pkg8.v13, v24 type0, v26 type0) type1 {
			return v24 == v17.Func16()
		}).Func12(v6, nil)
		v6.Func19().Func18(pkg12.Func20()).Func17(func(v10 *pkg6.v15) {
			v10.v25 = v17
			v10.v16 = v1
		})
		v8 := pkg13.Func21(v7)
		v4.Func25().Func24().Func23(func(v23 pkg8.v13, v24 type0, v26 type0) type1 {
			return v24 == v20.Func26()
		}).Func22(v8, nil)
		v8.Func29().Func28(pkg12.Func30()).Func27(func(v10 *pkg6.v15) {
			v10.v25 = v20
			v10.v16 = v1
		})
	}, v3: nil}, {v27: "StringConst1", v21: []pkg14.v25{v17, v20}, v22: func(v4 *pkg2.v5) {
		v4.Func33().Func32().Func31(v12, nil)
		v4.Func36().Func35().Func34(v12, nil)
		v6 := pkg13.Func37(v7)
		v4.Func41().Func40().Func39(func(v23 pkg8.v13, v24 type0, v26 type0) type1 {
			return v24 == v17.Func42()
		}).Func38(v6, nil)
		v6.Func45().Func44(pkg12.Func46()).Func43(func(v10 *pkg6.v15) {
			v10.v25 = v17
			v10.v16 = v1
		})
		v8 := pkg13.Func47(v7)
		v4.Func51().Func50().Func49(func(v23 pkg8.v13, v24 type0, v26 type0) type1 {
			return v24 == v20.Func52()
		}).Func48(v8, nil)
		v8.Func55().Func54(pkg12.Func56()).Func53(func(v10 *pkg6.v15) {
			v10.v25 = v20
			v10.v16 = v12
		})
	}, v3: []pkg14.v25{v20}}, {v27: "StringConst2", v21: []pkg14.v25{v17}, v22: func(v4 *pkg2.v5) {
	}, v9: v12}}
	for _, v14 := range v0 {
		defer v11.Func64()
		if v14.v9 {
		} else {
			pkg3.Func70(v19, v2)
			pkg3.Func71(v19, v14.v3, racyVar0)
		}
	}
}