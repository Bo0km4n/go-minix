sub sp, 20
mov bx, sp

mov word [bx +  2], 4
mov word [bx +  4], 1
mov word [bx +  6], 6
mov word [bx + 10], hello
int 0x20

mov word [bx +  2], 1
mov word [bx +  4], 0
int 0x20

hello: db "hello", 10
