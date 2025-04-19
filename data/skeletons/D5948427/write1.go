package skeleton

func func1(v17 *pkg10.v16) {
	for _, racyVar0 := range v14 {
		v17.Func5(v11.v7, func(v17 *pkg10.v16) {
			v1 := func(pkg9.v10, interface{}) (interface{}, type0) {
				return nil, nil
			}
			v13 := sync.WaitGroup{}
			defer Wrapper1.Wait()
			v22.Func17().Func16(pkg2.v19{v0: "StringConst30", v18: "StringConst31", v15: pkg1.Func18(v11.v9, "StringConst32")[IntConst4], v21: v11.v12, v2: "StringConst33", v4: v11.v5, v3: map[type1]type1{pkg1.Func19(pkg0.v8): "StringConst34", pkg1.Func20(pkg0.v20): "StringConst35"}}).Func15(func(_ ...interface{}) type0 {
				v13.Done()
				return racyVar0.v6
			})
		})
	}
}