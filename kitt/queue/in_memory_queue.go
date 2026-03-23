package queue

import (
	"context"
	"fmt"
)

type InMemoryQueue struct {
	jobs chan Job
}

func NewInMemoryQueue(buffer int) *InMemoryQueue {
	return &InMemoryQueue{
		jobs: make(chan Job, buffer),
	}
}

func (q *InMemoryQueue) Dispatch(ctx context.Context, job Job) error {
	select {
	case <-ctx.Done():
		return ctx.Err()
	case q.jobs <- job:
		return nil
	}
}

func (q *InMemoryQueue) Pop(ctx context.Context) (Job, error) {
	select {
	case <-ctx.Done():
		return nil, ctx.Err()
	case job := <-q.jobs:
		if job == nil {
			return nil, fmt.Errorf("nil job received")
		}
		return job, nil
	}
}
