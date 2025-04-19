package skeleton

func (v0 *type0) func1() {
	go func() {
		for {
			select {
			case <-v0.v1.Done():
				return
			}
			v0.v5.Lock()
			v0.v5.Unlock()
			if cap(v2) == v3 {
				continue
			}
			v0.racyVar0.Add(IntConst0)
			func() {
				defer v0.v4.Done()
			}()
		}
	}()
}