package operations

// OR model
type OR struct{}

// Analyze or analyze
func (or *OR) Analyze(ctx *Context, inst byte) (int, string) {
	switch inst {
	case 0x09:
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
	}
	return 0, ""
}
