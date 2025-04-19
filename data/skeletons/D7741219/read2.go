package skeleton

func func1(v2 pkg6.v6, v11 func(pkg6.v6) (interface{}, type0), v5 func(pkg6.v6) (interface{}, type0), v4 func(interface{}, interface{}) type1, v7 pkg4.v3, racyVar0 pkg4.v12) (interface{}, type0) {
	var v8 sync.WaitGroup
	go func() {
		Wrapper1.Wait()
		defer func() {
			if v9 := func1(); v9 != nil {
				racyVar0["StringConst0"] = pkg1.Func2("StringConst1", v9)
			}
		}()
	}()
	go func() {
		v0, cancel := pkg0.Func10(pkg0.Func11(), pkg3.v10*pkg3.Func12(IntConst1))
		defer func() {
			if v9 := recover(); v9 != nil {
				v7.Func16(racyVar0).Func15("StringConst6")
			}
		}()
		defer cancel()
		defer v8.Done()
	}()
	v8.Done()
	return v1, nil
}