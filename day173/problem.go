package day173

import "strings"

// FlattenMap takes an arbitrary map and flattens it.
// It concatenates keys with '.'
func FlattenMap(m map[string]interface{}) map[string]interface{} {
	result := make(map[string]interface{})
	bfsFlattenMap([]string{}, m, result)
	return result
}

func bfsFlattenMap(prefix []string, m, result map[string]interface{}) {
	for k, v := range m {
		if obj, ok := v.(map[string]interface{}); ok {
			bfsFlattenMap(append(prefix, k), obj, result)
		} else {
			result[strings.Join(append(prefix, k), ".")] = v
		}
	}
}
