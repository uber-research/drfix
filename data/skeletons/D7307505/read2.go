package skeleton

func (v0 *v9) func1(v3 pkg6.v6, v2 pkg3.v4) type0 {
	var v7 sync.WaitGroup
	var racyVar0 type1
	if v8 {
		go func() {
			defer v7.Done()
			racyVar0 = v0.func10(v3, v2)
		}()
		if racyVar0 != nil {
			return v1
		}
		return v5
	}
	return v1
}