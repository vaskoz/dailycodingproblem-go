package day282

import (
	"errors"
	"math"
)

var errNoAnswer = errors.New("no pythagorean triplet exists")

// ErrNoAnswer is the error returned if a solution doesn't exist.
func ErrNoAnswer() error {
	return errNoAnswer
}

// PythagoreanTripletBrute returns the triplets that fulfill
// a^2+b^2=c^2.
// Runs in O(N^3) time and O(1) space.
func PythagoreanTripletBrute(nums []int) (a, b, c int, err error) {
	for i := range nums {
		for j := range nums {
			for k := range nums {
				if i == j || j == k || i == k {
					continue
				}
				if nums[i]*nums[i]+nums[j]*nums[j] == nums[k]*nums[k] {
					if nums[i] < nums[j] {
						return nums[i], nums[j], nums[k], nil
					}
					return nums[j], nums[i], nums[k], nil
				}
			}
		}
	}
	return 0, 0, 0, errNoAnswer
}

// PythagoreanTripletBrute returns the triplets that fulfill
// a^2+b^2=c^2.
// Runs in O(N^2) time and O(N) space.
func PythagoreanTriplet(nums []int) (a, b, c int, err error) {
	squared := make(map[int]struct{}, len(nums))
	for _, num := range nums {
		squared[num*num] = struct{}{}
	}
	for i := range nums {
		for j := range nums {
			total := nums[i]*nums[i] + nums[j]*nums[j]
			c := int(math.Sqrt(float64(total)))
			if _, found := squared[total]; found {
				if nums[i] < nums[j] {
					return nums[i], nums[j], c, nil
				}
				return nums[j], nums[i], c, nil
			}
		}
	}
	return 0, 0, 0, errNoAnswer
}
