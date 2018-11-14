package day85

// MathematicalIf returns x if b=1 and y if b=0.
func MathematicalIf(x, y, b int32) int32 {
	var xMask, yMask int32
	for i := uint32(0); i < 32; i++ {
		xMask |= b << i
	}
	yMask = ^xMask
	return (x & xMask) | (y & yMask)
}
