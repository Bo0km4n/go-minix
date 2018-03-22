package operations

import "fmt"

// JNL model
type JNL struct{}

// Analyze jnl analyze
func (jnl *JNL) Analyze(ctx *Context, inst byte) (int, string) {
	disp := uint16(ctx.Body[ctx.Idx+1])
	retAddrStr := fmt.Sprintf("%04x", uint16(ctx.Idx+2)+disp)
	return 2, getResult(ctx.Idx, getOrgOpe(inst, byte(disp)), getOpeString("jnl", retAddrStr))
}
