func testProcessReplies(t *testing.T, actor string, nodes []string, sysMode config.SystemType) {
	updatedTime := time.Now()
	updatedTime = updatedTime.Add(-10 * time.Minute)

	store := makeStorage("someTable", []string{"someCol"}, 4, sysMode)
	store.Instance.PeerRegions = make(map[string]config.Peer)
	for _, node := range nodes {
		store.Instance.PeerRegions[node] = config.Peer{
			Name: node,
			ID:   someUtility.GetNodeID(node),
		}
	}

	u := store.newUpdater()
	u.client = newMockedClient(nil)
	u.lastUpdate = updatedTime
	u.delay = store.config.Herb.IntVar(u, "some_delay_ms", 0)

	uAck := make(chan sch.ReplicationReply, 1)
	var wg1 sync.WaitGroup
	wg1.Add(1)
	go func() {
		defer wg1.Done()
		u.processReplies(uAck)
	}()

	var wg2 sync.WaitGroup
	wg2.Add(1)
	go func() {
		defer wg2.Done()
		for _, node := range nodes {
			generateReplies(uAck, node, "someTable.someCol", []uint{1}, []uint64{12345})
			generateReplies(uAck, node, "someTable.someCol", []uint{2}, []uint64{12347})
			generateReplies(uAck, node, "someTable.someCol", []uint{1}, []uint64{1234})
		}
	}()
	wg2.Wait()
	close(uAck)
	wg1.Wait()

	var consumers []string
	for _, node := range nodes {
		peerID, err := u.getPeerID(node)
		require.NoError(t, err)
		consumers = append(consumers, peerID)
	}

	for _, c := range consumers {
		o := u.consumerOffsets[c]
		assert.True(t, len(o.newOffsets["someCol"].Offsets) > 0)
		assert.Equal(t, true, reflect.DeepEqual(o.newOffsets["someCol"].Offsets, o.origOffsets["someCol"].Offsets))
	}

	u.lastUpdate = time.Now()
	u.delay = store.config.Herb.IntVar(u, "some_delay_ms", 1000)
	uAck = make(chan sch.ReplicationReply, 1)
	var wg3 sync.WaitGroup
	wg3.Add(1)
	go func() {
		defer wg3.Done()
		u.processReplies(uAck)
	}()

	var wg4 sync.WaitGroup
	wg4.Add(1)
	go func() {
		defer wg4.Done()
		for _, node := range nodes {
			generateReplies(uAck, node, "someTable.someCol", []uint{1, 1}, []uint64{11, 778})
			generateReplies(uAck, node, "someTable.someCol", []uint{2}, []uint64{12349})
			generateReplies(uAck, node, "someTable.someCol", []uint{1}, []uint64{12345})
			generateReplies(uAck, node, "someTable.someCol", []uint{3}, []uint64{11, 99})
			generateReplies(uAck, node, "someTable.someCol", []uint{1, 1}, []uint64{7777, 12349})
			generateReplies(uAck, node, "someTable.someCol", []uint{2}, []uint64{13000, 99000})
			generateReplies(uAck, node, "someTable.someCol", []uint{1}, []uint64{90000})
		}
	}()
	wg4.Wait()
	close(uAck)
	wg3.Wait()

	for _, c := range consumers {
		o := u.consumerOffsets[c]

		o.maxOffsets = make(map[dsKey]uint64)

		assert.Equal(t, uint64(12345), o.newOffsets["someCol"].Offsets[1])
		assert.Equal(t, uint64(12347), o.newOffsets["someCol"].Offsets[2])

		assert.Equal(t, uint64(12345), o.origOffsets["someCol"].Offsets[1])
		assert.Equal(t, uint64(12347), o.origOffsets["someCol"].Offsets[2])

		assert.Equal(t, true, reflect.DeepEqual(o.newOffsets["someCol"].Offsets, o.origOffsets["someCol"].Offsets))
	}
}
