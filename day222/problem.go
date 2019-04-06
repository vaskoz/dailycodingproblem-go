package day222

import (
	"strings"
)

// ShortestStandardizedPath returns the shortest standardized path.
func ShortestStandardizedPath(path string) string {
	parts := strings.Split(path, "/")
	var result []string
	for _, part := range parts {
		switch part {
		case "..":
			result = result[:len(result)-1]
		case ".":
		default:
			result = append(result, part)
		}
	}
	return strings.Join(result, "/")
}
