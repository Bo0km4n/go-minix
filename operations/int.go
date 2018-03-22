package operations

import "fmt"

// INT operation
type INT struct{}

// Analyze implementation analyze for INT
func (t *INT) Analyze(ctx *Context, inst byte) (int, string) {
	switch inst {
	case 0xcc:
		// int3
	case 0xcd:
		// int ib
		ib := fmt.Sprintf("%02x", ctx.Body[ctx.Idx+1])
		return 2, getResult(ctx.Idx, getOrgOpe(ctx.Body[ctx.Idx:ctx.Idx+2]), getOpeString("int", ib))
	default:
		return 0, ""
	}
	return 0, ""
}
