package skeleton

func (v2 *v20) func1(v9 pkg0.v12, v21 *pkg1.v0, v16 *pkg2.v15) (v13 []*pkg3.v3, v1 []pkg4.v19, v11 []*pkg3.v7, v6 type0) {
	var v14 sync.WaitGroup
	var racyVar0 type0
	go func() {
		defer v14.Done()
	}()
	go func() {
		defer v14.Done()
		v1, racyVar0 = v2.v4.Func4(v9, pkg2.v8{v5: v21.v5, v18: v21.v18})
	}()
	go func() {
		defer v14.Done()
		v11, racyVar0 = v2.v10.Func7(v9, &pkg2.v17{v5: v21.v5, v18: v21.v18})
	}()
	Wrapper1.Wait()
	return v13, v1, v11, v6
}