package operations

// POP model
type POP struct{}

// Analyze pop analyze
func (pop *POP) Analyze(ctx *Context, inst byte) (int, string) {
	switch inst {
	case 0x58:
		reg := Reg16b(inst & maskLow3)
		return 1, getResult(ctx.Idx, getOrgOpe(ctx.Body[ctx.Idx:ctx.Idx+1]), getOpeString("pop", reg))
	case 0x59:
		reg := Reg16b(inst & maskLow3)
		return 1, getResult(ctx.Idx, getOrgOpe(ctx.Body[ctx.Idx:ctx.Idx+1]), getOpeString("pop", reg))
	case 0x5a:
		reg := Reg16b(inst & maskLow3)
		return 1, getResult(ctx.Idx, getOrgOpe(ctx.Body[ctx.Idx:ctx.Idx+1]), getOpeString("pop", reg))
	case 0x5b:
		reg := Reg16b(inst & maskLow3)
		return 1, getResult(ctx.Idx, getOrgOpe(ctx.Body[ctx.Idx:ctx.Idx+1]), getOpeString("pop", reg))
	case 0x5d:
		reg := Reg16b(inst & maskLow3)
		return 1, getResult(ctx.Idx, getOrgOpe(ctx.Body[ctx.Idx:ctx.Idx+1]), getOpeString("pop", reg))
	case 0x5f:
		reg := Reg16b(inst & maskLow3)
		return 1, getResult(ctx.Idx, getOrgOpe(ctx.Body[ctx.Idx:ctx.Idx+1]), getOpeString("pop", reg))
	case 0x5e:
		reg := Reg16b(inst & maskLow3)
		return 1, getResult(ctx.Idx, getOrgOpe(ctx.Body[ctx.Idx:ctx.Idx+1]), getOpeString("pop", reg))
	}
	return NOT_FOUND, ""
}
