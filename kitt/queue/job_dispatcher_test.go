package queue

import (
	"testing"
)

func Test_Job_Dispatcher(t *testing.T) {
	t.Run("it dispatches", func(t *testing.T) {
		job := someJob{SomeVar: "foo"}
		handler := &someJobHandler{}
		d := NewJobDispatcher()
		d.Register(someJob{}, handler)
		d.Dispatch(t.Context(), job)

		have := handler.SomeVar
		want := job.SomeVar

		if have != want {
			t.Fatalf("not equal: %s -> %s", have, want)
		}
	})
}
