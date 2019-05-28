package operations

import (
	"fmt"
)

// CALL model
type CALL struct{}

// Analyze mov analyze
func (call *CALL) Analyze(ctx *Context, inst byte) (int, string) {
	switch inst {
	case 0xe8:
		// FIX
		offset := uint16(ctx.Body[ctx.Idx+2])<<8 + uint16(ctx.Body[ctx.Idx+1])
		retAddrStr := fmt.Sprintf("%04x", offset+uint16(ctx.Idx+3))
		return 3, getResult(ctx.Idx, getOrgOpe(ctx.Body[ctx.Idx:ctx.Idx+3]), getOpeString("call", retAddrStr))
	default:
		return 0, ""
	}
}
