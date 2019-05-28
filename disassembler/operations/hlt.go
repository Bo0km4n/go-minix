package operations

type HLT struct{}

func (hlt *HLT) Analyze(ctx *Context, inst byte) (int, string) {
	return 1, getResult(ctx.Idx, getOrgOpe([]byte{ctx.Body[ctx.Idx]}), getOpeString("hlt"))
}
