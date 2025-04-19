package skeleton

func (v7 *type0) func1(v2 pkg8.v4, v11 []pkg14.v13) ([]pkg14.v13, pkg4) {
	var racyVar0 []pkg14.v13
	v1, cancel := pkg8.Func5(v2)
	defer cancel()
	var v6 sync.WaitGroup
	for _, v9 := range v11 {
		go func(v9 pkg14.v13) {
			if v0 != nil {
				v10 <- v0
			} else {
				if v5 {
					v3 <- v9
				}
			}
			v6.Done()
		}(v9)
	}
	go func() {
		Wrapper1.Wait()
		for v12 := range v3 {
			racyVar0 = append(racyVar0, v12)
		}
	}()
	select {
	case <-v8:
		break
	case v0 := <-v10:
		return nil, v0
	}
	return v14, nil
}