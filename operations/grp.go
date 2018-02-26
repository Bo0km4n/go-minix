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
		return 3, getResult(ctx.Idx, getOrgOpe(inst, ctx.Body[ctx.Idx+1], ctx.Body[ctx.Idx+2]), getOpeString("add", regStr, im))
	default:
		return 0, ""
	}
}
