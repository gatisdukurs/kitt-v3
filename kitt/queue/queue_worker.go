package queue

import (
	"context"
	"fmt"
)

type QueueWorker struct {
	name       string
	queue      *InMemoryQueue
	dispatcher JobDispatcher
}

func NewQueueWorker(name string, queue *InMemoryQueue, dispatcher JobDispatcher) *QueueWorker {
	return &QueueWorker{
		name:       name,
		queue:      queue,
		dispatcher: dispatcher,
	}
}

func (w *QueueWorker) Id() string {
	return w.name
}

func (w *QueueWorker) Run(ctx context.Context) error {
	for {
		job, err := w.queue.Pop(ctx)
		if err != nil {
			if ctx.Err() != nil {
				return nil
			}
			return fmt.Errorf("pop job: %w", err)
		}

		if err := w.dispatcher.Dispatch(ctx, job); err != nil {
			return fmt.Errorf("dispatch job %q: %w", job.Id(), err)
		}
	}
}
