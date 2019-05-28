package operations

type ADC struct{}

func (and *ADC) Analyze(ctx *Context, inst byte) (int, string) {
	switch inst {
	case 0x11: // Reg./Memory with Register to Either d = 0, w = 1
		opt := ctx.Body[ctx.Idx+1]
		mod := opt & maskTop2 >> 6
		rm := opt & maskLow3
		regFunc := Reg16b
		return getModRegRM(ctx, mod, rm, false, "", "adc", regFunc)

	}

	return NOT_FOUND, ""
}
