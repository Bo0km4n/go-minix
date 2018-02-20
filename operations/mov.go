package operations

import "fmt"

// MOV model
type MOV struct{}

// Analyze mov analyze
func (mov *MOV) Analyze(ctx *Context, inst byte) (int, string) {
	switch inst {
	case 0xb0:
	case 0xb1:
	case 0xb2:
	case 0xb3:
	case 0xb4:
	case 0xb5:
	case 0xb6:
	case 0xb7:
	case 0xb8:
	case 0xb9:
	case 0xba:
	case 0xbb:
		reg := inst & maskLow3
		rw := Reg16b(reg)
		iw := fmt.Sprintf("%02x%02x", ctx.Body[ctx.Idx+1], ctx.Body[ctx.Idx+2])
		return 3, getResult(ctx.Idx, getOrgOpe(inst, ctx.Body[ctx.Idx+1], ctx.Body[ctx.Idx+2]), getOpeString("mov", rw, iw))
	case 0xbc:
	case 0xbd:
	case 0xbe:
	case 0xbf:
	default:
		return 0, ""
	}
	return 0, ""
}
