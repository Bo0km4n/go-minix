package operations

import (
	"fmt"
)

// Context model
// Binary scanner
// disassemble binary to assembler code
type Context struct {
	Idx     int  // how many read bytes
	IdxByte byte // cursor byte
	Body    []byte
}

// Disassemble exec disassemble
func (ctx *Context) Disassemble(body []byte) {
	ctx.Body = body
	ctx.Idx = 0
	ctx.IdxByte = body[0]
	for {
		if ctx.Idx >= len(body)-1 {
			break
		}
		f := opeMap[ctx.Body[ctx.Idx]]
		if f == nil {
			ctx.Idx++
			fmt.Println("undefined function")
			break
		}
		next, ope := f(ctx, ctx.Body[ctx.Idx])
		ctx.Idx = ctx.Idx + next
		fmt.Println(ope)
	}
}
