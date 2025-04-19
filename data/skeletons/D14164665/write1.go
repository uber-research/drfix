package skeleton

func func1(v10 *pkg34.v5, v6 *v7, v2 *pkg11.v4) (*pkg34.v5, type0) {
	defer func() {
	}()
	var racyVar0 type0
	Wrapper1.Go(func() type0 {
		v3, racyVar0 = v6.Func6(v10)
		return v1
	})
	Wrapper1.Go(func() type0 {
		v9, racyVar0 = pkg0.Func7(v2, v6.v11, v6.v0, v10.v8)
		return v1
	})
	v1 = Wrapper1.Wait()
	return v3, v1
}