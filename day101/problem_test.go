package day101

import (
	"reflect"
	"testing"
)

var testcases = []struct {
	even     int
	expected []int
}{
	{4, []int{2, 2}},
	{6, []int{3, 3}},
	{8, []int{3, 5}},
	{10, []int{3, 7}},
	{10, []int{3, 7}},
	{18, []int{5, 13}},
}

func TestSieve(t *testing.T) {
	t.Parallel()
	expected := []int{2, 3, 5, 7, 11, 13, 17, 19, 23, 29, 31, 37,
		41, 43, 47, 53, 59, 61, 67, 71, 73, 79, 83, 89, 97}

	if result := Sieve(100); !reflect.DeepEqual(expected, result) {
		t.Errorf("Expected %v got %v", expected, result)
	}
}

func BenchmarkSieve(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Sieve(100)
	}
}

func TestGoldbachsConjecture(t *testing.T) {
	t.Parallel()
	for _, tc := range testcases {
		if result := GoldbachsConjecture(tc.even); !reflect.DeepEqual(tc.expected, result) {
			t.Errorf("Expected %v got %v", tc.expected, result)
		}
	}
}

func BenchmarkGoldbachsConjecture(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range testcases {
			GoldbachsConjecture(tc.even)
		}
	}
}

func TestGoldbachsConjectureBadInput(t *testing.T) {
	t.Parallel()
	if result := GoldbachsConjecture(3); result != nil {
		t.Errorf("Parameters less than 4 should result in nil return")
	}
	if result := GoldbachsConjecture(9); result != nil {
		t.Errorf("Odd parameters should result in nil return")
	}

}
