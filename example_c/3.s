.sect .text; .sect .rom; .sect .data; .sect .bss
.extern _a
.sect .data
_a:
.data2	1
.data2	2
.extern _b
.data2	3
.sect .text
_b:
	push	bp
	mov	bp,sp
add _a+2,#2
jmp .cret
