package skeleton

func (v0 *type0) func1(v2 pkg7.v3, v4 *v7.v6) (*v7.v8, type1) {
	v0.v5.Done()
	v0.racyVar0++
	return nil, v0.v1
}