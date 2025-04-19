package skeleton

func (v5 *type0) func1(v10 pkg8.v15, v25 *pkg3.v18) (*pkg3.v26, type1) {
	v10, cancel := v5.func1(v10, v25)
	defer cancel()
	var (
		v13	= pkg2.Func3(v10).Func2(map[type2]type2{"StringConst0": "StringConst1"})
		v14	= make(chan pkg7.v19, IntConst0)
		racyVar0	type1
		v0	= v6
		v23	= v6
		v12	= func() {
			var v22 pkg7.v19
			v22, racyVar0 = v5.v21.Func5(v10, v25)
			v14 <- v22
		}
	)
	select {
	case v22 := <-v14:
		if v7 != nil {
		} else if !v22.v1 {
			v17.v3 = v22.v3
			go func() {
				v4, cancel := pkg8.Func9(pkg8.Func10(), v5.v27.v2)
				defer cancel()
				v7 := v5.v20.Func11(v4, pkg11.v11{v8: v25.v8, v16: v25.v16, v9: pkg0.Func12(v10), v24: v22})
				if v7 != nil {
					pkg6.Func14().Func13("StringConst3", pkg6.Func15(v7))
				}
			}()
		}
	case <-v10.Done():
	}
	return v17, racyVar0
}