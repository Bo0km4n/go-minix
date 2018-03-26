package operations

// CBW model
type CBW struct{}

// Analyze cbw analyze
func (cbw *CBW) Analyze(ctx *Context, inst byte) (int, string) {
	return 1, getResult(ctx.Idx, getOrgOpe(ctx.Body[ctx.Idx:ctx.Idx+1]), getOpeString("cbw"))
}
