package skeleton

func (v2 *type0) func1(v4 pkg32.v7, v3 pkg13.v13, v10 *v8.v5, v11 pkg28.v12) (racyVar0 v14.v16, v1 type1) {
	Wrapper1.Go(func() {
	})
	Wrapper1.Go(func() {
		v2.func44(racyVar0, v3.v9, racyVar0.v6)
	})
	racyVar0 = pkg34.Func45(v4, racyVar0, v3, v2.v0)
	Wrapper1.Go(func() {
	})
	return v15, v1
}