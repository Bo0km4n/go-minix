package operations

import "fmt"

// LOOP model
type LOOP struct{}

// Analyze lea analyze
func (loop *LOOP) Analyze(ctx *Context, inst byte) (int, string) {
	switch inst {
	case 0xe2:
		retAddrStr := fmt.Sprintf("%04x", uint16(ctx.Idx+2)+signExtend(ctx.Body[ctx.Idx+1]))
		return 2, getResult(ctx.Idx, getOrgOpe(ctx.Body[ctx.Idx:ctx.Idx+2]), getOpeString("loop", retAddrStr))
	}
	return NOT_FOUND, ""
}
