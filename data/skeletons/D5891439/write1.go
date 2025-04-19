package skeleton

func func1(v26 *pkg11.v25) {
	v1 := []struct {
		v3	string
		v7	func(*pkg6.v5, *pkg9.v12, string, chan struct{}) v24
		v28	string
	}{{v3: "StringConst0", v7: func(v14 *pkg6.v5, _ *pkg9.v12, v31 string, v20 chan struct{}) v24 {
		v4 := make(chan pkg13.v38)
		var v22 type0
		v36 := func(v9 pkg10.v17) (string, <-chan pkg13.v38, pkg14.v19, type1) {
			if !v22 {
				v20 <- struct{}{}
				v22 = v16
				return "StringConst1", v4, pkg14.v19{}, nil
			}
			go func() {
				v4 <- pkg13.v38{v15: IntConst0}
			}()
			v13, v8 := pkg5.Func2(v31)
			pkg3.Func3(v26, v8, "StringConst2")
			pkg4.Func4(v26, "StringConst3", string(v13))
			return "StringConst4", nil, pkg14.v19{}, pkg4.v32
		}
		return v36
	}, v28: "StringConst5"}, {v3: "StringConst6", v7: func(_ *pkg6.v5, _ *pkg9.v12, v31 string, v20 chan struct{}) v24 {
		v4 := make(chan pkg13.v38)
		v6 := make(chan pkg13.v38)
		var v22 type0
		v36 := func(v9 pkg10.v17) (string, <-chan pkg13.v38, pkg14.v19, type1) {
			if !v22 {
				v20 <- struct{}{}
				v22 = v16
				return "StringConst7", v4, pkg14.v19{}, nil
			}
			go func() {
				v6 <- pkg13.v38{v15: IntConst1}
				v4 <- pkg13.v38{v15: IntConst2}
			}()
			return "StringConst8", v6, pkg14.v19{}, nil
		}
		return v36
	}, v28: "StringConst9"}, {v3: "StringConst10", v7: func(v37 *pkg6.v5, v30 *pkg9.v12, v31 string, v20 chan struct{}) v24 {
		v4 := make(chan pkg13.v38)
		v6 := make(chan pkg13.v38)
		v11 := IntConst3 * pkg8.v2
		v18 := pkg14.v19{v35: pkg17.v29{v10: v11}}
		var v22 type0
		v36 := func(v9 pkg10.v17) (string, <-chan pkg13.v38, pkg14.v19, type1) {
			if !v22 {
				v20 <- struct{}{}
				v22 = v16
				return "StringConst11", v4, v18, nil
			}
			v37.Func11().Func10(pkg12.Func12(), "StringConst12").Func9(func(_ pkg10.v17, v34 string) type1 {
				go func() {
					v4 <- pkg13.v38{v15: IntConst4}
					v6 <- pkg13.v38{v15: IntConst5}
				}()
				return nil
			})
			go func() {
				v30.Func13(pkg8.v27, "StringConst13")
				v30.Func14(v11)
			}()
			v37.Func17().Func16(pkg12.Func18(), "StringConst14").Func15(pkg0.v21{v33: "StringConst15"}, nil)
			return "StringConst16", v6, v18, nil
		}
		return v36
	}}, {v3: "StringConst17", v7: func(v37 *pkg6.v5, _ *pkg9.v12, v31 string, v20 chan struct{}) v24 {
		v4 := make(chan pkg13.v38)
		v6 := make(chan pkg13.v38)
		v11 := IntConst6 * pkg8.v2
		v18 := pkg14.v19{v35: pkg17.v29{v10: v11}}
		var (
			v22	type0
			v23	func()
		)
		v36 := func(v9 pkg10.v17) (string, <-chan pkg13.v38, pkg14.v19, type1) {
			if !v22 {
				v20 <- struct{}{}
				v22 = v16
				return "StringConst18", v4, v18, nil
			}
			go func() {
				v4 <- pkg13.v38{v15: IntConst7}
			}()
			v23 = pkg15.Func21(v26, v31, IntConst8)
			return "StringConst19", v6, v18, nil
		}
		v37.Func24().Func23(pkg12.Func25(), "StringConst20").Func22(func(_ pkg10.v17, v34 string) type1 {
			func26()
			go func() {
				v6 <- pkg13.v38{v15: IntConst9}
			}()
			return nil
		})
		return v36
	}}, {v3: "StringConst21", v7: func(v37 *pkg6.v5, _ *pkg9.v12, v31 string, v20 chan struct{}) v24 {
		v4 := make(chan pkg13.v38)
		v6 := make(chan pkg13.v38)
		v11 := IntConst10 * pkg8.v2
		v18 := pkg14.v19{v35: pkg17.v29{v10: v11}}
		var v22 type0
		v36 := func(v9 pkg10.v17) (string, <-chan pkg13.v38, pkg14.v19, type1) {
			if !v22 {
				v20 <- struct{}{}
				v22 = v16
				return "StringConst22", v4, v18, nil
			}
			go func() {
				v4 <- pkg13.v38{v15: IntConst11}
			}()
			v37.Func31().Func30(pkg12.Func32(), "StringConst23").Func29(pkg0.v21{}, pkg4.v32)
			return "StringConst24", v6, v18, nil
		}
		v37.Func35().Func34(pkg12.Func36(), "StringConst25").Func33(func(_ pkg10.v17, v34 string) type1 {
			go func() {
				v6 <- pkg13.v38{v15: IntConst12}
			}()
			return nil
		})
		return v36
	}}, {v3: "StringConst26", v7: func(v37 *pkg6.v5, _ *pkg9.v12, v31 string, v20 chan struct{}) v24 {
		v36 := func(v9 pkg10.v17) (string, <-chan pkg13.v38, pkg14.v19, type1) {
			if !v22 {
				v20 <- struct{}{}
				return "StringConst27", v4, v18, nil
			}
			go func() {
				v4 <- pkg13.v38{v15: IntConst14}
				v6 <- pkg13.v38{v15: IntConst15}
			}()
			return "StringConst30", v6, v18, nil
		}
		return v36
	}}}
	for _, racyVar0 := range v1 {
		v26.Func43(v0.v3, func(v26 *pkg11.v25) {
			go func() {
				if racyVar0.v28 != "StringConst32" {
					return
				}
			}()
		})
	}
}