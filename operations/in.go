package operations

import (
	"fmt"
)

// IN model
type IN struct{}

// Analyze in analyze
func (in *IN) Analyze(ctx *Context, inst byte) (int, string) {
	switch inst {
	case 0xe5:
		opt := ctx.Body[ctx.Idx+1]
		regStr := "ax"
		port := fmt.Sprintf("%02x", opt)
		return 2, getResult(ctx.Idx, getOrgOpe(ctx.Body[ctx.Idx], ctx.Body[ctx.Idx+1]), getOpeString("in", regStr, port))
	case 0xec:
		return 1, getResult(ctx.Idx, getOrgOpe(ctx.Body[ctx.Idx]), getOpeString("in", "al", "ax"))
	default:
		return 0, ""
	}
}
