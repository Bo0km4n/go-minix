package operations

type CWD struct{}

// Analyze cwd analyze
func (cwd *CWD) Analyze(ctx *Context, inst byte) (int, string) {
	return 1, getResult(ctx.Idx, getOrgOpe([]byte{ctx.Body[ctx.Idx]}), getOpeString("cwd"))
}
