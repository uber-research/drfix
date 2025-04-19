package skeleton

func (v1 *type0) func1() {
	v8, cancel := pkg6.Func2(pkg6.Func3(), IntConst0*pkg5.v13)
	defer cancel()
	if v15 {
		if pkg5.Func5(v1.v11) > v1.v0 {
			if v5 != nil {
				return
			}
			if len(v17.Func22()) > IntConst1 {
				for _, v7 := range v17.Func26() {
					if v5 != nil {
						return
					}
				}
			} else if len(v17.Func34()) == IntConst3 {
				v1.v12.Func35("StringConst8")
				v1.v20.Lock()
				v2 := v1.v2
				v1.v20.Unlock()
				if !v2 {
					if v1.v19.v16 == IntConst4 {
						v1.v3.v6.Func36(v4)
					} else {
						v1.v3.v14.Func37(v4)
					}
				}
				v1.v20.Lock()
				v1.v15 = v4
				v1.v2 = v4
				v1.v20.Unlock()
			}
		}
	} else if pkg5.Func38(v1.racyVar0) > v1.v9 {
		v1.v12.Func39("StringConst9")
		v1.v3.v18.Func40(func() type1 {
			v5 := v1.func41(v8)
			v1.v20.Lock()
			v1.v15 = v10
			v1.v20.Unlock()
			if v5 != nil {
				v1.v12.Func42("StringConst10", v5)
			} else {
				v1.v12.Func43("StringConst11")
			}
			return v5
		})
	}
}