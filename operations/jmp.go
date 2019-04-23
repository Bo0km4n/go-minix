package operations

import (
	"fmt"
)

// JMP model
type JMP struct{}

// Analyze jmp analyze
func (jmp *JMP) Analyze(ctx *Context, inst byte) (int, string) {
	switch inst {
	case 0xe9:
		retAddrTop8b := ctx.Body[ctx.Idx+2]
		retAddrBot8b := ctx.Body[ctx.Idx+1]
		retAddr16b := uint16(retAddrTop8b)
		retAddr16b = retAddr16b << 7
		retAddr16b += retAddr16b + uint16(retAddrBot8b)
		retAddrStr := fmt.Sprintf("%04x", uint16(ctx.Idx+3)+retAddr16b)
		return 3, getResult(ctx.Idx, getOrgOpe(ctx.Body[ctx.Idx:ctx.Idx+3]), getOpeString("jmp", retAddrStr))
	case 0xeb:
		retAddrStr := fmt.Sprintf("short %04x", uint16(ctx.Idx+2)+uint16(ctx.Body[ctx.Idx+1]))
		// retAddrStr := fmt.Sprintf("short %04x", int(ctx.Body[ctx.Idx+1])-int(ctx.Idx+2))
		return 2, getResult(ctx.Idx, getOrgOpe(ctx.Body[ctx.Idx:ctx.Idx+2]), getOpeString("jmp", retAddrStr))
	default:
		return 100, ""
	}
}
