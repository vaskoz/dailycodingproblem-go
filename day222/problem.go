package day222

import (
	"fmt"
	"strings"
)

// ShortestStandardizedPath returns the shortest standardized path.
func ShortestStandardizedPath(path string) string {
	parts := strings.Split(path, "/")
	var result []string
	for _, part := range parts {
		if part == ".." {
			result = result[:len(result)-1]
		} else if part == "." {
			continue
		} else {
			result = append(result, part)
		}
	}
	return fmt.Sprintf("%s", strings.Join(result, "/"))
}
