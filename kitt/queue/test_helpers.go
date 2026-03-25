package queue

import "context"

type someJob struct {
	SomeVar string
}

func (j someJob) Id() string {
	return "some.job"
}

type someJobHandler struct {
	SomeVar string
}

func (h *someJobHandler) Handle(ctx context.Context, job Job) error {
	j, ok := job.(someJob)

	if ok {
		h.SomeVar = j.SomeVar
	}
	return nil
}
