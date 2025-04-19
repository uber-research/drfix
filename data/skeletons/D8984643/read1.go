package skeleton

func func1(v14 *pkg9.v13) {
	defer v7.Func2()
	v17 := v18{v8: "StringConst0", v3: "StringConst1", v25: "StringConst2", v1: pkg6.v15, v5: func(v4 pkg8.v9) ([]v6, type0) {
		return []v6{{v21: "StringConst3", v23: []interface{}{"StringConst4", IntConst0}}}, nil
	}}
	var racyVar0 *pkg0.v2
	v10 := &sync.WaitGroup{}
	go func() {
		defer v10.Done()
		racyVar0 = pkg0.Func9(v14, pkg10.Func10(pkg10.Func11(pkg10.v19{v24: "StringConst5", v11: func() *v18 {
			return &v17
		}}, pkg10.v19{v24: "StringConst6", v11: func() *v18 {
			return nil
		}}, func() v16 {
			return v0
		}, func() pkg7.v12 {
			return pkg7.v20
		}), pkg10.Func12(v22))).Func8()
	}()
	racyVar0.Func14()
	Wrapper1.Wait()
}