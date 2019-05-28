package operations

import "fmt"

// JLE model
type JLE struct{}

// Analyze JLE analyze
func (JLE *JLE) Analyze(ctx *Context, inst byte) (int, string) {
	switch inst {
	case 0x7e:
		disp := signExtend(ctx.Body[ctx.Idx+1])
		retAddrStr := fmt.Sprintf("%04x", uint16(ctx.Idx+2)+disp)
		return 2, getResult(ctx.Idx, getOrgOpe(ctx.Body[ctx.Idx:ctx.Idx+2]), getOpeString("jle", retAddrStr))
	}

	return NOT_FOUND, ""
}
