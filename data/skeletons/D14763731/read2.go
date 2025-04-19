package skeleton

func (v6 *type0) func1(v1 pkg2.v8) {
	if v6.v2 == v1 {
		return
	}
	go func(v7 *pkg2.v5) {
		for v9 := range v7.v0 {
			v6.Lock()
			v3 := v6.func2(v9)
			v6.Unlock()
			if v3 != nil {
				return
			}
			if v6.racyVar0 != v7 {
				return
			}
		}
	}(v6.v4)
}