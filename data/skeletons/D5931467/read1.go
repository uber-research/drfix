package skeleton

func (v0 *type0) func1(v1 pkg20.v2, v3 []*pkg13.v4) {
	for _, racyVar0 := range v3 {
		v5 := func() {
			v6 := racyVar0.Func1()
		}
	}
}