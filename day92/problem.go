package day92

// CourseOrder returns how to take all the courses.
// If it's not possible to take all the courses, returns nil.
// Runs in O(N) time.
func CourseOrder(prereq map[string][]string) []string {
	var result []string
	marked := make(map[string]int)
	for n := range prereq {
		if _, found := marked[n]; !found {
			newest := visit(n, prereq, marked, result)
			if len(newest) > len(result) {
				result = newest
			}
		}
	}
	if len(marked) != len(prereq) {
		return nil
	}
	reverse(result)
	return result
}

func reverse(n []string) {
	for i := 0; i < len(n)/2; i++ {
		n[i], n[len(n)-1-i] = n[len(n)-1-i], n[i]
	}
}

func visit(node string, graph map[string][]string,
	marked map[string]int, result []string) []string {
	if mark := marked[node]; mark != 0 {
		return nil
	}
	marked[node] = 1
	for _, m := range graph[node] {
		newest := visit(m, graph, marked, result)
		if len(newest) > len(result) {
			result = newest
		}
	}
	marked[node] = 2
	result = append([]string{node}, result...)
	return result
}
