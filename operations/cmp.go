package operations

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

	}
	return 2, ""
}
