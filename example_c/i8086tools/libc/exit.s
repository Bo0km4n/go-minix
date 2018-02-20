.extern _exit
_exit:
	mov bp, sp
	mov ax, 2(bp)
	int 7
	.data1 1
