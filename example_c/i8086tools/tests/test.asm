mov ax, 1
mov ah, 2
mov al, 3
mov bx, 4
mov bh, 5
mov bl, 6
push ax
mov bp, sp
add ax, [bp]
mov cx, [bp]
mov [bp], bh
mov [bp+1], bl
pop di
push cx
mov [bp], word 0x1234
mov al, [0xffff]
add [bp], ah
add bp, [bp]
add ax, 0x7e00
pop cx
ret
