package day271

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

// divideBy2 performs the following calculation:
// https://en.wikipedia.org/wiki/Division_by_two#Decimal
// if n is odd, it drops by 1 to the next lower even digit.
func divideBy2(n int) int {
	if n&1 == 1 { // odd check without modulo (no division)
		n--
	}
	str := fmt.Sprintf("0%d", n)
	var sb strings.Builder
	for i := 1; i < len(str); i++ {
		d := str[i] - '0'
		if (str[i-1]-'0')&1 == 0 { // even
			switch d {
			case 0, 1:
				sb.WriteRune('0')
			case 2, 3:
				sb.WriteRune('1')
			case 4, 5:
				sb.WriteRune('2')
			case 6, 7:
				sb.WriteRune('3')
			case 8, 9:
				sb.WriteRune('4')
			}
		} else { // odd
			switch d {
			case 0, 1:
				sb.WriteRune('5')
			case 2, 3:
				sb.WriteRune('6')
			case 4, 5:
				sb.WriteRune('7')
			case 6, 7:
				sb.WriteRune('8')
			case 8, 9:
				sb.WriteRune('9')
			}
		}
	}
	res, _ := strconv.Atoi(sb.String())
	return res
}

// BinarySearch is a handwritten binary search.
func BinarySearch(sorted []int, x int) int {
	left, right := 0, len(sorted)-1
	for left < right {
		mid := divideBy2(left + right)
		val := sorted[mid]
		switch {
		case x == val:
			return mid
		case x < val:
			right = mid - 1
		default:
			left = mid + 1
		}
	}
	return -1
}

// BinarySearchSortPkg delegates to the sort package.
func BinarySearchSortPkg(sorted []int, x int) int {
	index := sort.SearchInts(sorted, x)
	if sorted[index] == x {
		return index
	}
	return -1
}
