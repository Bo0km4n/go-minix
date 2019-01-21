package operations

// NEG model
type NEG struct{}

// Analyze neg analyze
func (neg *NEG) Analyze(ctx *Context, inst byte) (int, string) {
	switch inst {
	case 0xf7:
		opt := ctx.Body[ctx.Idx+1]
		w := inst & 0x01
		mod := opt & maskTop2 >> 6
		rm := opt & maskLow3
		regFunc := getRegFunc(w)
		return getModRegRM(ctx, mod, rm, false, "", "neg", regFunc)
	}

	return OVER_RANGE, ""
}
