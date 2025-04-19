package skeleton

func (v0 *type0) func1(racyVar0 pkg8.v1) {
	defer v0.v2.Unlock()
	v0.racyVar0 = racyVar0
}