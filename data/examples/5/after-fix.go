type asyncQueueTaskItem struct {
	item     queue.QueueItem
	deadline time.Time
	manager  *manager
}

func (w *asyncQueueTaskItem) Run(ctx context.Context) {
	w.manager.processItemAfterDequeue(w.item.(*queue.Item))
}

type asyncWorkerQueue struct {
	queue  queue.TimedQueue
	manager *manager

	taskChan chan queue.QueueItem
}

func newAsyncWorkerQueue(
	ddlQueue queue.TimedQueue,
	manager *manager,
) *asyncWorkerQueue {

	return &asyncWorkerQueue{
		queue:   ddlQueue,
		manager: manager,
	}
}

func (q *asyncWorkerQueue) Run(stopChan chan struct{}) {
	q.mu.Lock()
	q.taskChan = make(chan queue.QueueItem)
	taskChan := q.taskChan
	q.mu.Unlock()

	go func() {
		for {
			queueItem := q.queue.Dequeue(stopChan)
			if queueItem == nil {
				taskChan <- nil
				return
			}
			select {
			case taskChan <- queueItem:
				continue
			case <-stopChan:
				close(taskChan)
				return
			}
		}
	}()
}

func (q *asyncWorkerQueue) Enqueue(job async.Job) {
	asyncQueueItem := job.(*asyncQueueTaskItem)
	q.queue.Enqueue(asyncQueueItem.item, asyncQueueItem.deadline)
	return
}

func (q *asyncWorkerQueue) Dequeue() async.Job {
	q.mu.Lock()
	taskChan := q.taskChan
	q.mu.Unlock()

	queueItem := <-taskChan
	if queueItem == nil {
		return nil
	}

	return &asyncQueueTaskItem{
		item:    queueItem,
		manager: q.manager,
	}
}
