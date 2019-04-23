package operations

type AND struct{}

func (and *AND) Analyze(ctx *Context, inst byte) (int, string) {
	opt := ctx.Body[ctx.Idx+1]
	switch inst {
	case 0x20:
		w := byte(0x00)
		mod := opt & maskTop2 >> 6
		reg := opt & maskMid3 >> 3
		rm := opt & maskLow3
		ea := getRM(mod, rm, 0)
		regStr := getRegFunc(w)(reg)
		return 2, getResult(ctx.Idx, getOrgOpe(ctx.Body[ctx.Idx:ctx.Idx+2]), getOpeString("and", ea, regStr))
	case 0x21:
		w := byte(0x01)
		mod := opt & maskTop2 >> 6
		reg := opt & maskMid3 >> 3
		rm := opt & maskLow3
		ea := getRM(mod, rm, 0)
		regStr := getRegFunc(w)(reg)
		return 2, getResult(ctx.Idx, getOrgOpe(ctx.Body[ctx.Idx:ctx.Idx+2]), getOpeString("and", ea, regStr))
	case 0x22:
		w := byte(0x00)
		mod := opt & maskTop2 >> 6
		reg := opt & maskMid3 >> 3
		rm := opt & maskLow3
		ea := getRM(mod, rm, 0)
		regStr := getRegFunc(w)(reg)
		return 2, getResult(ctx.Idx, getOrgOpe(ctx.Body[ctx.Idx:ctx.Idx+2]), getOpeString("and", regStr, ea))
	case 0x23:
		w := byte(0x01)
		mod := opt & maskTop2 >> 6
		reg := opt & maskMid3 >> 3
		rm := opt & maskLow3
		ea := getRM(mod, rm, 0)
		regStr := getRegFunc(w)(reg)
		return 2, getResult(ctx.Idx, getOrgOpe(ctx.Body[ctx.Idx:ctx.Idx+2]), getOpeString("and", regStr, ea))
	case 0x80:
		// w := 0x00 TODO
	case 0x81:
		// w := 0x01 TODO

	}

	return OVER_RANGE, ""
}
