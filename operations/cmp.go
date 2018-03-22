package operations

// CMP model
type CMP struct{}

// Analyze je analyze
func (cmp *CMP) Analyze(ctx *Context, inst byte) (int, string) {
	return 2, ""
}
