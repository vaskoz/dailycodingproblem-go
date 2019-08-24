package day365

import "testing"

type step struct {
	action string
	data   interface{}
	err    error
}

// nolint
var testcases = []struct {
	steps []step
}{
	{
		[]step{
			{"push", "foo", nil},
			{"push", "bcd", nil},
			{"push", "def", nil},
			{"pop", "def", nil},
			{"pop", "bcd", nil},
			{"pop", "foo", nil},
			{"pop", nil, ErrEmptyStack()},
			{"pull", nil, ErrEmptyStack()},
		},
	},
	{
		[]step{
			{"push", "abc", nil},
			{"push", "bcd", nil},
			{"pull", "abc", nil},
			{"push", "def", nil},
			{"pop", "def", nil},
			{"pop", "bcd", nil},
		},
	},
	{
		[]step{
			{"push", "abc", nil},
			{"push", "bcd", nil},
			{"push", "def", nil},
			{"push", "efg", nil},
			{"push", "fgi", nil},
			{"pull", "abc", nil},
			{"pop", "fgi", nil},
			{"pop", "efg", nil},
			{"pop", "def", nil},
			{"pop", "bcd", nil},
			{"pull", nil, ErrEmptyStack()},
			{"pop", nil, ErrEmptyStack()},
		},
	},
}

func TestQuack(t *testing.T) {
	t.Parallel()
	for _, tc := range testcases {
		q := &Quack{}
		for _, s := range tc.steps {
			switch s.action {
			case "push":
				q.Push(s.data)
			case "pop":
				if x, err := q.Pop(); x != s.data || err != s.err {
					t.Errorf("Expected (%v,%v), got (%v,%v)", s.data, s.err, x, err)
				}
			case "pull":
				if x, err := q.Pull(); x != s.data || err != s.err {
					t.Errorf("Expected (%v,%v), got (%v,%v)", s.data, s.err, x, err)
				}
			}
		}
	}
}

func BenchmarkQuack(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range testcases {
			q := &Quack{}
			for _, s := range tc.steps {
				switch s.action {
				case "push":
					q.Push(s.data)
				case "pop":
					q.Pop() // nolint
				case "pull":
					q.Pull() // nolint
				}
			}
		}
	}
}
