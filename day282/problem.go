package day282

import "errors"

var errNoAnswer = errors.New("no pythagorean triplet exists")

// ErrNoAnswer is the error returned if a solution doesn't exist.
func ErrNoAnswer() error {
	return errNoAnswer
}

// PythagoreanTriplet returns the triplets that fulfill
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
