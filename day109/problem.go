package day109

// SwapEvenOddBits swaps the even and odd bits
// of a uint8.
func SwapEvenOddBits(ui uint8) uint8 {
	return (ui&0xaa)>>1 | (ui&0x55)<<1
}
