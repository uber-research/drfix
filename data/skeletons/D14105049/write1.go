package skeleton

func (v16 *type0) func1(v21 *v1.v15, v4 v1.v0) (racyVar0 type1) {
	defer v11.Func6()
	defer close(v10)
	defer close(v9)
	defer close(v2)
	if v7 != nil {
		return pkg7.Func15("StringConst3", v21.v17.Func16(), v7)
	}
	var v14 sync.WaitGroup
	v8, cancel := pkg19.Func17(v4.Func18())
	defer cancel()
	v5, racyVar0 := v16.v18.Func20(v21.v17.Func21())
	if v7 != nil {
		return pkg7.Func23("StringConst5", v21.v17.Func24(), v7)
	}
	if v5.Func25() != pkg24.v12 {
		if v7 != nil {
			return pkg7.Func31("StringConst6", v21.v17.Func32(), v7)
		}
		return
	}
	go func() {
		defer v14.Done()
		for {
			select {
			case <-v11.v3:
				racyVar0 = v16.func34(v8, v21, v9, v2, v5, !v6)
				if v7 != nil {
					v10 <- v7
					continue
				}
			case <-v22:
				return
			}
		}
	}()
	for {
		select {
		case <-v4.Func38().Done():
			cancel()
			Wrapper1.Wait()
			return
		case v20, v19 := <-v9:
			if !v19 {
				return nil
			}
			if racyVar0 = v4.Func43(&v20); racyVar0 != nil {
				cancel()
				Wrapper1.Wait()
				return pkg7.Func48("StringConst15", v21.v17.Func49(), v7)
			}
		case v13 := <-v10:
			cancel()
			Wrapper1.Wait()
			return pkg7.Func53("StringConst17", v21.v17.Func54(), v13)
		case <-v2:
			cancel()
			Wrapper1.Wait()
			if v7 != nil {
				return pkg7.Func63("StringConst20", v7)
			}
			return
		}
	}
}