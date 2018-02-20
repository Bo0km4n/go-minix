.extern _main
_main:
mov ax, #6
push ax
mov ax, #hello
push ax
mov ax, #1
push ax
call _write
add sp, #6

mov ax, #0
push ax
call _exit

.sect .data
hello: .data1 'h', 'e', 'l', 'l', 'o', '\n'
