package operations

// SBB model
type SBB struct{}

// Analyze sbb analyze
func (sbb *SBB) Analyze(ctx *Context, inst byte) (int, string) {
	switch inst {
	case 0x18:
		opt := ctx.Body[ctx.Idx+1]
		d := 0x00
		w := 0x00
		regFunc := getRegFunc(byte(w))
		mod := opt & maskTop2 >> 6
		reg := opt & maskMid3 >> 3
		rm := opt & maskLow3

		fromOrTo := d == 0x01
		regStr := regFunc(reg)

		return getModRegRM(ctx, mod, rm, fromOrTo, regStr, "sbb", regFunc)
	case 0x19: // d = 0, w = 1
		opt := ctx.Body[ctx.Idx+1]
		mod := opt & maskTop2 >> 6
		reg := opt & maskMid3 >> 3
		rm := opt & maskLow3

		return getModRegRM(ctx, mod, rm, false, Reg16b(reg), "sbb", Reg16b)

	}
	return NOT_FOUND, ""
}
