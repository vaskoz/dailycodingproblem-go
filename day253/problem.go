package day253

import (
	"fmt"
	"io"
	"os"
	"strings"
)

var (
	// nolint
	out io.Writer = os.Stdout
)

// PrintZigZag prints k lines of a zig-zag pattern for
// the letters given.
func PrintZigZag(letters string, k int) {
	lines := make([]strings.Builder, k)
	pos := 0
	goingUp := false
	for _, r := range letters {
		switch {
		case pos == k:
			goingUp = true
			pos -= 2
		case pos == -1:
			goingUp = false
			pos += 2
		}
		for i := range lines {
			if i == pos {
				lines[i].WriteRune(r)
			} else {
				lines[i].WriteRune(' ')
			}
		}
		switch {
		case goingUp:
			pos--
		case !goingUp:
			pos++
		}
	}
	for _, line := range lines {
		fmt.Fprintln(out, line.String())
	}
}
