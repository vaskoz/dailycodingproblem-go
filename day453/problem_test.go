package day453

import (
	"reflect"
	"testing"
)

// nolint
var testcases = []struct {
	head     *BST
	k        int
	expected []int
}{
	{&BST{10, &BST{5, nil, nil}, &BST{
		15, &BST{11, nil, nil}, &BST{15, nil, nil},
	}}, 20, []int{5, 15}},
	{&BST{10, &BST{5, nil, nil}, &BST{
		15, &BST{11, nil, nil}, &BST{15, nil, nil},
	}}, 100, []int{}},
	{&BST{10, &BST{5, nil, nil}, &BST{
		15, &BST{11, nil, nil}, &BST{15, nil, nil},
	}}, 15, []int{5, 10}},
	{&BST{10, &BST{5, nil, nil}, &BST{
		15, &BST{11, nil, nil}, &BST{15, nil, nil},
	}}, 26, []int{11, 15}},
}

func TestFindTwoSumKInBST(t *testing.T) {
	t.Parallel()

	for _, tc := range testcases {
		if result := FindTwoSumKInBST(tc.head, tc.k); !reflect.DeepEqual(result, tc.expected) {
			t.Errorf("Expected %v got %v", tc.expected, result)
		}
	}
}

func BenchmarkFindTwoSumInBST(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range testcases {
			FindTwoSumKInBST(tc.head, tc.k)
		}
	}
}
