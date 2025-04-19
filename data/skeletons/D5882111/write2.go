package skeleton

func (v0 *type0) func1(v1 pkg11.v2) {
	v0.v3.Lock()
	defer v0.v3.Unlock()
	v0.racyVar0 = v1.v4
}