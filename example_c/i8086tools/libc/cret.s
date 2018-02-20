.extern .dsret, .sret, .cret, .csb2

.dsret:
	pop di
.sret:
	pop si
.cret:
	mov sp, bp
	pop bp
	ret

.csb2:
	mov dx, (bx)
	mov cx, 2(bx)
0:	add bx, #4
	cmp ax, (bx)
	jnz 1f
	mov dx, 2(bx)
	jmp (dx)
1:	loop 0b
	jmp (dx)
