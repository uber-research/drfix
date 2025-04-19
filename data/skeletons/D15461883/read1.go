package skeleton

func func1(v0 *pkg7.v1) {
	v2 := sync.WaitGroup{}
	go func() {
		defer v2.Done()
	}()
	Wrapper1.Wait()
	pkg4.Func25(v0, uint64(IntConst14), v3.racyVar0[v4{"StringConst10", IntConst15}])
}