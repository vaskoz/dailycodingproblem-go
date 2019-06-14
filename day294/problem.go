package day294

import (
	"errors"
)

// Elevation is a map of locations to their elevations.
type Elevation map[int]int

// PathDistance represents the distance for a particular path.
type PathDistance map[int]map[int]int

// MaxInt is the largest int.
const MaxInt = int(^uint(0) >> 1)

var errNotPossible = errors.New("no path possible")

// ErrNotPossible is the error returned if a path doesn't exist.
func ErrNotPossible() error {
	return errNotPossible
}

// ShortestRunnerPath with the condition that the route goes entirely
// uphill at first, and then entirely downhill.
// Always starts at location 0 and end at location 0.
func ShortestRunnerPath(elevations Elevation, paths PathDistance) (int, error) {
	min := MaxInt
	startElev := elevations[0]
	for next, distance := range paths[0] {
		nextElev := elevations[next]
		if nextElev > startElev {
			dist, err := shortestRunnerPath(elevations, paths, map[int]struct{}{}, true, next)
			if err == nil && dist+distance < min {
				min = dist + distance
			}
		}
	}
	if min == MaxInt {
		return 0, ErrNotPossible()
	}
	return min, nil
}

func shortestRunnerPath(
	elevations Elevation,
	paths PathDistance,
	visited map[int]struct{},
	up bool,
	n int) (int, error) {
	if n == 0 {
		return 0, nil
	}
	min := MaxInt
	currentElevation := elevations[n]
	for next, distance := range paths[n] {
		if _, visit := visited[next]; !visit {
			nextElevation := elevations[next]
			if up && nextElevation > currentElevation {
				visited[next] = struct{}{}
				dist, err := shortestRunnerPath(elevations, paths, visited, true, next)
				totalDist := dist + distance
				if err == nil && totalDist < min {
					min = totalDist
				}
				delete(visited, next)
			} else if nextElevation < currentElevation {
				visited[next] = struct{}{}
				dist, err := shortestRunnerPath(elevations, paths, visited, false, next)
				totalDist := dist + distance
				if err == nil && totalDist < min {
					min = totalDist
				}
				delete(visited, next)
			}
		}
	}
	if min == MaxInt {
		return 0, ErrNotPossible()
	}
	return min, nil
}
