package queue

import (
	"context"
	"testing"
	"time"
)

func Test_Queue_Worker(t *testing.T) {
	t.Run("it", func(t *testing.T) {
		ctx, cancel := context.WithCancel(context.Background())
		defer cancel()

		job := someJob{SomeVar: "foo"}
		handler := &someJobHandler{}

		d := NewJobDispatcher()
		q := NewInMemoryQueue(10)
		worker := NewQueueWorker("queue", q, d)

		d.Register(someJob{}, handler)
		q.Dispatch(ctx, job)

		done := make(chan error, 1)
		go func() {
			done <- worker.Run(ctx)
		}()

		time.Sleep(time.Millisecond * 20)
		cancel()

		have := handler.SomeVar
		want := job.SomeVar

		if have != want {
			t.Fatalf("not eq: %s -> %s", have, want)
		}
	})
}
