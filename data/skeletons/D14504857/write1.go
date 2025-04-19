package skeleton

func (v0 *type0) func1() <-chan struct{} {
	go func() {
		v0.racyVar0.Wait()
	}()
	return v1
}