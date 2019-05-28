package operations

import "fmt"

// JBE model
type JBE struct{}

// Analyze JBE analyze
func (JBE *JBE) Analyze(ctx *Context, inst byte) (int, string) {
	switch inst {
	case 0x76:
		disp := signExtend(ctx.Body[ctx.Idx+1])
		retAddrStr := fmt.Sprintf("%04x", uint16(ctx.Idx+2)+disp)
		return 2, getResult(ctx.Idx, getOrgOpe(ctx.Body[ctx.Idx:ctx.Idx+2]), getOpeString("jbe", retAddrStr))
	}

	return NOT_FOUND, ""
}
