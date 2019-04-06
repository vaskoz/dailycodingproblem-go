package day105

import "time"

// Debounced returns a debounced function that will not call its
// provided function until after N milliseconds despite calling it.
func Debounced(f func(), n int) func() {
	created := time.Now()
	return func() {
		now := time.Now()
		if elapsed := int(now.Sub(created) / time.Millisecond); elapsed > n {
			f()
		}
	}
}
