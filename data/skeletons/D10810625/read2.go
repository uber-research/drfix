package skeleton

func (v37 *v54) func1(v13 pkg16.v27, v45 *pkg7.v15, v29 *v10.v14, v16 v10.v57, v47 type0, v46 type1, v49 type0, v19 type0) (v21 *v10.v18, v11 *v36, v9 type2) {
	defer v37.v51.v44.Func2(v13, v37.v51.v50.Func3(), "StringConst0").Func1(func(pkg16 pkg16.v27, v24 *pkg12.v23) {
		v24.v52 = pkg12.v59{"StringConst1": v58.Func4(v29), "StringConst2": v48, "StringConst3": v32, "StringConst4": v40, "StringConst5": v49}
		v24.v8 = v9
	})
	if v37.func7(v45) || v55 {
		v45.racyVar0 = v37.func8(v45)
	}
	v17 := map[type1]type1{"StringConst6": pkg5.Func12(v55), "StringConst7": pkg5.Func13(v45.racyVar0), "StringConst8": v46, "StringConst9": pkg5.Func14(v49), "StringConst10": pkg5.Func15(v45.v12), "StringConst11": pkg5.Func16(v45.v42), "StringConst12": pkg5.Func17(func18(v45.Func19()) == IntConst0), "StringConst13": pkg5.Func20(func21(v45)), "StringConst14": pkg5.Func22(v4), "StringConst15": pkg5.Func23(v41)}
	defer func() {
	}()
	if v9 != nil {
		return nil, v11, v9
	}
	if v49 {
		if v56 != nil {
			Wrapper1.Go(func() type2 {
				return nil
			})
		}
		if v38 != nil {
			Wrapper1.Go(func() type2 {
				return nil
			})
		}
	} else {
		Wrapper1.Go(func() type2 {
			if v56 == nil {
				return nil
			}
			v1 := v56.v30
			v56 = v37.func82(v13, v45, v56)
			v28 := pkg15.Func83()
			if v4 {
				v37.v6.Func85("StringConst38").Func84(IntConst4)
				v39, _, v32 = v37.func86(v13, v45, v56, v46, v7)
			} else {
				v39, _, v32 = v37.func87(v13, v45, v56, v16, v47, v46, v19)
			}
			if v39 != nil {
				v39 = func88(v39, v1)
				pkg17.Func89(v37.v6, "StringConst39", len(v39.v0), v17)
			}
			v5 = pkg15.Func91(v28)
			v37.v6.Func94(v17).Func93("StringConst40", pkg4.v53).Func92(v5)
			v37.v6.Func97(v17).Func96("StringConst41", pkg4.v33).Func95(v5)
			v37.v6.Func100(v43).Func99("StringConst42", pkg4.v2).Func98(v5)
			if v32 != nil {
				v37.v6.Func102("StringConst43").Func101(IntConst5)
				v37.v31.Func103("StringConst44", pkg11.Func104("StringConst45", v58.Func105(v29)), pkg11.Func106("StringConst46", v58.Func107(v56)), pkg11.Func108(v32))
			}
			return nil
		})
	}
	Wrapper1.Go(func() type2 {
		if v3 == nil {
			return nil
		}
		if v37.func109(v45) && func110(v3, v25) {
			return nil
		}
		return nil
	})
	if v40 = Wrapper1.Wait(); v40 != nil {
		return nil, v11, v40
	}
	if v49 {
		if v32 != nil && v48 != nil && v20 != nil {
			return nil, v11, pkg14.Func123("StringConst50" + v32.Func124() + "StringConst51" + v48.Func125() + "StringConst52" + v20.Func126())
		}
	} else {
		if v32 != nil && v48 != nil {
			return nil, v11, pkg14.Func128(v32, v48.Func129())
		}
		v37.v6.Func131("StringConst53").Func130(v34)
		v35 := []*v10.v18{v39, v22}
		v26 := []type2{v32, v48}
		v21 = func132(v35, v26, v29.v30)
	}
	return v21, v11, nil
}