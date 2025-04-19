package skeleton

func (v0 *type0) func1(v1 pkg0.v2) type1 {
	if v0.racyVar0 != nil {
		return pkg5.Func1(v0.v3, "StringConst0")
	}
	if _, v4 := v1.Func2(); !v4 {
		return pkg5.Func3("StringConst1")
	}
	v0.RLock()
	v0.RUnlock()
	if v5 != nil {
		select {
		case <-v1.Done():
			return pkg5.Func5("StringConst3")
		case <-v5:
			return nil
		}
	}
	return nil
}