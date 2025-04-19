package skeleton

func (v0 *type0) func1(v3 pkg6.v4, v5 pkg2.v6, v7 pkg2.v2) type1 {
	for racyVar0 := range v9 {
		Wrapper1.Go(func() type1 {
			return v0.pkg8.Func2(v3, v9[racyVar0])
		})
	}
	if v1 := Wrapper1.Wait(); v1 != nil {
		return v1
	}
	return nil
}