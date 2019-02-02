package day163

import "strconv"

// ReversePolishCalculator calculates the result of the expression
// given in reverse polish notation.
// It assumes the given expression is always valid.
// Bad input causes a panic.
func ReversePolishCalculator(input []string) int {
	var stack []int
	for _, v := range input {
		switch v {
		case "+":
			r := stack[len(stack)-2] + stack[len(stack)-1]
			stack = append(stack[:len(stack)-2], r)
		case "-":
			r := stack[len(stack)-2] - stack[len(stack)-1]
			stack = append(stack[:len(stack)-2], r)
		case "*":
			r := stack[len(stack)-2] * stack[len(stack)-1]
			stack = append(stack[:len(stack)-2], r)
		case "/":
			r := stack[len(stack)-2] / stack[len(stack)-1]
			stack = append(stack[:len(stack)-2], r)
		default:
			i, err := strconv.Atoi(v)
			if err != nil {
				panic("invalid input")
			}
			stack = append(stack, i)
		}
	}
	if len(stack) != 1 {
		panic("invalid input")
	}
	return stack[0]
}
