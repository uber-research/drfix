package skeleton

func (v2 *v1) func1(v3 pkg0.v4, v10 pkg7.v7, v8 pkg5.v9, v0 pkg7.v7, v5 *sync.WaitGroup, racyVar0 map[pkg7.v7]type0) {
	defer func() {
		if v6 := recover(); v6 != nil {
			v5.Done()
			racyVar0[v0] = pkg1.Func2("StringConst0")
		}
	}()
	v5.Done()
}