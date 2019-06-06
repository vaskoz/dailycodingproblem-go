package day274

import (
	"strings"
)

// EvalSimpleMathExpression parses very simple math expressions.
// Expr consisting of parentheses, single digits, and positive and negative signs.
func EvalSimpleMathExpression(expr string) int {
	ex := trimAllSpaces(expr)
	return simpleMath([]rune(ex))
}

func simpleMath(expr []rune) int {
	var result int
	for i := 0; i < len(expr); i++ {
		switch r := expr[i]; r {
		case '0', '1', '2', '3', '4', '5', '6', '7', '8', '9':
			if i > 0 && expr[i-1] == '-' {
				result -= int(r - '0')
			} else {
				result += int(r - '0')
			}
		case '(':
			count := 1
			j := i + 1
			for count != 0 {
				r = expr[j]
				switch r {
				case '(':
					count++
				case ')':
					count--
				}
				j++
			}
			if i > 0 && expr[i-1] == '-' {
				result -= simpleMath(expr[i+1 : j-1])
			} else {
				result += simpleMath(expr[i+1 : j-1])
			}
			i = j
		}
	}
	return result
}

func trimAllSpaces(expr string) string {
	var sb strings.Builder
	for _, r := range expr {
		if r != ' ' {
			sb.WriteRune(r)
		}
	}
	return sb.String()
}
