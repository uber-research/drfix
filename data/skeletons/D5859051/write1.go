package skeleton

func (v7 *type0) func1(v1 pkg6.v3, v2 type1, v12 []type1) ([]*pkg1.v5, type2) {
	var v4 sync.WaitGroup
	var racyVar0 type2
	go func() {
		defer v4.Done()
		v8, racyVar0 = v7.func3(v1, v2, v12)
		if v0 != nil {
			v9 <- v0
		}
	}()
	go func() {
		defer v4.Done()
		v10, racyVar0 = v7.func4(v1, v2, v12)
		if v0 != nil {
			v9 <- v0
		}
	}()
	Wrapper1.Wait()
	if !v6.Func7() {
		return nil, v6
	}
	return v7.v11.Func8(v1, v8, v10, nil)
}