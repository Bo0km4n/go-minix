package operations

// ADD model
type ADD struct{}

// Analyze add analyze
func (add *ADD) Analyze(ctx *Context, inst byte) (int, string) {
	switch inst {
	case 0x00:
		opt := ctx.Body[ctx.Idx+1]
		d := inst & 0x02 >> 1
		w := inst & 0x01
		regFunc := getRegFunc(w)

		mod := opt & maskTop2 >> 6
		reg := opt & maskMid3 >> 3
		rm := opt & maskLow3

		fromOrTo := d == 0x01
		regStr := regFunc(reg)

		return getModRegRM(ctx, mod, rm, fromOrTo, regStr, "add", regFunc)

	case 0x03:
		opt := ctx.Body[ctx.Idx+1]
		d := inst & 0x02 >> 1
		w := inst & 0x01
		regFunc := getRegFunc(w)

		mod := opt & maskTop2 >> 6
		reg := opt & maskMid3 >> 3
		rm := opt & maskLow3

		fromOrTo := getFromOrTo(d)
		regStr := regFunc(reg)

		return getModRegRM(ctx, mod, rm, fromOrTo, regStr, "add", regFunc)
	}
	return 999, ""
}
