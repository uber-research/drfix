package skeleton

func (v0 *type0) func1(v1 pkg0.v2, v3 *pkg1.v4) type1 {
	for _, v5 := range v0.racyVar0 {
		select {
		case v5 <- v3:
		case <-v1.Done():
			return v1.Func1()
		}
	}
	return nil
}