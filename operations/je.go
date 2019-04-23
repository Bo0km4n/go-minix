package operations

import "fmt"

// JE model
type JE struct{}

// Analyze je analyze
func (je *JE) Analyze(ctx *Context, inst byte) (int, string) {
	switch inst {
	case 0x74:
		disp := uint16(ctx.Body[ctx.Idx+1])
		retAddrStr := fmt.Sprintf("%04x", uint16(ctx.Idx+2)+disp)
		return 2, getResult(ctx.Idx, getOrgOpe(ctx.Body[ctx.Idx:ctx.Idx+2]), getOpeString("je", retAddrStr))
	}

	return OVER_RANGE, ""
}
