package operations

type XCHG struct{}

func (xchg *XCHG) Analyze(ctx *Context, inst byte) (int, string) {
	switch inst {
	case 0x86: // Register/Memory with Register w = 1
		opt := ctx.Body[ctx.Idx+1]
		mod := opt & maskTop2 >> 6
		reg := opt & maskMid3 >> 3
		rm := opt & maskLow3
		return getModRegRM(ctx, mod, rm, false, Reg8b(reg), "xchg", Reg8b)
	case 0x87: // Register/Memory with Register w = 1
		opt := ctx.Body[ctx.Idx+1]
		mod := opt & maskTop2 >> 6
		reg := opt & maskMid3 >> 3
		rm := opt & maskLow3
		return getModRegRM(ctx, mod, rm, false, Reg16b(reg), "xchg", Reg16b)

	case 0x90: // Register with Accumulator
		regStr := Reg16b(inst & maskLow3)
		return 1, getResult(ctx.Idx, getOrgOpe(ctx.Body[ctx.Idx:ctx.Idx+1]), getOpeString("xchg", regStr, "ax"))
	case 0x91: // Register with Accumulator
		regStr := Reg16b(inst & maskLow3)
		return 1, getResult(ctx.Idx, getOrgOpe(ctx.Body[ctx.Idx:ctx.Idx+1]), getOpeString("xchg", regStr, "ax"))
	case 0x92: // Register with Accumulator
		regStr := Reg16b(inst & maskLow3)
		return 1, getResult(ctx.Idx, getOrgOpe(ctx.Body[ctx.Idx:ctx.Idx+1]), getOpeString("xchg", regStr, "ax"))
	case 0x93: // Register with Accumulator
		regStr := Reg16b(inst & maskLow3)
		return 1, getResult(ctx.Idx, getOrgOpe(ctx.Body[ctx.Idx:ctx.Idx+1]), getOpeString("xchg", regStr, "ax"))
	case 0x94: // Register with Accumulator
		regStr := Reg16b(inst & maskLow3)
		return 1, getResult(ctx.Idx, getOrgOpe(ctx.Body[ctx.Idx:ctx.Idx+1]), getOpeString("xchg", regStr, "ax"))
	case 0x95: // Register with Accumulator
		regStr := Reg16b(inst & maskLow3)
		return 1, getResult(ctx.Idx, getOrgOpe(ctx.Body[ctx.Idx:ctx.Idx+1]), getOpeString("xchg", regStr, "ax"))
	case 0x96: // Register with Accumulator
		regStr := Reg16b(inst & maskLow3)
		return 1, getResult(ctx.Idx, getOrgOpe(ctx.Body[ctx.Idx:ctx.Idx+1]), getOpeString("xchg", regStr, "ax"))
	case 0x97: // Register with Accumulator
		regStr := Reg16b(inst & maskLow3)
		return 1, getResult(ctx.Idx, getOrgOpe(ctx.Body[ctx.Idx:ctx.Idx+1]), getOpeString("xchg", regStr, "ax"))
	}
	return NOT_FOUND, ""
}
