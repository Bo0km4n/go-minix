.extern csv, cret

csv:
    pop ax
    push bp
    mov bp, sp
    push cx
    push si
    push di
    sub sp, #2
    jmp (ax)

cret:
    lea sp, -6(bp)
    pop di
    pop si
    pop cx
    pop bp
    ret
