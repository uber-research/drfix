package skeleton

func func1(v9 *type0) v10 {
	return func(v2 pkg2.v5) (interface{}, type1) {
		var v4 type2
		v9.v3.Lock()
		v4 = v9.v7
		v9.racyVar0++
		v9.v3.Unlock()
		if v9.v8[v4].v0 > IntConst0 {
			pkg1.Func1(v9.v8[v4].v0)
		}
		return v9.v8[v4].v6, v9.v8[v4].v1
	}
}