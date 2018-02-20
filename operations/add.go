package operations

// ADD model
type ADD struct{}

// Analyze add analyze
func (add *ADD) Analyze(ctx *Context, inst byte) (int, string) {
	switch inst {
	case 0x00:
		opt := ctx.Body[ctx.Idx+1]
		ea := EA(opt & maskLow3)
		reg := Reg8b(opt & maskMid3)
		return 2, getResult(ctx.Idx, getOrgOpe(ctx.Body[ctx.Idx], ctx.Body[ctx.Idx+1]), getOpeString("add", ea, reg))
	case 0x01:
	case 0x02:
	case 0x03:
	case 0x04:
	case 0x05:
	default:
		return 0, ""
	}
	return 0, ""
}
