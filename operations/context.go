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

// operation Objects
var (
	mov    MOV
	add    ADD
	intOpe INT
)

// operation mask list
var (
	maskMid3 = byte(0x38)
	maskLow3 = byte(0x07)
	maskTop2 = byte(0xc0)
)

var opeMap = map[byte]func(*Context, byte) (int, string){

	// mov
	0xb0: mov.Analyze,
	0xb1: mov.Analyze,
	0xb2: mov.Analyze,
	0xb3: mov.Analyze,
	0xb4: mov.Analyze,
	0xb5: mov.Analyze,
	0xb6: mov.Analyze,
	0xb7: mov.Analyze,
	0xb8: mov.Analyze,
	0xb9: mov.Analyze,
	0xba: mov.Analyze,
	0xbb: mov.Analyze,
	0xbc: mov.Analyze,
	0xbd: mov.Analyze,
	0xbe: mov.Analyze,
	0xbf: mov.Analyze,

	// int
	0xcc: intOpe.Analyze,
	0xcd: intOpe.Analyze,

	// add
	0x00: add.Analyze,
	0x01: add.Analyze,
	0x02: add.Analyze,
	0x03: add.Analyze,
	0x04: add.Analyze,
	0x05: add.Analyze,
}

// Disassemble exec disassemble
func (ctx *Context) Disassemble(body []byte) {
	ctx.Body = body
	ctx.Idx = 0
	ctx.IdxByte = body[0]
	for {
		if ctx.Idx >= len(body) {
			break
		}
		f := opeMap[ctx.Body[ctx.Idx]]
		if f == nil {
			ctx.Idx++
			continue
		}
		next, ope := f(ctx, ctx.Body[ctx.Idx])
		ctx.Idx = ctx.Idx + next
		fmt.Println(ope)
	}
}
