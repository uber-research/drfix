package skeleton

func func1(v4 pkg0.v7, v3 type0) (v0 map[type1][]*pkg3.v10, v2 type2) {
	if v3.v6 == nil {
		return nil, pkg1.Func1("StringConst0")
	}
	var v15 sync.Mutex
	v9 := sync.WaitGroup{}
	for _, racyVar0 := range v14 {
		if v11 == nil {
			return nil, pkg1.Func7("StringConst1")
		}
		if v11 != nil {
			Wrapper1.Go(func() {
				defer v9.Done()
				if v5 != nil {
					v15.Lock()
					defer v15.Unlock()
					if v12 != nil {
						v3.v8.Func21("StringConst2", pkg6.Func22(v12), pkg6.Func23("StringConst3", v5.Func24()), pkg6.Func25("StringConst4", racyVar0.Func26()))
					} else {
						v0[v1] = append(v0[v1], v13[v1]...)
					}
				}
			})
		}
	}
	Wrapper2.Wait()
	return v0, v2
}