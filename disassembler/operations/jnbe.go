package operations

import "fmt"

// JNBE model
type JNBE struct{}

// Analyze JNBE analyze
func (JNBE *JNBE) Analyze(ctx *Context, inst byte) (int, string) {
	switch inst {
	case 0x77:
		disp := signExtend(ctx.Body[ctx.Idx+1])
		retAddrStr := fmt.Sprintf("%04x", uint16(ctx.Idx+2)+disp)
		return 2, getResult(ctx.Idx, getOrgOpe(ctx.Body[ctx.Idx:ctx.Idx+2]), getOpeString("jnbe", retAddrStr))
	}

	return NOT_FOUND, ""
}
