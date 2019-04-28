package operations

const (
	NOT_FOUND = -1
)

// operation Objects
var (
	mov    MOV
	add    ADD
	and    AND
	intOpe INT
	push   PUSH
	call   CALL
	grp    GRP
	jmp    JMP
	in     IN
	sbb    SBB
	lea    LEA
	pop    POP
	or     OR
	je     JE
	cmp    CMP
	jnl    JNL
	ret    RET
	cbw    CBW
	jne    JNE
	inc    INC
	xor    XOR
	sub    SUB
	jl     JL
	jb     JB
	jnle   JNLE
	jle    JLE
	jnb    JNB
	jnbe   JNBE
	jbe    JBE
	hlt    HLT
	dec    DEC
	cwd    CWD
	cld    CLD
	test   TEST
	rep    REP
	adc    ADC
	xchg   XCHG
)

// operation mask list
var (
	maskMid3 = byte(0x38)
	maskLow3 = byte(0x07)
	maskTop2 = byte(0xc0)
)

var opeMap = map[byte]func(*Context, byte) (int, string){

	// push
	0x50: push.Analyze,
	0x51: push.Analyze,
	0x52: push.Analyze,
	0x53: push.Analyze,
	0x54: push.Analyze,
	0x55: push.Analyze,
	0x56: push.Analyze,
	0x57: push.Analyze,

	// mov
	0x88: mov.Analyze,
	0x89: mov.Analyze,
	0x8a: mov.Analyze,
	0x8b: mov.Analyze,
	0xb8: mov.Analyze,
	0xb9: mov.Analyze,
	0xba: mov.Analyze,
	0xbb: mov.Analyze,
	0xc6: mov.Analyze,
	0xc7: mov.Analyze,

	// int
	// 0xcc: intOpe.Analyze,
	0xcd: intOpe.Analyze,

	// add
	0x00: add.Analyze,
	0x01: add.Analyze,
	0x03: add.Analyze,

	// and
	0x20: and.Analyze,
	0x21: and.Analyze,
	0x22: and.Analyze,
	0x23: and.Analyze,

	// call
	0xe8: call.Analyze,

	// grp
	0x80: grp.Analyze,
	0x81: grp.Analyze,
	0x83: grp.Analyze,
	0xf7: grp.Analyze,
	0xf6: grp.Analyze,
	0xd1: grp.Analyze,
	0xff: grp.Analyze,

	// jmp
	0xe9: jmp.Analyze,
	0xeb: jmp.Analyze,

	// in
	0xe5: in.Analyze,
	0xec: in.Analyze,

	// sbb
	0x18: sbb.Analyze,

	// lea
	0x8d: lea.Analyze,

	// pop
	0x58: pop.Analyze,
	0x59: pop.Analyze,
	0x5a: pop.Analyze,
	0x5b: pop.Analyze,
	0x5d: pop.Analyze,
	0x5e: pop.Analyze,
	0x5f: pop.Analyze,

	// or
	0x08: or.Analyze,
	0x09: or.Analyze,
	0x0a: or.Analyze,
	0x0b: or.Analyze,

	// je
	0x74: je.Analyze,

	// jnl
	0x7d: jnl.Analyze,

	// jnle
	0x7f: jnle.Analyze,

	// ret
	0xc3: ret.Analyze,
	0xc2: ret.Analyze,

	// cbw
	0x98: cbw.Analyze,

	// jne
	0x75: jne.Analyze,

	// inc
	0x40: inc.Analyze,
	0x41: inc.Analyze,
	0x42: inc.Analyze,
	0x43: inc.Analyze,
	0x44: inc.Analyze,
	0x45: inc.Analyze,
	0x46: inc.Analyze,
	0x47: inc.Analyze,

	// xor
	0x30: xor.Analyze,
	0x31: xor.Analyze,
	0x32: xor.Analyze,
	0x33: xor.Analyze,

	// sub
	0x28: sub.Analyze,
	0x29: sub.Analyze,
	0x2a: sub.Analyze,
	0x2b: sub.Analyze,
	0x2c: sub.Analyze,
	0x2d: sub.Analyze,

	// jl
	0x7c: jl.Analyze,

	// jnb
	0x73: jnb.Analyze,

	// hlt
	0xf4: hlt.Analyze,

	// dec
	0x48: dec.Analyze,
	0x49: dec.Analyze,
	0x4a: dec.Analyze,
	0x4b: dec.Analyze,
	0x4c: dec.Analyze,
	0x4d: dec.Analyze,
	0x4e: dec.Analyze,
	0x4f: dec.Analyze,

	// cwd
	0x99: cwd.Analyze,

	// cmp
	0x38: cmp.Analyze,
	0x39: cmp.Analyze,
	0x3a: cmp.Analyze,
	0x3b: cmp.Analyze,
	0x3d: cmp.Analyze,

	// jb
	0x72: jb.Analyze,

	// jle
	0x7e: jle.Analyze,

	// jnbe
	0x77: jnbe.Analyze,

	// jbe
	0x76: jbe.Analyze,

	// cld
	0xfc: cld.Analyze,

	// test
	0x85: test.Analyze,
	0xa8: test.Analyze,

	// rep
	0xf2: rep.Analyze,
	0xf3: rep.Analyze,

	// adc
	0x11: adc.Analyze,

	// std
	0xfd: func(ctx *Context, inst byte) (int, string) {
		return 1, getResult(ctx.Idx, getOrgOpe(ctx.Body[ctx.Idx:ctx.Idx+1]), "std")
	},

	// xchg
	0x86: xchg.Analyze,
	0x87: xchg.Analyze,
	0x90: xchg.Analyze,
	0x91: xchg.Analyze,
	0x92: xchg.Analyze,
	0x93: xchg.Analyze,
	0x94: xchg.Analyze,
	0x95: xchg.Analyze,
	0x96: xchg.Analyze,
	0x97: xchg.Analyze,
}
