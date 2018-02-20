.code16
.intel_syntax noprefix

sub sp, 20
mov bx, sp

mov word ptr [bx +  2], 4
mov word ptr [bx +  4], 1
mov word ptr [bx +  6], 6
mov word ptr [bx + 10], offset hello
int 0x20

mov word ptr [bx +  2], 1
mov word ptr [bx +  4], 0
int 0x20

.data
hello: .ascii "hello\n"
