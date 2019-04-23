package operations

// INC model
type INC struct{}

// Analyze inc analyze
func (inc *INC) Analyze(ctx *Context, inst byte) (int, string) {
	switch inst {
	case 0x40:
		reg := inst & maskLow3
		regStr := Reg16b(reg)
		return 1, getResult(ctx.Idx, getOrgOpe(ctx.Body[ctx.Idx:ctx.Idx+1]), getOpeString("inc", regStr))
	case 0x41:
		reg := inst & maskLow3
		regStr := Reg16b(reg)
		return 1, getResult(ctx.Idx, getOrgOpe(ctx.Body[ctx.Idx:ctx.Idx+1]), getOpeString("inc", regStr))
	case 0x42:
		reg := inst & maskLow3
		regStr := Reg16b(reg)
		return 1, getResult(ctx.Idx, getOrgOpe(ctx.Body[ctx.Idx:ctx.Idx+1]), getOpeString("inc", regStr))
	case 0x43:
		reg := inst & maskLow3
		regStr := Reg16b(reg)
		return 1, getResult(ctx.Idx, getOrgOpe(ctx.Body[ctx.Idx:ctx.Idx+1]), getOpeString("inc", regStr))
	case 0x44:
		reg := inst & maskLow3
		regStr := Reg16b(reg)
		return 1, getResult(ctx.Idx, getOrgOpe(ctx.Body[ctx.Idx:ctx.Idx+1]), getOpeString("inc", regStr))
	case 0x45:
		reg := inst & maskLow3
		regStr := Reg16b(reg)
		return 1, getResult(ctx.Idx, getOrgOpe(ctx.Body[ctx.Idx:ctx.Idx+1]), getOpeString("inc", regStr))
	case 0x46:
		reg := inst & maskLow3
		regStr := Reg16b(reg)
		return 1, getResult(ctx.Idx, getOrgOpe(ctx.Body[ctx.Idx:ctx.Idx+1]), getOpeString("inc", regStr))
	case 0x47:
		reg := inst & maskLow3
		regStr := Reg16b(reg)
		return 1, getResult(ctx.Idx, getOrgOpe(ctx.Body[ctx.Idx:ctx.Idx+1]), getOpeString("inc", regStr))
	}
	return OVER_RANGE, ""
}
