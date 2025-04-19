package skeleton

func (v0 *type0) func1() {
	v0.v1.RLock()
	defer v0.v1.RUnlock()
	if v0.v2.v3 == IntConst2 && !v0.v2.racyVar0 {
		v0.v2.racyVar0 = v4
	}
}