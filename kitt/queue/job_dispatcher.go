package queue

import (
	"context"
	"fmt"
)

type JobDispatcher interface {
	Register(job Job, handler JobHandler)
	Dispatch(ctx context.Context, job Job) error
}

type jobDispatcher struct {
	handlers map[string]JobHandler
}

func NewJobDispatcher() JobDispatcher {
	return &jobDispatcher{
		handlers: make(map[string]JobHandler),
	}
}

func (d *jobDispatcher) Register(job Job, handler JobHandler) {
	d.handlers[job.Id()] = handler
}

func (d *jobDispatcher) Dispatch(ctx context.Context, job Job) error {
	handler, ok := d.handlers[job.Id()]
	if !ok {
		return fmt.Errorf("no handler registered for job %q", job.Id())
	}

	if err := handler.Handle(ctx, job); err != nil {
		return fmt.Errorf("job %q failed: %w", job.Id(), err)
	}

	return nil
}
