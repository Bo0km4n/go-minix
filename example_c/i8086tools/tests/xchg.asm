mov ax, 0x1234
mov cx, 0xabcd
xchg ax, cx
push ax
mov ax, 0x5678
mov bp, sp
xchg ax, [bp]
pop ax
ret
test ax, [bp]
xchg [bp], ax
test [bp], ax
xchg ax, bx
xchg bx, ax
