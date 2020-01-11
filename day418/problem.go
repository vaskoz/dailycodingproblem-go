package day418

// Bonuses determines how much bonus to give based on lines of code.
func Bonuses(loc []int) []int {
	if len(loc) == 0 {
		return nil
	} else if len(loc) == 1 {
		return []int{1}
	}

	bonuses := make([]int, 0, len(loc))
	segs := segments(loc)

	for _, seg := range segs {
		asc, run := seg.ascending, seg.run
		segBonuses := make([]int, run)

		for i := range segBonuses {
			segBonuses[i] = i
		}

		if !asc {
			for i := 0; i < len(segBonuses)/2; i++ {
				segBonuses[i], segBonuses[len(segBonuses)-1-i] = segBonuses[len(segBonuses)-1-i], segBonuses[i]
			}
		}

		bonuses = append(bonuses, segBonuses...)
	}

	for i := range bonuses {
		bonuses[i]++
	}

	return bonuses
}

type segment struct {
	ascending bool
	run       int
}

func segments(loc []int) []segment {
	asc := loc[1] > loc[0]
	prev := loc[0]

	var start int

	var segs []segment

	slice := loc[1:]
	for i, num := range slice {
		if (asc && num < prev) || (!asc && num > prev) {
			segs = append(segs, segment{asc, i - start + 1})
			start = i + 1
			asc = !asc
		}

		prev = num
	}

	return append(segs, segment{asc, len(loc) - start})
}
