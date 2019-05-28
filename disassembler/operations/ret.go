package operations

import "fmt"

// RET model
type RET struct{}

// Analyze ret analyze
func (ret *RET) Analyze(ctx *Context, inst byte) (int, string) {
	switch inst {
	case 0xc2: // Within Seg Adding Immed to SP
		dispStr := fmt.Sprintf("%4x", joinDispHighAndLow(ctx.Body[ctx.Idx+1], ctx.Body[ctx.Idx+2]))
		return 3, getResult(ctx.Idx, getOrgOpe(ctx.Body[ctx.Idx:ctx.Idx+3]), getOpeString("ret", dispStr))
	case 0xc3: // Within Segment
		return 1, getResult(ctx.Idx, getOrgOpe(ctx.Body[ctx.Idx:ctx.Idx+1]), getOpeString("ret"))
	}
	return NOT_FOUND, ""
}
