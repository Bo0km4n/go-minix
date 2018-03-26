package operations

// RET model
type RET struct{}

// Analyze ret analyze
func (ret *RET) Analyze(ctx *Context, inst byte) (int, string) {
	switch inst {
	case 0xc3:
		return 1, getResult(ctx.Idx, getOrgOpe(ctx.Body[ctx.Idx:ctx.Idx+1]), getOpeString("ret"))
	}
	return 999, ""
}
