package operations

import "fmt"

// XOR model
type XOR struct{}

// Analyze xor analyze
func (xor *XOR) Analyze(ctx *Context, inst byte) {
	switch inst {
	case 0x30:
		// Reg./Memory and Register to Either 00001000
	case 0x31:
		// Reg./Memory and Register to Either 00001001
	case 0x32:
		// Reg./Memory and Register to Either 00001010
	case 0x33:
		// Reg./Memory and Register to Either 00001011
	case 0x80:
		// Immediate to Register/Memory 10000000
	case 0x81:
		// Immediate to Register/Memory 10000001
	case 0x34:
		// Immediate to Accumulator 00110100
	case 0x35:
		// Immediate to Accumulator 00110101
	}
}

// Next xor next
func (xor *XOR) Next() (int, string) {
	return 1, fmt.Sprintf("%04d: %-10x %s", 2, 0x89e3, "mov bx, sp")
}
