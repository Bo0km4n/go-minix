start:
	pop ax
	mov dx, sp
	push dx
	push ax
	call _main
	push ax
	call _exit
