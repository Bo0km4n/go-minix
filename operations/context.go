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
	push   PUSH
	call   CALL
	grp    GRP
	jmp    JMP
	in     IN
	sbb    SBB
	lea    LEA
)

// operation mask list
var (
	maskMid3 = byte(0x38)
	maskLow3 = byte(0x07)
	maskTop2 = byte(0xc0)
)

var opeMap = map[byte]func(*Context, byte) (int, string){

	// push
	// 0x06: push.Analyze,
	// 0x16: push.Analyze,
	0x50: push.Analyze,
	// 0x51: push.Analyze,
	// 0x52: push.Analyze,
	// 0x53: push.Analyze,
	// 0x54: push.Analyze,
	0x55: push.Analyze,
	0x56: push.Analyze,
	0x57: push.Analyze,
	0xff: push.Analyze,

	// mov
	0x89: mov.Analyze,
	0x8b: mov.Analyze,
	// 0xb0: mov.Analyze,
	// 0xb1: mov.Analyze,
	// 0xb2: mov.Analyze,
	// 0xb3: mov.Analyze,
	// 0xb4: mov.Analyze,
	// 0xb5: mov.Analyze,
	// 0xb6: mov.Analyze,
	// 0xb7: mov.Analyze,
	0xb8: mov.Analyze,
	// 0xb9: mov.Analyze,
	// 0xba: mov.Analyze,
	0xbb: mov.Analyze,
	// 0xbc: mov.Analyze,
	// 0xbd: mov.Analyze,
	// 0xbe: mov.Analyze,
	// 0xbf: mov.Analyze,

	// int
	// 0xcc: intOpe.Analyze,
	0xcd: intOpe.Analyze,

	// add
	0x00: add.Analyze,

	// call
	0xe8: call.Analyze,

	// grp
	0x83: grp.Analyze,

	// jmp
	0xe9: jmp.Analyze,

	// in
	0xe5: in.Analyze,
	0xec: in.Analyze,

	// sbb
	0x18: sbb.Analyze,

	// lea
	0x8d: lea.Analyze,
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
			fmt.Println("undefined function")
			break
		}
		next, ope := f(ctx, ctx.Body[ctx.Idx])
		ctx.Idx = ctx.Idx + next
		fmt.Println(ope)
	}
}
