package queue

import (
	"context"
)

type Queue interface {
	Dispatch(ctx context.Context, job Job) error
	Pop(ctx context.Context) (Job, error)
}
