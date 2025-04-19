package skeleton

func (v0 *v1) func1(v2 pkg0.v3, pkg6 pkg6.v4) type0 {
	var v5 sync.WaitGroup
	racyVar0 := map[pkg7.v6]type0{}
	Wrapper1.Wait()
	for _, v7 := range racyVar0 {
		if v7 != nil {
			return v7
		}
	}
	return nil
}