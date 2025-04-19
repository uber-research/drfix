package skeleton

func (v0 *type0) func1(v1 pkg0.v2) type1 {
	select {
	case <-v0.v3:
		return nil
	case <-v1.Done():
		v0.racyVar0 = v4
		return v4
	}
}