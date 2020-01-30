package day434

import "testing"

// nolint
var testcases = []struct {
	root             *BinaryTree
	target           int
	expectedFloor    int
	expectedFloorErr error
	expectedCeil     int
	expectedCeilErr  error
}{
	{
		&BinaryTree{
			11,
			&BinaryTree{9, nil, nil},
			&BinaryTree{12, nil, nil},
		},
		10,
		9,
		nil,
		11,
		nil,
	},
	{
		&BinaryTree{
			11,
			&BinaryTree{9, nil, nil},
			&BinaryTree{12, nil, nil},
		},
		12,
		12,
		nil,
		12,
		nil,
	},
	{
		&BinaryTree{
			11,
			&BinaryTree{9, nil, nil},
			&BinaryTree{12, nil, nil},
		},
		13,
		12,
		nil,
		0,
		ErrNoAnswer(),
	},
	{
		&BinaryTree{
			11,
			&BinaryTree{9, nil, nil},
			&BinaryTree{12, nil, nil},
		},
		8,
		0,
		ErrNoAnswer(),
		9,
		nil,
	},
	{
		&BinaryTree{
			12,
			&BinaryTree{9, nil, nil},
			&BinaryTree{14, nil, nil},
		},
		11,
		9,
		nil,
		12,
		nil,
	},
}

func TestFloorRecursive(t *testing.T) {
	t.Parallel()

	for tcid, tc := range testcases {
		if result, err := FloorRecursive(tc.root, tc.target); result != tc.expectedFloor ||
			err != tc.expectedFloorErr {
			t.Errorf("TCID %d Expected (%v,%v), got (%v,%v)", tcid, tc.expectedFloor, tc.expectedFloorErr, result, err)
		}
	}
}

func BenchmarkFloorRecursive(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range testcases {
			FloorRecursive(tc.root, tc.target) // nolint
		}
	}
}

func TestFloorFaster(t *testing.T) {
	t.Parallel()

	for tcid, tc := range testcases {
		if result, err := FloorFaster(tc.root, tc.target); result != tc.expectedFloor ||
			err != tc.expectedFloorErr {
			t.Errorf("TCID %d Expected (%v,%v), got (%v,%v)", tcid, tc.expectedFloor, tc.expectedFloorErr, result, err)
		}
	}
}

func BenchmarkFloorFaster(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range testcases {
			FloorRecursive(tc.root, tc.target) // nolint
		}
	}
}

func TestCeilingRecursive(t *testing.T) {
	t.Parallel()

	for tcid, tc := range testcases {
		if result, err := CeilingRecursive(tc.root, tc.target); result != tc.expectedCeil ||
			err != tc.expectedCeilErr {
			t.Errorf("TCID %d Expected (%v,%v), got (%v,%v)", tcid, tc.expectedCeil, tc.expectedCeilErr, result, err)
		}
	}
}

func BenchmarkCeiling(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range testcases {
			CeilingRecursive(tc.root, tc.target) // nolint
		}
	}
}

func TestCeilingFaster(t *testing.T) {
	t.Parallel()

	for tcid, tc := range testcases {
		if result, err := CeilingFaster(tc.root, tc.target); result != tc.expectedCeil ||
			err != tc.expectedCeilErr {
			t.Errorf("TCID %d Expected (%v,%v), got (%v,%v)", tcid, tc.expectedCeil, tc.expectedCeilErr, result, err)
		}
	}
}

func BenchmarkCeilingFaster(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range testcases {
			CeilingFaster(tc.root, tc.target) // nolint
		}
	}
}
