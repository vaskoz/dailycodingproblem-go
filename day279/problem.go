package day279

// FriendGraph is an adjacency list of friendships using
// integer IDs for friends.
type FriendGraph map[int]map[int]struct{}

// NumberOfFriendGroups returns the number of disconnected
// friend groups.
// Runs DFS for each separate group and counts the number of
// times DFS had to be run.
// Runs in O(V+E) time and O(V) space just like DFS.
func NumberOfFriendGroups(friends FriendGraph) int {
	visits := make([]bool, len(friends))
	friendGroups := 0
	for id, visited := range visits {
		if !visited {
			dfs(friends, id, visits)
			friendGroups++
		}
	}
	return friendGroups
}

func dfs(g FriendGraph, start int, visits []bool) {
	for friendID := range g[start] {
		if visited := visits[friendID]; !visited {
			visits[friendID] = true
			dfs(g, friendID, visits)
		}
	}
}
