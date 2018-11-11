package day82

import (
	"io"
)

// Read7 reads only 7 bytes at a time from the given io.Reader.
func Read7(r io.Reader) string {
	lr := &io.LimitedReader{R: r, N: 7}
	b := make([]byte, 7)
	n, err := lr.Read(b)
	if err != nil {
		return ""
	}
	return string(b[:n])
}

// ReadN reads N-bytes using only the Read7 function that reads 7-bytes at a time.
func ReadN(r io.Reader, N int) string {
	result := make([]byte, N)
	readBytes := 0
	for readBytes < N {
		b := Read7(r)
		if len(b) == 0 {
			break
		}
		copy(result[readBytes:], []byte(b))
		readBytes += len(b)
	}
	return string(result[:readBytes])
}
