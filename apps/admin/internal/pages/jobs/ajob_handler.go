package jobs

import (
	"context"
	"fmt"
	"kitt/kitt/queue"
)

type AjobHandler struct {
}

func (h AjobHandler) Handle(ctx context.Context, job queue.Job) error {
	j, ok := job.(Ajob)
	if ok {
		fmt.Println("SomeVariable: " + j.SomeVar)
	}
	return nil
}
