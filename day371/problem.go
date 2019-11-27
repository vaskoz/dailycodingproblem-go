package day371

import (
	"errors"
	"sort"
	"strconv"
	"strings"
)

var errNoAnswerPossible = errors.New("an answer doesn't exist")

type eqData struct {
	vars  map[string]int
	total int
}

// SolveArithmeticEquations returns the answer to these equations
// or an error if no answer exists.
// Will panic on malformed input.
func SolveArithmeticEquations(equations string) (map[string]int, error) {
	eqs := strings.Split(strings.TrimSpace(equations), "\n")
	eqDatas := make([]eqData, 0, len(eqs))

	for _, eq := range eqs {
		sides := strings.Split(strings.TrimSpace(eq), "=")
		if len(sides) != 2 {
			panic("only one equal sign per line/equation")
		}

		lhs, rhs := strings.TrimSpace(sides[0]), strings.TrimSpace(sides[1])
		vars := make(map[string]int)
		total := 0

		for _, v := range strings.Split(lhs, "+") {
			if i, err := strconv.Atoi(v); err == nil {
				total -= i
			} else {
				vars[v]++
			}
		}

		for _, v := range strings.Split(rhs, "+") {
			v = strings.TrimSpace(v)
			if i, err := strconv.Atoi(v); err == nil {
				total += i
			} else {
				vars[v]--
			}
		}

		eqDatas = append(eqDatas, eqData{vars, total})
	}

	return solveArithmeticEquations(eqDatas)
}

func solveArithmeticEquations(equations []eqData) (map[string]int, error) {
	ans := make(map[string]int)

	sort.Slice(equations, func(i, j int) bool {
		return len(equations[i].vars) < len(equations[j].vars)
	})

	for _, eq := range equations {
		for v, coef := range eq.vars {
			if a, exists := ans[v]; exists {
				eq.total -= a * coef
				delete(eq.vars, v)
			}
		}

		if len(eq.vars) != 1 {
			return nil, errNoAnswerPossible
		}

		for k, v := range eq.vars {
			ans[k] = eq.total / v
		}
	}

	return ans, nil
}
