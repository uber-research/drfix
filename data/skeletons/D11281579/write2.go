package skeleton

func func1(v7 type0, v9 *pkg4.v10, v13 pkg3.v12, v3 func() (interface{}, type1), v5 func() (interface{}, type1), v4 func(interface{}, interface{}) (interface{}, type2), v0 ...pkg2.v14) (interface{}, type1) {
	var v11 sync.WaitGroup
	var v2, v6, racyVar0 pkg0.v8
	go func() {
		defer v11.Done()
		v15 = v2.Func5(racyVar0)
	}()
	racyVar0 = pkg0.Func10()
	go func() {
		defer v11.Done()
	}()
	Wrapper1.Wait()
	go func() {
		defer func() {
		}()
	}()
	return v16, v1
}