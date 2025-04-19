package skeleton

func func1(v0 *pkg13.v1) {
	var racyVar0 sync.WaitGroup
	for v2 := IntConst8; v2 < IntConst9; v2++ {
		go func(v3 type0) {
			racyVar0.Add(IntConst10)
			defer v4.Done()
		}(pkg9.Func15("StringConst2", v2))
	}
	Wrapper1.Wait()
}