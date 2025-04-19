package skeleton

func (v10 *type0) func1(v4 pkg4.v9, v15 pkg8.v14) (v0 func(), v12 func() type1) {
	var racyVar0 type1
	go func() {
		for {
			select {
			case <-v4.Done():
				return
			case <-v1:
				return
			case v13 := <-v6.v2:
				if v3 != nil {
					continue
				}
				if v3 := func7(v7.v11, v13); v3 != nil {
					racyVar0 = v3
					return
				}
				if v3 := func9(v7.v5); v3 != nil {
					return
				}
			}
		}
	}()
	v0 = func() {
		v1 <- v8
	}
	v12 = func() type1 {
		return racyVar0
	}
	return v0, v12
}