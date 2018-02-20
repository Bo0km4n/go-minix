.extern _write
_write:
	push bp
	mov bp, sp
	mov ax, 6(bp)
	mov 1f, ax
	mov ax, 8(bp)
	mov 2f, ax
	mov ax, 4(bp)
	int 7
	.data1 0
	.data2 0f
	mov sp, bp
	pop bp
	ret

.sect .data
0:	int 7
	.data1 4
1:	.data2 0
2:	.data2 0
