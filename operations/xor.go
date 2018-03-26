package operations

import "fmt"

// XOR model
type XOR struct{}

// Analyze xor analyze
func (xor *XOR) Analyze(ctx *Context, inst byte) (int, string) {
	switch inst {
	case 0x30:
		// Reg./Memory and Register to Either 00001000
	case 0x31:
		// Reg./Memory and Register to Either 00001001
		d := 0x00
		w := 0x01
		opt := ctx.Body[ctx.Idx+1]
		mod := opt & maskTop2 >> 6
		reg := opt & maskMid3 >> 3
		rm := opt & maskLow3
		regFunc := getRegFunc(byte(w))
		regStr := regFunc(reg)

		return getModRegRM(ctx, mod, rm, getFromOrTo(byte(d)), regStr, "xor", regFunc)
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
	return 999, ""
}

// Next xor next
func (xor *XOR) Next() (int, string) {
	return 1, fmt.Sprintf("%04d: %-10x %s", 2, 0x89e3, "mov bx, sp")
}
