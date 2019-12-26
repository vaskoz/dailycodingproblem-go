package day329

// Matches is a mapping from key=male to value=female.
type Matches map[string]string

// StableMarriage returns N stable marriage pairs.
func StableMarriage(males, females map[string][]string) Matches {
	free := make(map[string]struct{}, len(males)*2)
	proposals := make(map[string]int, len(males)*2)
	engaged := make(map[string]string, len(males)*2)

	for person := range males {
		free[person] = struct{}{}
		proposals[person] = 0
	}

	for person := range females {
		free[person] = struct{}{}
	}

	for len(proposals) != 0 {
		var m string

		for man := range proposals {
			if _, ok := free[man]; ok {
				m = man
				break
			}
		}

		if m == "" {
			break
		}

		w := males[m][proposals[m]]

		if _, ok := free[w]; ok {
			engaged[m] = w
			engaged[w] = m
			delete(free, m)
			delete(free, w)
		} else {
			mp := engaged[w]

			for _, suitor := range females[w] {
				if suitor == m {
					delete(engaged, mp)
					engaged[w] = m
					engaged[m] = w
					delete(free, m)
					free[mp] = struct{}{}
					break
				} else if suitor == mp {
					break
				}
			}
		}
		proposals[m]++
	}

	for female := range females {
		delete(engaged, female)
	}

	return engaged
}
