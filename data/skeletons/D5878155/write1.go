package skeleton

func (v25 *pkg7) func1(v5 pkg0.v15, v29 *v24.v6) (*v24.v11, type0) {
	if v29 == nil {
		return nil, v28
	}
	if v29.v26 != nil {
		v30 := sync.Mutex{}
		v2 := pkg3.Func6(IntConst2, len(v27), IntConst3, v9, func(v10 type1) type0 {
			if v23 != nil {
				return v23
			}
			v20.racyVar0 = v16
			v30.Lock()
			defer v30.Unlock()
			return nil
		})
	} else if v29.v3 != nil {
		for _, v1 := range v29.v3 {
			v13 := v0
			v13 = append(v13, pkg4.Func20(v1.v4))
			if v29.v17 != nil {
				v13 = append(v13, pkg4.Func22(*v29.v17))
			}
			if v1.v14 != nil {
				v13 = append(v13, pkg4.Func24(*v1.v14))
			}
			if v1.v8 != nil {
				v13 = append(v13, pkg4.Func26(*v1.v8))
			}
			if v1.v19 != nil {
				v13 = append(v13, pkg4.Func28(*v1.v19))
			}
			v21, v16, v23 := v25.v12.Func29(v5, v13...)
			if v23 != nil {
				v25.v18.Func30("StringConst3", pkg6.Func31(v23), pkg6.Func32("StringConst4", v29))
				v20.v7 = pkg2.Func33("StringConst5", v23.Func34())
				break
			}
			v20.v17 = v16
			v20.v22 = append(v20.v22, v21...)
		}
	} else {
		v21, v16, v23 := v25.v12.Func36(v5, v0...)
		if v23 != nil {
			v25.v18.Func37("StringConst6", pkg6.Func38(v23), pkg6.Func39("StringConst7", v29))
			v20.v7 = pkg2.Func40("StringConst8", v23.Func41())
		}
		v20.v17 = v16
		v20.v22 = v21
	}
	return &v20, nil
}