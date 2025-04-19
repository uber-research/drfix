package skeleton

func (v2 *type0) func1(v9 pkg11.v15, v3 type1, v20 type1, v6 []*pkg3.v16, v21 []*pkg3.v19, v14 []type1, v5 pkg3.v4, v13 type2, v18 type3) (*type4, type5) {
	v7, racyVar0 := v2.factory.func29(v3)
	if v8 != nil {
		return nil, v8
	}
	var v17 sync.WaitGroup
	go func() {
		defer v17.Done()
		v0, racyVar0 = v7.Func33(v9, v10)
	}()
	go func() {
		defer v17.Done()
		v12, racyVar0 = v7.Func35(v9, v11)
	}()
	Wrapper1.Wait()
	return &type4{v1: v0, v22: func37(v12)}, nil
}