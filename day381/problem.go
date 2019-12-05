package day381

import (
	"encoding/base64"
	"encoding/hex"
)

// Base64Delegate just delegates
func Base64Delegate(hexs string) string {
	data, err := hex.DecodeString(hexs)
	if err != nil {
		panic("bad hex input")
	}

	return base64.StdEncoding.EncodeToString(data)
}
