package skeleton

func (v3 *v14) func1(v5 pkg6.v9, v10 pkg0.v0, v11 []pkg8.v15, v6 *pkg5.v2) type0 {
	for _, racyVar0 := range v12 {
		v16.Func12(v7, func() {
			v4 := v3.v13.Func13(v5, v10, *racyVar0)
			if v4 != nil {
				v8 <- v1
			} else {
				v8 <- nil
			}
		})
	}
	return nil
}