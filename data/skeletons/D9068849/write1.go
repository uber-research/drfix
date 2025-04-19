package skeleton

func (v12 *type0) func1(v5 pkg4.v9, v6 type1, v1 type1, v11 type1, v15 func(pkg4.v9, []pkg7.v13) (interface{}, type2), v8 func(pkg4.v9, []pkg7.v13) (interface{}, type2)) (interface{}, type2) {
	racyVar0 := Func1(v5)
	if v11 == "StringConst0" {
		return func2(v5, v0)
	}
	if v11 == "StringConst1" {
		return func3(v5, v0)
	}
	var v10 sync.WaitGroup
	Wrapper1.Go(func() {
		Wrapper2.Wait()
		defer func() {
		}()
	})
	Wrapper1.Go(func() {
		v3, cancel := pkg4.Func20(pkg4.Func21(), pkg3.v14*pkg3.Func22(v16))
		defer func() {
		}()
		defer v10.Done()
		defer cancel()
		defer v7.Func32()
		racyVar0 = append(racyVar0, pkg7.Func35(v2, "StringConst15"))
	})
	defer v10.Done()
	v4, v17 = func37(v5, racyVar0)
	return v4, v17
}