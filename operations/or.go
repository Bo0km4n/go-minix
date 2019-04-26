package operations

// OR model
type OR struct{}

// Analyze or analyze
func (or *OR) Analyze(ctx *Context, inst byte) (int, string) {
	switch inst {
	case 0x08: // Reg./Memory and Register to Either d = 0, w = 0
		opt := ctx.Body[ctx.Idx+1]
		mod := opt & maskTop2 >> 6
		reg := opt & maskMid3 >> 3
		rm := opt & maskLow3
		fromOrTo := false
		return getModRegRM(ctx, mod, rm, fromOrTo, Reg8b(reg), "or", Reg8b)
	case 0x09: // Reg./Memory and Register to Either d = 0, w = 1
		opt := ctx.Body[ctx.Idx+1]
		d := 0x00
		w := 0x01
		regFunc := getRegFunc(byte(w))

		mod := opt & maskTop2 >> 6
		reg := opt & maskMid3 >> 3
		rm := opt & maskLow3
		fromOrTo := d == 0x01
		regStr := regFunc(reg)

		return getModRegRM(ctx, mod, rm, fromOrTo, regStr, "or", regFunc)
	case 0x0a: // Reg./Memory and Register to Either d = 1, w = 0
		opt := ctx.Body[ctx.Idx+1]
		mod := opt & maskTop2 >> 6
		reg := opt & maskMid3 >> 3
		rm := opt & maskLow3
		fromOrTo := true
		return getModRegRM(ctx, mod, rm, fromOrTo, Reg8b(reg), "or", Reg8b)
	case 0x0b: // Reg./Memory and Register to Either d = 1, w = 1
		opt := ctx.Body[ctx.Idx+1]
		mod := opt & maskTop2 >> 6
		reg := opt & maskMid3 >> 3
		rm := opt & maskLow3
		fromOrTo := true
		return getModRegRM(ctx, mod, rm, fromOrTo, Reg16b(reg), "or", Reg16b)

	}
	return 0, ""
}
