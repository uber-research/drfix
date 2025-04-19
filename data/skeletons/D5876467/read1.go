package skeleton

func func1(v4 v1, v3 pkg2.v0, v7 v8, racyVar0 *type0, v5 *type0, v2 *sync.WaitGroup) {
	v4.Func1(v3, v7, func() {
		v6 := atomic.AddInt32(v5, IntConst0)
		if v6 > *racyVar0 {
			*racyVar0 = v6
		}
		atomic.AddInt32(v5, -IntConst2)
		v2.Done()
	})
}