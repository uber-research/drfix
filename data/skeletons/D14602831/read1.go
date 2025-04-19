package skeleton

func (v5 *type0) func1(v10 pkg18.v15, v6 []pkg2.v22, v20 pkg2.v24, v2 type1, v3 type1, v1 type2, v11 type2) (pkg2.v25, type3) {
	var racyVar0 type3
	if v8 != nil {
		return pkg2.v25{}, v8
	}
	if v3 == v13 {
		return v5.func45(v10, v6, v2, v7)
	}
	if v3 != v12 && v3 != v17 {
		return pkg2.v25{}, nil
	}
	if v19 && !v11 {
		go v5.v26.Func46(func() {
			v9, cancel := pkg18.Func47(pkg18.Func48(), IntConst0*pkg15.v18)
			defer cancel()
			_, v16 := v5.Func49(pkg8.Func50(v9, pkg8.v14, v21), v6, v20, v2, v3, v4, v1, v11)
			if v16 != nil {
				racyVar0 = pkg12.Func51(racyVar0, v16)
			}
			if v8 != nil {
				v7.Func53(pkg7.Func54(v8)).Func52("StringConst16")
			} else {
				v7.Func55("StringConst17", pkg7.Func56("StringConst18", "StringConst19"), pkg7.Func57("StringConst20", (pkg2.v0)(v6)))
			}
		})()
		return pkg2.v25{}, racyVar0
	}
	return v23, nil
}