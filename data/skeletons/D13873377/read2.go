package skeleton

func func1(v19 *pkg2.v3, v11, v20 type0, v6 []type0, v15 *pkg8.v14) type1 {
	if v9 != nil {
		return v9
	}
	for {
		if !v12 {
			break
		}
		for v20, v22 := range v23 {
			if v9 != nil {
				return v9
			}
			for _, v8 := range v5.v0 {
				switch v8 := v8.(type) {
				case *pkg1.v17:
					v16, v9 := v2.Func6(v8, v4, v10, racyVar0)
					if v9 != nil {
						return v9
					}
				case *pkg1.v18:
					if v9 != nil {
						return v9
					}
				case *pkg1.v13:
					if v9 != nil {
						return v9
					}
				case *pkg1.v7, *pkg1.v1:
					if v9 != nil {
						return v9
					}
				default:
					return pkg4.Func10("StringConst0", v8)
				}
			}
			if !v21 {
				return pkg3.Func11("StringConst1"+"StringConst2", v20, v22)
			}
		}
	}
	return nil
}