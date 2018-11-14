package day85

// MathematicalIf returns x if b=1 and y if b=0.
func MathematicalIf(x, y, b int32) int32 {
	var xMask, yMask int32
	xMask = b
	xMask |= xMask << 1
	xMask |= xMask << 2
	xMask |= xMask << 4
	xMask |= xMask << 8
	xMask |= xMask << 16
	yMask = ^xMask
	return (x & xMask) | (y & yMask)
}
