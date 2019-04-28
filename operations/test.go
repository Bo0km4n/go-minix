package operations

type TEST struct{}

// Analyze test analyze
func (test *TEST) Analyze(ctx *Context, inst byte) (int, string) {
	switch inst {
	case 0xa8:
		return NOT_FOUND, ""
	}
	return NOT_FOUND, ""
}
