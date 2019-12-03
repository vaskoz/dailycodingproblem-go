package day378

import "encoding/json"

// ToJSON converts arbitrary data into a json string.
// Delegates to json.Marshal because it exists.
func ToJSON(data interface{}) string {
	b, err := json.Marshal(data)
	if err != nil {
		return ""
	}

	return string(b)
}
