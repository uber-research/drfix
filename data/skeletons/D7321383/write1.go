package skeleton

func (v4 *type0) func1(v5 pkg17.v7) (v9 map[type1]pkg7.v0, v2 map[type1]*pkg1.v1, racyVar0 type2) {
	var v6 sync.WaitGroup
	go func() {
		v2, racyVar0 = v4.v10.Func2(v5)
		v6.Done()
	}()
	v9, racyVar0 = v4.v8.Func5(v5)
	if v3 != nil {
		return nil, nil, v3
	}
	Wrapper1.Wait()
	return v9, v2, nil
}