package operations

// PUSH model
type PUSH struct{}

// Analyze PUSH analyze
func (push *PUSH) Analyze(ctx *Context, inst byte) (int, string) {
	switch inst {
	case 0x06:
	case 0x16:
	case 0x50:
		reg := Reg16b(inst & maskLow3)
		return 1, getResult(ctx.Idx, getOrgOpe(inst), getOpeString("push", reg))
	case 0x51:
	case 0x52:
	case 0x53:
	case 0x54:
	case 0x55:
		reg := Reg16b(inst & maskLow3)
		return 1, getResult(ctx.Idx, getOrgOpe(inst), getOpeString("push", reg))
	case 0x56:
	case 0x57:
	default:
		return 0, ""
	}
	return 0, ""
}
