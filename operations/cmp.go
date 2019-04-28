package operations

import "fmt"

// CMP model
type CMP struct{}

// Analyze je analyze
func (cmp *CMP) Analyze(ctx *Context, inst byte) (int, string) {
	switch inst {
	case 0x38: // Register/Memory and Register d = 0, w = 0
		opt := ctx.Body[ctx.Idx+1]
		mod := opt & maskTop2 >> 6
		reg := opt & maskMid3 >> 3
		rm := opt & maskLow3
		return getModRegRM(ctx, mod, rm, false, Reg8b(reg), "cmp", Reg8b)

	case 0x39: // Register/Memory and Register d = 0, w = 1
		opt := ctx.Body[ctx.Idx+1]
		mod := opt & maskTop2 >> 6
		reg := opt & maskMid3 >> 3
		rm := opt & maskLow3
		return getModRegRM(ctx, mod, rm, false, Reg16b(reg), "cmp", Reg16b)

	case 0x3a: // Register/Memory and Register d = 1, w = 0
		opt := ctx.Body[ctx.Idx+1]
		mod := opt & maskTop2 >> 6
		reg := opt & maskMid3 >> 3
		rm := opt & maskLow3
		return getModRegRM(ctx, mod, rm, true, Reg8b(reg), "cmp", Reg8b)

	case 0x3b: // Register/Memory and Register d = 1, w = 1
		opt := ctx.Body[ctx.Idx+1]
		mod := opt & maskTop2 >> 6
		reg := opt & maskMid3 >> 3
		rm := opt & maskLow3
		return getModRegRM(ctx, mod, rm, true, Reg16b(reg), "cmp", Reg16b)

	case 0x3d: // Immediate with Accumulator w = 1
		opt := ctx.Body[ctx.Idx+1]
		mod := opt & maskTop2 >> 6
		// rm := opt & maskLow3

		switch mod {
		case 0x00:
			regStr := Reg16b(0x00)
			dataStr := fmt.Sprintf("%04x", joinDispHighAndLow(ctx.Body[ctx.Idx+1], ctx.Body[ctx.Idx+2]))
			return 3, getResult(ctx.Idx, getOrgOpe(ctx.Body[ctx.Idx:ctx.Idx+4]), getOpeString("cmp", regStr, dataStr))
		}
	}
	return 2, ""
}
