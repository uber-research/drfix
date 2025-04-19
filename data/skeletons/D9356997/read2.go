package skeleton

func (v4 *type0) func1(v3 pkg7.v5, v12 *v8.v11) (*v8.v10, type1) {
	v9, racyVar0 := pkg0.Func3(v3, v12)
	if v1 != nil {
		return nil, pkg0.Func4(v1)
	}
	var v6 sync.WaitGroup
	go func() {
		defer v6.Done()
	}()
	go func() {
		defer v6.Done()
		v7, racyVar0 = v4.v2.Func9(v3, v9)
	}()
	if racyVar0 != nil {
		return nil, pkg0.Func10(v1)
	}
	Wrapper1.Wait()
	return pkg0.Func11(v3, v7, v0)
}