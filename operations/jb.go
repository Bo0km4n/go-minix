package operations

import "fmt"

type JB struct{}

func (jb *JB) Analyze(ctx *Context, inst byte) (int, string) {
	disp := signExtend(ctx.Body[ctx.Idx+1])
	retAddrStr := fmt.Sprintf("%04x", uint16(ctx.Idx+2)+disp)
	return 2, getResult(ctx.Idx, getOrgOpe(ctx.Body[ctx.Idx:ctx.Idx+2]), getOpeString("jb", retAddrStr))
}
