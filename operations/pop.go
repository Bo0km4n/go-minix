package operations

// POP model
type POP struct{}

// Analyze pop analyze
func (pop *POP) Analyze(ctx *Context, inst byte) (int, string) {
	switch inst {
	case 0x5b:
		reg := Reg16b(inst & maskLow3)
		return 1, getResult(ctx.Idx, getOrgOpe(inst), getOpeString("pop", reg))
	}
	return 0, ""
}
