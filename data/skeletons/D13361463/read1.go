package skeleton

func func1(v3 type0) (*type1, type2) {
	var racyVar0 chan type3
	v3.v6.Func8(pkg17.v1{v2: func(pkg13.v7) type2 {
		if v9 != nil {
			v3.v8.Func9("StringConst4")
			go func() {
				racyVar0 = v9.Func10()
			}()
			v3.v8.Func11("StringConst5")
		}
		return nil
	}, v0: func(v4 pkg13.v7) type2 {
		if v9 != nil {
			close(racyVar0)
		}
		return nil
	}})
	return v5, nil
}