package queue

import "context"

type JobHandler interface {
	Handle(ctx context.Context, job Job) error
}
