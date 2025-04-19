package skeleton

func (v0 *type0) func1(v1 pkg0.v2, v3 type1) (pkg2.v4, type2) {
	select {
	case <-v1.Done():
		return nil, v1.Func10()
	}
	if v5.racyVar0 != nil {
		return nil, v5.v6
	}
	return v7, nil
}