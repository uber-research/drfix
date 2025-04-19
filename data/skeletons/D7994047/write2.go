package skeleton

func (v0 *type0) func1(v2 pkg0.v4, v8 *type1) (*type2, type3) {
	var racyVar0 []type3
	var v5 sync.WaitGroup
	for v9 := range v8.v7 {
		go func() {
			defer v5.Done()
			if v1 == nil {
			} else {
				racyVar0 = append(racyVar0, v1)
			}
		}()
	}
	Wrapper1.Wait()
	if len(v3) > IntConst1 {
		return nil, pkg6.Func5(v3...)
	}
	return &type2{v7: v6}, nil
}