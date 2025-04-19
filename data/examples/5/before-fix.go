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
	q.taskChan = make(chan queue.QueueItem)

	go func() {
		for {
			queueItem := q.queue.Dequeue(stopChan)
			if queueItem == nil {
				q.taskChan <- nil
				return
			}
			select {
			case q.taskChan <- queueItem:
				continue
			case <-stopChan:
				close(q.taskChan)
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
	queueItem := <-q.taskChan
	if queueItem == nil {
		return nil
	}

	return &asyncQueueTaskItem{
		item:    queueItem,
		manager: q.manager,
	}
}
