package operations

type REP struct{}

func (rep *REP) Analyze(ctx *Context, inst byte) (int, string) {
	switch inst {
	case 0xf2:
		return 2, getResult(ctx.Idx, getOrgOpe(ctx.Body[ctx.Idx:ctx.Idx+2]), getOpeString("rep", stringManipulations(ctx.Body[ctx.Idx+1])))
	}
	return NOT_FOUND, ""
}

func stringManipulations(opt byte) string {
	switch opt {
	case 0xa4:
		return "movsb"
	case 0xa5:
		return "movsw"
	case 0xa6:
		return "cmpsb"
	case 0xa7:
		return "cmpsw"
	case 0xae:
		return "scasb"
	case 0xaf:
		return "scasw"
	case 0xac:
		return "lodsb"
	case 0xad:
		return "lodsw"
	case 0xaa:
		return "stosb"
	case 0xab:
		return "stosw"
	}
	return "NOT FOUND"
}
