package day292

import (
	"errors"
	"sort"
)

var errTeamNotPossible = errors.New("building a team without enemies is impossible")

// ErrTeamNotPossible returns the error when a team isn't possible.
func ErrTeamNotPossible() error {
	return errTeamNotPossible
}

// AssignTeams attempts to assign teammates using a constraint of enemies.
// It uses Breadth-First-Search to accomplish this.
func AssignTeams(enemies map[int]map[int]struct{}) ([]int, []int, error) {
	team1 := make(map[int]bool)
	team2 := make(map[int]bool)
	var q []int
	for id := range enemies {
		team1[id] = true
		for e := range enemies[id] {
			q = append(q, e)
		}
		break
	}
	for len(q) != 0 {
		var nextQ []int
		var countNew int
		for _, id := range q {
			if _, found := team1[id]; found {
				return nil, nil, errTeamNotPossible
			}
			if _, found := team2[id]; !found {
				team2[id] = true
				countNew++
				for enemy := range enemies[id] {
					nextQ = append(nextQ, enemy)
				}
			}
		}
		if countNew > 0 {
			q = nextQ
		} else if countNew == 0 {
			q = nil
		}
		team1, team2 = team2, team1
	}
	return buildResult(team1, team2)
}

func buildResult(team1, team2 map[int]bool) ([]int, []int, error) {
	result1 := make([]int, 0, len(team1))
	for id := range team1 {
		result1 = append(result1, id)
	}
	result2 := make([]int, 0, len(team2))
	for id := range team2 {
		result2 = append(result2, id)
	}
	sort.Slice(result1, func(i, j int) bool { return result1[i] < result1[j] })
	sort.Slice(result2, func(i, j int) bool { return result2[i] < result2[j] })
	if result1[0] < result2[0] {
		return result1, result2, nil
	}
	return result2, result1, nil
}
