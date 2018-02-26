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

		switch mod {
		case 0x00:
			disp := 0
			ea := getRM(mod, rm, disp)

			if fromOrTo {
				return 2, getResult(ctx.Idx, getOrgOpe(ctx.Body[ctx.Idx], ctx.Body[ctx.Idx+1]), getOpeString("add", regStr, ea))
			}
			return 2, getResult(ctx.Idx, getOrgOpe(ctx.Body[ctx.Idx], ctx.Body[ctx.Idx+1]), getOpeString("add", ea, regStr))

		case 0x01:
			disp := signExtend(ctx.Body[ctx.Idx+2])
			ea := getRM(mod, rm, int(int16(disp)))

			if fromOrTo {
				return 3, getResult(ctx.Idx, getOrgOpe(ctx.Body[ctx.Idx], ctx.Body[ctx.Idx+1], ctx.Body[ctx.Idx+2]), getOpeString("add", regStr, ea))
			}
			return 3, getResult(ctx.Idx, getOrgOpe(ctx.Body[ctx.Idx], ctx.Body[ctx.Idx+1], ctx.Body[ctx.Idx+2]), getOpeString("add", ea, regStr))

		case 0x02:
			disp := joinDispHighAndLow(ctx.Body[ctx.Idx+2], ctx.Body[ctx.Idx+3])
			ea := getRM(mod, rm, disp)

			if fromOrTo {
				return 4, getResult(ctx.Idx, getOrgOpe(ctx.Body[ctx.Idx], ctx.Body[ctx.Idx+1], ctx.Body[ctx.Idx+2], ctx.Body[ctx.Idx+3]), getOpeString("add", regStr, ea))
			}
			return 4, getResult(ctx.Idx, getOrgOpe(ctx.Body[ctx.Idx], ctx.Body[ctx.Idx+1], ctx.Body[ctx.Idx+2], ctx.Body[ctx.Idx+3]), getOpeString("add", ea, regStr))

		case 0x03:
			rmReg := regFunc(rm)

			if fromOrTo {
				return 2, getResult(ctx.Idx, getOrgOpe(ctx.Body[ctx.Idx], ctx.Body[ctx.Idx+1]), getOpeString("add", regStr, rmReg))
			}
			return 2, getResult(ctx.Idx, getOrgOpe(ctx.Body[ctx.Idx], ctx.Body[ctx.Idx+1]), getOpeString("add", rmReg, regStr))
		}
	}

	return 0, ""
}
