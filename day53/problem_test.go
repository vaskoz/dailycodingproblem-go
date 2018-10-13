package day53

import "testing"

var testcases = []struct {
	inserts []interface{}
}{
	{[]interface{}{"foo", "bar", 1, 3, 5, 6}},
}

func TestQueue(t *testing.T) {
	t.Parallel()
	for _, tc := range testcases {
		q := NewQueue()
		for _, v := range tc.inserts {
			q.Enqueue(v)
		}
		for _, v := range tc.inserts {
			if res, err := q.Dequeue(); err != nil {
				t.Errorf("Expected no errors while dequeuing")
			} else if res != v {
				t.Errorf("Expected results in FIFO order. Expected %v, got %v", v, res)
			}
		}
		if _, err := q.Dequeue(); err == nil {
			t.Errorf("Dequeue on empty queue should return an error")
		}
	}
}

func BenchmarkQueue(b *testing.B) {
}
