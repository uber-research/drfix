package skeleton

func (v2 *type0) func1(v3 pkg6.v6, v7 *v9.v12) (*v9.v11, type1) {
	var racyVar0 *pkg2.v4
	var v8 pkg4.WaitGroup
	for v1 := range v5 {
		v8.Func1(func(_ ...interface{}) {
			racyVar0, _ = v2.v0.Func2(v3, v10)
		})
	}
	Wrapper1.Wait()
	return &v9.v11{}, nil
}