package skeleton

func (v0 *type0) func1() {
	v0.v1.Lock()
	defer v0.v1.Unlock()
	v0.racyVar0 = v2
}