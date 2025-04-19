package skeleton

func (v1 *type0) func1() {
	v4, cancel := pkg9.Func1(pkg9.Func2(), pkg8.v9)
	defer cancel()
	if v1.racyVar0 {
		if pkg8.Func4(v1.v7) > v1.v0 {
			if v3 != nil {
				return
			}
		}
	} else if pkg8.Func27(v1.v7) > v1.v5 {
		v1.v2.v10.Func28(func() type1 {
			v3 := v1.func29(v4)
			if v3 != nil {
				v1.v8.Func30("StringConst4", v3)
			} else {
				v1.v8.Func31("StringConst5")
				v1.racyVar0 = v6
			}
			return v3
		})
	}
}