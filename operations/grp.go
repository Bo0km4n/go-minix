package operations

import (
	"fmt"

	"github.com/k0kubun/pp"
)

// GRP model
type GRP struct{}

// Analyze grp analyze
func (grp *GRP) Analyze(ctx *Context, inst byte) (int, string) {
	switch inst {
	case 0x83:
		mode := (ctx.Body[ctx.Idx+1] & maskMid3) >> 3
		return grp.matchOpe1WB(ctx, inst, mode)
	case 0x80:
		mode := (ctx.Body[ctx.Idx+1] & maskMid3) >> 3
		return grp.matchOpe1B(ctx, inst, mode)
	case 0x81:
		mode := (ctx.Body[ctx.Idx+1] & maskMid3) >> 3
		return grp.matchOpe1W(ctx, inst, mode)
	case 0xf7:
		mode := (ctx.Body[ctx.Idx+1] & maskMid3) >> 3
		return grp.matchOpe3W(ctx, inst, mode)
	case 0xf6:
		mode := (ctx.Body[ctx.Idx+1] & maskMid3) >> 3
		return grp.matchOpe3B(ctx, inst, mode)
	case 0xd1:
		mode := (ctx.Body[ctx.Idx+1] & maskMid3) >> 3
		return grp.matchOpe2(ctx, inst, mode)
	case 0xff:
		mode := (ctx.Body[ctx.Idx+1] & maskMid3) >> 3
		return grp.matchOpe5(ctx, inst, mode)
	default:
		return NOT_FOUND, ""
	}
}

func (grp *GRP) matchOpe1B(ctx *Context, inst byte, mode byte) (int, string) {
	switch mode {
	case 0x00:
		regMode := inst & 0x01
		regFunc := getRegFunc(regMode)
		regAddr := ctx.Body[ctx.Idx+1] & maskLow3
		regStr := regFunc(regAddr)
		im := fmt.Sprintf("%02x", ctx.Body[ctx.Idx+2])
		return 3, getResult(ctx.Idx, getOrgOpe(ctx.Body[ctx.Idx:ctx.Idx+3]), getOpeString("add", regStr, im))
	case 0x05:
		s := inst & 0x02 >> 1
		w := inst & 0x01
		opt := ctx.Body[ctx.Idx+1]
		mod := opt & maskTop2 >> 6
		rm := opt & maskLow3
		data := ctx.Body[ctx.Idx+2]

		if s == 0x00 || w == 0x01 {
			addtionalData := ctx.Body[ctx.Idx+3]
			ea := ""
			if mod == 0x01 {
				disp := signExtend(data)
				ea = getRM(mod, rm, int(int16(disp)))
			} else {
				ea = getRM(mod, rm, int(data))
				pp.Println(ea)
			}
			dataStr := fmt.Sprintf("%d", signExtend(addtionalData))

			return 4, getResult(ctx.Idx, getOrgOpe(ctx.Body[ctx.Idx:ctx.Idx+4]), getOpeString("sub", ea, dataStr))
		}
		ea := getRM(mod, rm, int(data))
		dataStr := fmt.Sprintf("%02x", data)

		return 3, getResult(ctx.Idx, getOrgOpe(ctx.Body[ctx.Idx:ctx.Idx+3]), getOpeString("cmp", ea, dataStr))
	case 0x07:
		s := inst & 0x02 >> 1
		w := inst & 0x01
		opt := ctx.Body[ctx.Idx+1]
		mod := opt & maskTop2 >> 6
		rm := opt & maskLow3
		data := ctx.Body[ctx.Idx+2]

		if s == 0x00 && w == 0x01 {
			addtionalData := ctx.Body[ctx.Idx+3]
			ea := ""
			if mod == 0x01 {
				disp := signExtend(data)
				ea = getRM(mod, rm, int(int16(disp)))
			} else {
				ea = getRM(mod, rm, int(data))
			}
			dataStr := fmt.Sprintf("%d", signExtend(addtionalData))

			return 4, getResult(ctx.Idx, getOrgOpe(ctx.Body[ctx.Idx:ctx.Idx+4]), getOpeString("cmp", ea, dataStr))
		}
		ea := getRM(mod, rm, int(data))
		dataStr := fmt.Sprintf("%02x", data)

		return 3, getResult(ctx.Idx, getOrgOpe(ctx.Body[ctx.Idx:ctx.Idx+3]), getOpeString("cmp", ea, dataStr))

	default:
		return NOT_FOUND, ""
	}
}

