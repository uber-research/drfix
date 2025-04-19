package skeleton

func (v6 *type0) func1() (map[type1]type2, type3) {
	var (
		v7	= make(map[type1]type2)
		v5	sync.Mutex
		v4	= pkg5.Func2(pkg0.Func3())
	)
	for _, racyVar0 := range v9.v2 {
		Wrapper1.Go(func(v1 pkg0.v3) type3 {
			v8, v0 := v9.func4(racyVar0)
			if v0 != nil {
				return v0
			}
			v5.Lock()
			defer v5.Unlock()
			return nil
		})
	}
	if v0 := Wrapper1.Wait(); v0 != nil {
		return nil, v0
	}
	return v7, nil
}