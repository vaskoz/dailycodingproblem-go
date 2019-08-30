package day356

import "testing"

type action struct {
	act  string
	data interface{}
	err  error
}

// nolint
var testcases = []struct {
	actions []action
}{
	{
		[]action{
			{"E", 100, nil}, {"S", 1, nil}, {"D", 100, nil},
			{"D", nil, ErrEmpty()},
		},
	},
	{
		[]action{
			{"E", 100, nil}, {"E", 200, nil}, {"E", 300, nil}, {"E", 400, nil}, {"E", 500, nil},
			{"S", 5, nil}, {"E", 1000, ErrFull()},
			{"D", 100, nil}, {"D", 200, nil},
			{"S", 3, nil},
			{"D", 300, nil}, {"D", 400, nil},
			{"S", 1, nil},
			{"E", 600, nil}, {"E", 700, nil}, {"E", 800, nil}, {"E", 900, nil},
			{"E", 1000, ErrFull()},
			{"D", 500, nil}, {"D", 600, nil}, {"D", 700, nil}, {"D", 800, nil},
			{"S", 1, nil},
			{"E", 1000, nil}, {"E", 1100, nil}, {"E", 1200, nil}, {"E", 1300, nil},
			{"S", 5, nil},
			{"D", 900, nil}, {"D", 1000, nil}, {"D", 1100, nil}, {"D", 1200, nil}, {"D", 1300, nil},
			{"S", 0, nil},
			{"D", nil, ErrEmpty()},
		},
	},
}

func TestQueueFLA(t *testing.T) {
	t.Parallel()
	for _, tc := range testcases {
		q := NewQueueFLA(5)
		for _, action := range tc.actions {
			switch action.act {
			case "E":
				if err := q.Enqueue(action.data); err != action.err {
					t.Errorf("Expected %v, got %v", action.err, err)
				}
			case "D":
				if item, err := q.Dequeue(); item != action.data || err != action.err {
					t.Errorf("Expected (%v,%v), got (%v,%v)", action.data, action.err, item, err)
				}
			case "S":
				if s := q.Size(); s != action.data {
					t.Errorf("Expected a size of %d, got %d", action.data, s)
				}
			}
		}
	}
}

func BenchmarkQueueFLA(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range testcases {
			q := NewQueueFLA(5)
			for _, action := range tc.actions {
				switch action.act {
				case "E":
					q.Enqueue(action.data) // nolint
				case "D":
					q.Dequeue() // nolint
				case "S":
					q.Size()
				}
			}
		}
	}
}
