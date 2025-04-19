package skeleton

func (v2 *type0) func1(v3 pkg9.v7, v9 []type1) ([]*v10.v8, pkg3) {
	racyVar0 := make([]*v10.v8, IntConst0)
	for v4 := IntConst1; v4 < len(v9); v4 += v6 {
		v5.Func6(func() pkg3 {
			if v0 != nil {
				return v0
			}
			racyVar0 = append(racyVar0, v1.Func9()...)
			return nil
		})
	}
	if v0 := Wrapper1.Wait(); v0 != nil {
		return nil, Func10(v0)
	}
	return v11, nil
}