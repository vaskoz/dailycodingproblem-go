package day275

import (
	"fmt"
	"strings"
)

func LookAndSay(n int) string {
	if n < 1 {
		panic("look and say number must be 1 or greater")
	}
	start := "1"
	for i := 1; i < n; i++ {
		var parts []string
		var c rune
		count := 0
		for i, r := range start {
			switch {
			case i == 0:
				c = r
				count = 1
			case c != r:
				parts = append(parts, fmt.Sprintf("%d%s", count, string(c)))
				c = r
				count = 1
			default:
				count++
			}
		}
		parts = append(parts, fmt.Sprintf("%d%s", count, string(c)))
		start = strings.Join(parts, "")
	}
	return start
}
