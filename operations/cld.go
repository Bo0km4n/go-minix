package operations

type CLD struct{}

// Analyze cld analyze
func (cwd *CLD) Analyze(ctx *Context, inst byte) (int, string) {
	return 1, getResult(ctx.Idx, getOrgOpe([]byte{ctx.Body[ctx.Idx]}), getOpeString("cld"))
}
