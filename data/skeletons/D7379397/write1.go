package skeleton

func (v7 *type0) func1(v3 pkg8.v6, v11 []pkg13.v12) (map[pkg13.v12]*v5.v13, type1) {
	for _, racyVar0 := range v9 {
		Wrapper1.Go(func() type1 {
			v0 := make([]interface{}, func11(racyVar0))
			if v1 := v7.v2.Func12(v3, v10, v0...); v1 != nil {
				v8 <- v1
				return v1
			}
			v4 <- v0
			return nil
		})
	}
	return Func16(v3, v4, v8, v9, v11)
}