func (grp *GRP) matchOpe1WB(ctx *Context, inst byte, mode byte) (int, string) {
	switch mode {
	case 0x00:
		regMode := inst & 0x01
		regFunc := getRegFunc(regMode)
		regAddr := ctx.Body[ctx.Idx+1] & maskLow3
		regStr := regFunc(regAddr)
		im := fmt.Sprintf("%02x", ctx.Body[ctx.Idx+2])
		return 3, getResult(ctx.Idx, getOrgOpe(ctx.Body[ctx.Idx:ctx.Idx+3]), getOpeString("add", regStr, im))
	case 0x05:
		s := inst & 0x02 >> 1
		w := inst & 0x01
		opt := ctx.Body[ctx.Idx+1]
		mod := opt & maskTop2 >> 6
		rm := opt & maskLow3
		data := ctx.Body[ctx.Idx+2]

		if (s == 0x01 || w == 0x01) && !(s == 0x01 && w == 0x01) {
			addtionalData := ctx.Body[ctx.Idx+3]
			ea := ""
			if mod == 0x01 {
				disp := signExtend(data)
				ea = getRM(mod, rm, int(int16(disp)))
			} else {
				ea = getRM(mod, rm, int(data))
			}
			dataStr := fmt.Sprintf("%d", signExtend(addtionalData))

			return 4, getResult(ctx.Idx, getOrgOpe(ctx.Body[ctx.Idx:ctx.Idx+4]), getOpeString("sub", ea, dataStr))
		}
		ea := getRM(mod, rm, int(data))
		dataStr := fmt.Sprintf("%02x", data)

		return 3, getResult(ctx.Idx, getOrgOpe(ctx.Body[ctx.Idx:ctx.Idx+3]), getOpeString("sub", ea, dataStr))
	case 0x07: // CMP Immediate with Register/Memory s = 1, w = 1
		opt := ctx.Body[ctx.Idx+1]
		mod := opt & maskTop2 >> 6
		rm := opt & maskLow3
		disp := ctx.Body[ctx.Idx+2]
		data := ctx.Body[ctx.Idx+3]

		ea := getRM(mod, rm, int(int16(signExtend(disp))))
		dataExtended := fmt.Sprintf("%x", signExtend(data))
		return 4, getResult(ctx.Idx, getOrgOpe(ctx.Body[ctx.Idx:ctx.Idx+4]), getOpeString("cmp", ea, dataExtended))

	default:
		return NOT_FOUND, ""
	}
}

func (grp *GRP) matchOpe3W(ctx *Context, inst, mode byte) (int, string) {
	switch mode {
	case 0x03: // NEG: Change sign w = 1
		opt := ctx.Body[ctx.Idx+1]
		w := inst & 0x01
		mod := opt & maskTop2 >> 6
		rm := opt & maskLow3
		regFunc := getRegFunc(w)
		return getModRegRM(ctx, mod, rm, false, "", "neg", regFunc)
	case 0x04:
		opt := ctx.Body[ctx.Idx+1]
		mod := opt & maskTop2 >> 6
		rm := opt & maskLow3
		ea := getRM(mod, rm, 0)

		return 2, getResult(ctx.Idx, getOrgOpe(ctx.Body[ctx.Idx:ctx.Idx+2]), getOpeString("mul", ea))
	}
	return NOT_FOUND, ""
}

func (grp *GRP) matchOpe1W(ctx *Context, inst, mode byte) (int, string) {
	switch mode {
	case 0x07:
		opt := ctx.Body[ctx.Idx+1]
		mod := opt & maskTop2 >> 6
		rm := opt & maskLow3
		dataHigh8bit := ctx.Body[ctx.Idx+2]
		dataLow8bit := ctx.Body[ctx.Idx+3]
		dataStr := fmt.Sprintf("%02x%02x", dataLow8bit, dataHigh8bit)
		return 4, getResult(ctx.Idx, getOrgOpe(ctx.Body[ctx.Idx:ctx.Idx+4]), getOpeString("cmp", getRM(mod, rm, 0), dataStr))
	default:
		return NOT_FOUND, ""
	}
}

func (grp *GRP) matchOpe3B(ctx *Context, inst, mode byte) (int, string) {
	switch mode {
	case 0x00: // test reg data
		opt := ctx.Body[ctx.Idx+1]
		rm := opt & maskLow3
		w := inst & 0x01
		regStr := getRegFunc(w)(rm)
		dataStr := fmt.Sprintf("%02x", ctx.Body[ctx.Idx+2])
		return 3, getResult(ctx.Idx, getOrgOpe(ctx.Body[ctx.Idx:ctx.Idx+3]), getOpeString("test", regStr, dataStr))
	default:
		return NOT_FOUND, ""
	}
}

func (grp *GRP) matchOpe2(ctx *Context, inst, mode byte) (int, string) {
	switch mode {
	case 0x04: // shl reg data
		var countStr string
		opt := ctx.Body[ctx.Idx+1]
		// mod := opt & maskTop2 >> 6
		rm := opt & maskLow3
		w := inst & 0x01
		v := inst & 0x02
		if v == 0x00 {
			countStr = fmt.Sprintf("%d", 1)
		} else {
			countStr = "cl"
		}
		regStr := getRegFunc(w)(rm)
		return 2, getResult(ctx.Idx, getOrgOpe(ctx.Body[ctx.Idx:ctx.Idx+2]), getOpeString("shl", regStr, countStr))
	}
	return NOT_FOUND, ""
}

func (grp *GRP) matchOpe5(ctx *Context, inst, mode byte) (int, string) {
	switch mode {
	case 0x06:
		opt := ctx.Body[ctx.Idx+1]
		disp := ctx.Body[ctx.Idx+2]
		mod := opt & maskTop2 >> 6
		rm := opt & maskLow3
		ea := getRM(mod, rm, int(disp))
		return 3, getResult(ctx.Idx, getOrgOpe(ctx.Body[ctx.Idx:ctx.Idx+3]), getOpeString("push", ea))
	case 0x02:
		opt := ctx.Body[ctx.Idx+1]
		mod := opt & maskTop2 >> 6
		rm := opt & maskLow3
		ea := getRM(mod, rm, 0)
		return 2, getResult(ctx.Idx, getOrgOpe(ctx.Body[ctx.Idx:ctx.Idx+2]), getOpeString("call", ea))
	}

	return NOT_FOUND, ""
}
