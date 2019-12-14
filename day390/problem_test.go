package day390

import (
	"math/rand"
	"reflect"
	"testing"
	"time"
)

func TestMissingThousandInMillion(t *testing.T) {
	t.Parallel()

	missing := []int{10_395, 25_011, 345_335, 565_332, 999_993}
	input := setup(missing)

	if res := MissingThousandInMillion(input); !reflect.DeepEqual(res, missing) {
		t.Errorf("Expected %v, got %v", missing, res)
	}
}

func BenchmarkMissingThousandInMillion(b *testing.B) {
	missing := []int{25_011, 345_335, 565_332, 999_993, 10_395}
	input := setup(missing)

	for i := 0; i < b.N; i++ {
		MissingThousandInMillion(input)
	}
}

func setup(miss []int) []int {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	ans := make([]int, 1_000_000-len(miss)+1)
	missMap := make(map[int]struct{}, len(miss))

	for _, n := range miss {
		missMap[n] = struct{}{}
	}

	pos := 0

	for n := 1; n <= 1_000_000; n++ {
		if _, skip := missMap[n]; !skip {
			ans[pos] = n
			pos++
		}
	}

	r.Shuffle(len(ans), func(i, j int) {
		ans[i], ans[j] = ans[j], ans[i]
	})

	return ans
}
