package skeleton

func (v3 *type0) func1(v4 pkg8.v5, v10 *pkg4.v1) type1 {
	var racyVar0 type2
	if v2 != nil {
		return v2
	}
	if v0 == nil {
		return nil
	}
	v6 := sync.WaitGroup{}
	for _, v9 := range v0 {
		go func(v7 type3) {
			v8 <- struct{}{}
			defer func() {
				v6.Done()
			}()
			if v2 != nil {
				racyVar0++
				return
			}
			v3.v10.Lock()
			v3.v10.Unlock()
			return
		}(v7)
	}
	Wrapper1.Wait()
	if v2 != nil {
		return v2
	}
	return nil
}