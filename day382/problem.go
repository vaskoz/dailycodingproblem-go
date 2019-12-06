package day382

import (
	"encoding/base64"
	"encoding/hex"
)

// Base64DecodeDelegate just delegates
func Base64DecodeDelegate(b64 string) string {
	data, err := base64.StdEncoding.DecodeString(b64)
	if err != nil {
		panic("bad base64 input")
	}

	return hex.EncodeToString(data)
}
