package operations

import "fmt"

type TEST struct{}

// Analyze test analyze
func (test *TEST) Analyze(ctx *Context, inst byte) (int, string) {
	switch inst {
	case 0x85: // Reg./Memory and Register d = 0, w = 1
		opt := ctx.Body[ctx.Idx+1]
		mod := opt & maskTop2 >> 6
		reg := opt & maskMid3 >> 3
		rm := opt & maskLow3
		return getModRegRM(ctx, mod, rm, false, Reg16b(reg), "test", Reg16b)

	case 0xa8: // Immediate Data and Accumulator w = 0
		dataStr := fmt.Sprintf("%x", ctx.Body[ctx.Idx+1])
		regStr := Reg8b(0x00)
		return 2, getResult(ctx.Idx, getOrgOpe(ctx.Body[ctx.Idx:ctx.Idx+2]), getOpeString("test", regStr, dataStr))
	}
	return NOT_FOUND, ""
}
