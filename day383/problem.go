package day383

import (
	"sort"
	"strings"
)

type interval struct {
	start, end int
}

// Embolden surrounds substrings (lst) in <b> tags
// and combines overlapping substrings.
func Embolden(s string, lst []string) string {
	intervals := make([]interval, 0, len(lst))

	for i := range lst {
		pos := 0
		for j := strings.Index(s[pos:], lst[i]); j != -1; j = strings.Index(s[pos:], lst[i]) {
			pos += j
			intervals = append(intervals, interval{pos, pos + len(lst[i])})
			pos += len(lst[i])
		}
	}

	sort.Slice(intervals, func(i, j int) bool {
		if intervals[i].start < intervals[j].start {
			return true
		} else if intervals[i].start == intervals[j].start {
			return intervals[i].end < intervals[j].end
		}
		return false
	})

	intervals = mergeIntervals(intervals)

	var sb strings.Builder

	pos := 0

	for _, inter := range intervals {
		sb.WriteString(s[pos:inter.start])
		sb.WriteString("<b>")
		sb.WriteString(s[inter.start:inter.end])
		sb.WriteString("</b>")

		pos = inter.end
	}

	sb.WriteString(s[pos:])

	return sb.String()
}

func mergeIntervals(ints []interval) []interval {
	merged := make([]interval, 0, len(ints))
	merged = append(merged, ints[0])
	last := 0

	for i := 1; i < len(ints); i++ {
		if ints[i].start > merged[last].start && ints[i].start < merged[last].end {
			merged[last] = interval{min(merged[last].start, ints[i].start),
				max(merged[last].end, ints[i].end)}
		} else {
			merged = append(merged, ints[i])
			last++
		}
	}

	return merged
}

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}

	return b
}
