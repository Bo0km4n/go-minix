package operations

type AND struct{}

func (and *AND) Analyze(ctx *Context, inst byte) (int, string) {
	opt := ctx.Body[ctx.Idx+1]
	switch inst {
	case 0x20: // Reg./Memory and Register to Either d = 0, w = 0
		w := byte(0x00)
		mod := opt & maskTop2 >> 6
		reg := opt & maskMid3 >> 3
		rm := opt & maskLow3
		regStr := getRegFunc(w)(reg)
		return getModRegRM(ctx, mod, rm, getFromOrTo(0x00), regStr, "and", getRegFunc(w))
	case 0x21: // Reg./Memory and Register to Either d = 0, w = 1
		w := byte(0x01)
		mod := opt & maskTop2 >> 6
		reg := opt & maskMid3 >> 3
		rm := opt & maskLow3
		regStr := getRegFunc(w)(reg)
		return getModRegRM(ctx, mod, rm, getFromOrTo(0x00), regStr, "and", getRegFunc(w))
	case 0x22: // Reg./Memory and Register to Either d = 1, w = 0
		w := byte(0x00)
		mod := opt & maskTop2 >> 6
		reg := opt & maskMid3 >> 3
		rm := opt & maskLow3
		regStr := getRegFunc(w)(reg)
		return getModRegRM(ctx, mod, rm, getFromOrTo(0x01), regStr, "and", getRegFunc(w))
	case 0x23: // Reg./Memory and Register to Either d = 1, w = 1
		w := byte(0x01)
		mod := opt & maskTop2 >> 6
		reg := opt & maskMid3 >> 3
		rm := opt & maskLow3
		regStr := getRegFunc(w)(reg)
		return getModRegRM(ctx, mod, rm, getFromOrTo(0x01), regStr, "and", getRegFunc(w))
	case 0x80:
		// w := 0x00 TODO
	case 0x81:
		// w := 0x01 TODO

	}

	return OVER_RANGE, ""
}
