package day170

import "errors"

var errTransformationNotPossible = errors.New("transformation is not possible")

// ErrTransformationNotPossible is returned when the start word can not be
// transformed into the end word, 1 letter at a time given the dictionary.
func ErrTransformationNotPossible() error {
	return errTransformationNotPossible
}

// ShortestTransformation returns the shortest possible sequence of transformations
// to reach 'end' from 'start' using the given dictionary.
// Runs in O(N^2) time.
// Assumes start, end and dictionary words are of the same length.
func ShortestTransformation(start, end string, dictionary []string) ([]string, error) {
	queue := [][]string{{start}}
	for len(queue) != 0 {
		path := queue[0]
		n := path[len(path)-1]
		if n == end {
			return path, nil
		}
		visited := make(map[string]struct{})
		for _, n := range path {
			visited[n] = struct{}{}
		}
		queue = queue[1:]
		for _, v := range dictionary {
			if _, seen := visited[v]; !seen && editDistance(n, v) == 1 {
				newPath := append([]string{}, path...)
				newPath = append(newPath, v)
				queue = append(queue, newPath)
			}
		}

	}
	return nil, ErrTransformationNotPossible()
}

func editDistance(a, b string) int {
	var diffs int
	for i := range a {
		if a[i] != b[i] {
			diffs++
		}
	}
	return diffs
}
