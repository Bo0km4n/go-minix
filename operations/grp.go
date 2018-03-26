package operations

import (
	"fmt"
)

// GRP model
type GRP struct{}

// Analyze grp analyze
func (grp *GRP) Analyze(ctx *Context, inst byte) (int, string) {
	switch inst {
	case 0x83:
		mode := (ctx.Body[ctx.Idx+1] & maskMid3) >> 3
		return grp.matchOpe(ctx, inst, mode)
	case 0x80:
		mode := (ctx.Body[ctx.Idx+1] & maskMid3) >> 3
		return grp.matchOpe(ctx, inst, mode)
	default:
		return 0, ""
	}
}

func (grp *GRP) matchOpe(ctx *Context, inst byte, mode byte) (int, string) {
	switch mode {
	case 0x00:
		regMode := inst & 0x01
		regFunc := getRegFunc(regMode)
		regAddr := ctx.Body[ctx.Idx+1] & maskLow3
		regStr := regFunc(regAddr)
		im := fmt.Sprintf("%02x", ctx.Body[ctx.Idx+2])
		return 3, getResult(ctx.Idx, getOrgOpe(ctx.Body[ctx.Idx:ctx.Idx+3]), getOpeString("add", regStr, im))
	case 0x07:
		s := inst & 0x02 >> 1
		w := inst & 0x01
		opt := ctx.Body[ctx.Idx+1]
		mod := opt & maskTop2 >> 6
		rm := opt & maskLow3
		data := ctx.Body[ctx.Idx+2]

		if s == 0x00 && w == 0x01 {
			addtionalData := ctx.Body[ctx.Idx+3]
			ea := getRM(mod, rm, int(data))
			dataStr := fmt.Sprintf("%d", signExtend(addtionalData))

			return 4, getResult(ctx.Idx, getOrgOpe(ctx.Body[ctx.Idx:ctx.Idx+4]), getOpeString("cmp", ea, dataStr))
		}
		ea := getRM(mod, rm, int(data))
		dataStr := fmt.Sprintf("%02x", data)

		return 3, getResult(ctx.Idx, getOrgOpe(ctx.Body[ctx.Idx:ctx.Idx+3]), getOpeString("cmp", ea, dataStr))

	default:
		return 0, ""
	}
}
