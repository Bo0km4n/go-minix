package operations

import "fmt"

// JNB model
type JNB struct{}

// Analyze JNB analyze
func (j *JNB) Analyze(ctx *Context, inst byte) (int, string) {
	switch inst {
	case 0x73:
		toAddr := byte(ctx.Idx+2) + ctx.Body[ctx.Idx+1]
		toAddrStr := fmt.Sprintf("%04x", toAddr)
		return 2, getResult(ctx.Idx, getOrgOpe(ctx.Body[ctx.Idx:ctx.Idx+2]), getOpeString("jnb", toAddrStr))
	default:
		return OVER_RANGE, ""
	}
}
