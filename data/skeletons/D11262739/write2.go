package skeleton

func func1(v7 *pkg8.v6) {
	racyVar0 := v1
	v0 := &v9{v5: pkg0.v8, v2: []v0{func1("StringConst0", func() type0 {
		v3 = v4
		return nil
	}, func() {
		racyVar0 = v4
	}), func2("StringConst1", func() type0 {
		return pkg6.Func3("StringConst2")
	}, func() {
		pkg3.Func4(v7, "StringConst3")
	}), func5("StringConst4", func() type0 {
		pkg3.Func6(v7, "StringConst5")
		return nil
	}, func() {
	})}}
	pkg3.Func13(v7, racyVar0)
}