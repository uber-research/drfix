package skeleton

func (v8 *type0) func1(v4 pkg10.v12, v5 type1, v2 []type1, v0 type2) (*type3, map[type1]type4) {
	racyVar0, v11 := make([]type1, IntConst0), func2([]type1, IntConst1)
	var v7 sync.WaitGroup
	for _, v14 := range v2 {
		go func() {
			defer v7.Done()
			if v3 != nil {
				return
			}
			if v9 {
				racyVar0 = append(racyVar0, v10)
			} else {
				v11 = append(v11, v10)
			}
		}()
	}
	Wrapper1.Wait()
	return &type3{v13: v15, v6: v11}, v1
}