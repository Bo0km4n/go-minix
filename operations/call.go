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
		retAddrTop8b := ctx.Body[ctx.Idx+2]
		retAddrBot8b := ctx.Body[ctx.Idx+1]
		retAddr16b := uint16(retAddrTop8b)
		retAddr16b = retAddr16b << 8
		retAddr16b += retAddr16b + uint16(retAddrBot8b)
		retAddrStr := fmt.Sprintf("%04x", int16(uint16(ctx.Idx+3)+retAddr16b))
		return 3, getResult(ctx.Idx, getOrgOpe(ctx.Body[ctx.Idx:ctx.Idx+3]), getOpeString("call", retAddrStr))
	default:
		return 0, ""
	}
}
