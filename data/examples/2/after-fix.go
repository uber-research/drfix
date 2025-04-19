func (cl *ConcurrencyLimiter) backgroundWorker() {
	cl.backgroundWg.Add(1)
	go func() {
		defer cl.backgroundWg.Done()
		for {
			select {
			case <-cl.backgroundChan:
			case <-cl.backgroundCtx.Done():
				return
			}

			cl.updateMu.Lock()
			limit := cl.limit
			oldTp := cl.pool()
			cl.updateMu.Unlock()

			if cap(oldTp) == limit {
				continue
			}

			newTp := newTokenPool(limit)
			tokensToMove := limit
			delta := limit - cap(oldTp)
			if delta > 0 {
				addTokens(newTp, delta)
				tokensToMove -= delta
			}
			cl.setPool(newTp)
			drainTokens(cl.backgroundCtx, oldTp, tokensToMove, func(t *token) { newTp.addToken() })
			if delta < 0 {
				drainTokens(cl.backgroundCtx, oldTp, -delta, func(*token) {})
			}
			close(oldTp)
		}
	}()
}
