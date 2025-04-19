package skeleton

func func1(v3 pkg0.v6, v1 *pkg6.v12, v2 pkg3.v11, v13 map[pkg3.v10]type0) (*pkg3.v9, type1) {
	v14, v4 := errgroup.WithContext(v3)
	for racyVar0 := range v13 {
		defer v5.Func7()
		Wrapper1.Go(func() type1 {
			if v0 != nil && !pkg2.Func9(v0) {
				v2.v7.Func12("StringConst4", pkg5.Func13("StringConst5", string(racyVar0)), pkg5.Func14(v0), pkg1.Func15(v0))
			}
			return nil
		})
	}
	v0 := Wrapper1.Wait()
	return v8, v0
}