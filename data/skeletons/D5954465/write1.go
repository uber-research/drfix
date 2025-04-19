package skeleton

func (v0 *type0) func1(v1 type1, racyVar0 pkg13.v2, v3 type1, v4 type2) (pkg13.v5, type3) {
	v0.racyVar0 = racyVar0
	if v0.v6 == IntConst0 {
		return v0, pkg1.Func1("StringConst0")
	}
	return v0, nil
}