package operations

type DEC struct{}

func (dec *DEC) Analyze(ctx *Context, inst byte) (int, string) {
	reg := inst & maskLow3
	regStr := getRegFunc(1)(reg)
	return 1, getResult(ctx.Idx, getOrgOpe([]byte{ctx.Body[ctx.Idx]}), getOpeString("dec", regStr))
}
