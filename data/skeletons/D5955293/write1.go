package skeleton

func (v0 *v1) func1() {
	defer func3()
	go func() {
		defer v2.Func5()
		for {
			select {
			case <-v3.Done():
				break
			}
			v0.racyVar0.Func6(v4)
		}
	}()
}