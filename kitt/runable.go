package kitt

import "context"

type Runnable interface {
	Id() string
	Run(ctx context.Context) error
}
