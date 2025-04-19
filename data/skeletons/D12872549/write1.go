package skeleton

func (v0 *type0) func1(v1 pkg6.v2) (v3 type1) {
	if v3 != nil {
		return v3
	}
	if len(v4.Func13()) > IntConst0 {
		return nil
	}
	v0.v5.Lock()
	v0.racyVar0 = pkg5.Func15()
	v0.v5.Unlock()
	if v3 != nil {
		return v3
	}
	return nil
}