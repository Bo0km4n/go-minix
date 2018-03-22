package operations

// LEA model
type LEA struct{}

// Analyze lea analyze
func (lea *LEA) Analyze(ctx *Context, inst byte) (int, string) {
	switch inst {
	case 0x8d:
		opt := ctx.Body[ctx.Idx+1]
		d := 0x01
		w := 0x01
		regFunc := getRegFunc(byte(w))
		mod := opt & maskTop2 >> 6
		reg := opt & maskMid3 >> 3
		rm := opt & maskLow3

		fromOrTo := d == 0x01
		regStr := regFunc(reg)

		return getModRegRM(ctx, mod, rm, fromOrTo, regStr, "lea", regFunc)
	}
	return 0, ""
}
