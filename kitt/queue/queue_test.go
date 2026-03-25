package queue

import "testing"

func Test_In_Memory_Queue(t *testing.T) {
	t.Run("it adds to the queue and pop", func(t *testing.T) {
		want := someJob{SomeVar: "foo"}
		q := NewInMemoryQueue(100)
		q.Dispatch(t.Context(), want)

		have, err := q.Pop(t.Context())

		if err != nil {
			t.Fatal(err)
		}

		if have != want {
			t.Fatalf("not eq: %s -> %s", have, want)
		}
	})
}
