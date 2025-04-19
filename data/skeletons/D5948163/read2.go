package skeleton

func (v0 *type0) func1() {
	for v1 {
		select {
		case <-v0.racyVar0:
			return
		}
	}
}