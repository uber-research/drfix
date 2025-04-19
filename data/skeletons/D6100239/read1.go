package skeleton

func func1(v16 *pkg7.v15) {
	racyVar0 := IntConst1
	v0 := map[type0]struct {
		v14	func(pkg6.v12) type1
		v6	func(pkg6.v12) type1
		v21	type2
		v5	pkg4.v18
		v20	type3
		v3	type3
		v9	type2
	}{"StringConst0": {v6: func(v4 pkg6.v12) type1 {
		pkg4.Func1(pkg4.v1 * IntConst2)
		v19++
		return nil
	}, v14: func(v4 pkg6.v12) type1 {
		v8++
		return nil
	}, v5: pkg4.v1 * IntConst3, v21: v2, v20: IntConst4, v3: IntConst5}, "StringConst1": {v6: func(v4 pkg6.v12) type1 {
		v19++
		return nil
	}, v14: func(v4 pkg6.v12) type1 {
		pkg4.Func2(pkg4.v1 * IntConst6)
		v8++
		return nil
	}, v5: pkg4.v1 * IntConst7, v21: v2, v20: IntConst8, v3: IntConst9}, "StringConst2": {v6: func(v4 pkg6.v12) type1 {
		v19++
		return pkg3.Func3("StringConst3")
	}, v14: func(v4 pkg6.v12) type1 {
		pkg4.Func4(pkg4.v1 * IntConst10)
		v8++
		return nil
	}, v5: pkg4.v1 * IntConst11, v21: v2, v20: IntConst12, v3: IntConst13}, "StringConst4": {v6: func(v4 pkg6.v12) type1 {
		v19++
		return pkg3.Func5("StringConst5")
	}, v14: func(v4 pkg6.v12) type1 {
		pkg4.Func6(pkg4.v1 * IntConst14)
		v8++
		return nil
	}, v5: pkg4.v1 * IntConst15, v21: v2, v20: IntConst16, v3: IntConst17, v9: v10}, "StringConst6": {v6: func(v4 pkg6.v12) type1 {
		racyVar0++
		panic("StringConst7")
		return nil
	}, v14: func(v4 pkg6.v12) type1 {
		pkg4.Func8(pkg4.v1 * IntConst18)
		v8++
		return nil
	}, v5: pkg4.v1 * IntConst19, v21: v2, v20: IntConst20, v3: IntConst21}, "StringConst8": {v6: func(v4 pkg6.v12) type1 {
		pkg4.Func9(pkg4.v1 * IntConst22)
		v19++
		return nil
	}, v14: func(v4 pkg6.v12) type1 {
		v8++
		return nil
	}, v5: IntConst23 * pkg4.v1, v21: v10, v20: IntConst24, v3: IntConst25}, "StringConst9": {v6: func(v4 pkg6.v12) type1 {
		v19++
		return nil
	}, v14: func(v4 pkg6.v12) type1 {
		pkg4.Func10(pkg4.v1 * IntConst26)
		v8++
		return nil
	}, v5: IntConst27 * pkg4.v1, v21: v10, v20: IntConst28, v3: IntConst29}, "StringConst10": {v6: func(v4 pkg6.v12) type1 {
		v19++
		return nil
	}, v14: func(v4 pkg6.v12) type1 {
		pkg4.Func11(pkg4.v1 * IntConst30)
		v8++
		return pkg3.Func12("StringConst11")
	}, v5: IntConst31 * pkg4.v1, v21: v10, v20: IntConst32, v3: IntConst33}, "StringConst12": {v6: func(v4 pkg6.v12) type1 {
		v19++
		return pkg3.Func13("StringConst13")
	}, v14: func(v4 pkg6.v12) type1 {
		pkg4.Func14(pkg4.v1 * IntConst34)
		v8++
		return nil
	}, v5: IntConst35 * pkg4.v1, v21: v10, v20: IntConst36, v3: IntConst37, v9: v10}, "StringConst14": {v6: func(v4 pkg6.v12) type1 {
		v19++
		return nil
	}, v14: func(v4 pkg6.v12) type1 {
		return nil
	}, v5: IntConst39 * pkg4.v1, v21: v10, v20: IntConst40, v3: IntConst41}}
	for v7, v13 := range v0 {
		v11, cancel := pkg6.Func17(pkg6.Func18(), pkg4.v17)
		cancel()
		pkg2.Func38(v16, v13.v3, racyVar0)
	}
}