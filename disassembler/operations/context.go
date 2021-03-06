package operations

import (
	"fmt"
)

// Context model
// Binary scanner
// disassemble binary to assembler code
type Context struct {
	Idx  int // how many read bytes
	Body []byte
}

// Disassemble exec disassemble
func (ctx *Context) Disassemble(body []byte) {
	ctx.Body = body
	ctx.Idx = 0
	for {
		if ctx.Idx >= len(body)-1 {
			break
		}
		f := opeMap[ctx.Body[ctx.Idx]]
		if f == nil {
			ctx.Idx++
			panic("undefined function")
		}
		offset, ope := f(ctx, ctx.Body[ctx.Idx])
		if offset < 0 {
			panic(fmt.Errorf("Not found operator: %02x", ctx.Body[ctx.Idx]))
		}
		ctx.Idx = ctx.Idx + offset
		fmt.Println(ope)
	}
}
