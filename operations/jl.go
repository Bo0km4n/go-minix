package operations

import "fmt"

// JL model
type JL struct{}

// Analyze jl analyze
func (jl *JL) Analyze(ctx *Context, inst byte) (int, string) {
	retAddrStr := fmt.Sprintf("%04x", byte(ctx.Idx+2)+ctx.Body[ctx.Idx+1])
	return 2, getResult(ctx.Idx, getOrgOpe(ctx.Body[ctx.Idx:ctx.Idx+2]), getOpeString("jl", retAddrStr))
}
