#!/usr/bin/env python
# This file is in the public domain.
import sys

target = "hello.s"
if len(sys.argv) == 2:
    target = sys.argv[1]
elif len(sys.argv) > 2:
    print "usage: pdp2i8086.py pdp11.s"
    sys.exit(1)

with open(target) as f:
    lines = f.readlines()

def isspace(ch):
    return ch <= ' '

def isletter(ch):
    return ch.isalnum() or ch == '_' or ch == '~'

class Lexer:
    def __init__(self, s):
        self.p = 0
        self.s = s
        self.read()

    def canread(self):
        return self.p < len(self.s)

    def readwhile(self, f):
        ret = ""
        while self.canread() and f(self.s[self.p]):
            ret += self.s[self.p]
            self.p += 1
        return ret

    def read(self):
        if not self.canread():
            self.text = ""
            return False
        ch = self.s[self.p]
        self.p += 1
        if isspace(ch):
            self.readwhile(isspace)
            self.read()
        elif isletter(ch):
            self.text = ch + self.readwhile(isletter)
        elif ch == '.':
            self.text = ch + self.readwhile(str.isalnum)
        else:
            self.text = ch
        return True

regs = { "r0": "ax",
         "r1": "dx",
         "r2": "cx",
         "r3": "si",
         "r4": "di",
         "r5": "bp",
         "r6": "sp",
         "sp": "sp",
         "r7": "ip",
         "pc": "ip" }

def getnum(o):
    if not o.isdigit():
        return o
    d = int(o, 8)
    if d < 10:
        return str(d)
    return "0x%x" % d

written = False

def write(a):
    global written
    sys.stdout.write(a)
    written = True

class Operand:
    def __init__(self, lexer):
        self.s = ""
        self.r = ""
        self.t = lexer.text
        self.m = -1
        if lexer.text == "":
            return
        if lexer.text == "$":
            if not lexer.read():
                return
            self.t += lexer.text
            self.s = "#" + getnum(lexer.text)
            lexer.read()
            return
        self.m = 0
        if lexer.text == "-":
            if not lexer.read():
                return
            self.t += lexer.text
            if lexer.text.isdigit():
                self.s += "-"
            else:
                self.m = 4 # -(R)
        if lexer.text.isdigit():
            self.s += getnum(lexer.text)
            if not lexer.read():
                return
            self.t += lexer.text
        if lexer.text == "(":
            if self.m == 0: self.m = 1
            self.s += lexer.text
            if not lexer.read():
                return
            self.t += lexer.text
        r = lexer.text
        if not regs.has_key(r):
            return
        self.r = regs[r]
        self.s += self.r
        lexer.read()
        if lexer.text == ")":
            self.t += lexer.text
            self.s += lexer.text
            lexer.read()
            if lexer.text == "+":
                self.t += lexer.text
                self.m = 2 # (R)+
                lexer.read()

    def incdec(self):
        return 2 <= self.m <= 5

    def conv(self):
        if self.s[0] == "#":
            write("mov bx, " + self.s + "; ")
        elif self.m < 1:
            return
        else:
            src = self.s
            if self.r != "bp" and self.r != "si" and self.r != "di":
                write("mov bx, " + self.r + "; ")
                src = src.replace(self.r, "bx")
            write("mov bx, " + src + "; ")
        self.s = "bx"
        self.r = "bx"
        self.m = 0

for line in lines:
    written = False
    lexer = Lexer(line)
    while lexer.text != "":
        tok = lexer.text
        lexer.read()
        if tok == ";":
            write("; ")
        elif tok == ".globl":
            if lexer.text != "":
                write(".extern " + lexer.text)
                lexer.read()
        elif tok == ".text":
            write(".sect .text")
        elif tok == ".data":
            write(".sect .data")
        elif tok == ".byte":
            write(".data1 ")
            while lexer.text != "":
                if str.isdigit(lexer.text):
                    write("0x%02x" % int(lexer.text, 8))
                elif lexer.text == ",":
                    write(", ")
                else:
                    break
                lexer.read()
        elif lexer.text == ":":
            if tok[0] != "~":
                write(tok + ":")
                written = lexer.text != ""
            lexer.read()
        elif tok == "jsr":
            lexer.read()
            if lexer.text == "," and lexer.read():
                if lexer.text == "*" and lexer.read():
                    if lexer.text == "$" and lexer.read():
                        pass
                write("call " + lexer.text)
        elif tok == "jmp":
            write("jmp " + lexer.text)
            lexer.read()
        elif tok == "mov":
            src = Operand(lexer)
            if lexer.text == "," and lexer.read():
                dst = Operand(lexer)
                if dst.t == "(sp)" or dst.t == "-(sp)":
                    if dst.t == "(sp)":
                        write("add sp, #2; ")
                    src.conv()
                    write("push " + src.s)
                else:
                    write("mov " + dst.s + ", " + src.s)
        elif tok == "tst":
            src = Operand(lexer)
            if src.t == "(sp)+":
                write("add sp, #2")
            elif src.t == "-(sp)":
                write("sub sp, #2")
            else:
                assert not src.incdec(), line
                write("cmp " + src.s + ", #0")
        elif tok == "cmp":
            src = Operand(lexer)
            if lexer.text == "," and lexer.read():
                dst = Operand(lexer)
                if src.t == "(sp)+" and dst.t == "(sp)+":
                    write("add sp, #4")
                elif src.t == "-(sp)" and dst.t == "-(sp)":
                    write("sub sp, #4")
                else:
                    src.conv()
                    assert not dst.incdec(), line
                    write("cmp " + src.s + ", " + dst.s)
        elif tok == "add" or tok == "sub":
            src = Operand(lexer)
            if lexer.text == "," and lexer.read():
                dst = Operand(lexer)
                src.conv()
                assert not dst.incdec(), line
                write(tok + " " + dst.s + ", " + src.s)
        elif tok == "clr":
            dst = Operand(lexer)
            assert not dst.incdec(), line
            if dst.m == 0:
                write("xor " + dst.s + ", " + dst.s)
            else:
                write("mov " + dst.s + ", #0")
        elif tok == "inc" or tok == "dec":
            dst = Operand(lexer)
            assert not dst.incdec(), line
            write(tok + " " + dst.s)
        elif tok == "asl":
            dst = Operand(lexer)
            assert dst.m == 0, line
            write("sal " + dst.s + ", #1")
        elif tok == "jbr":
            write("jmp " + lexer.text)
            lexer.read()
        elif tok == "jle":
            write("jle " + lexer.text)
            lexer.read()
    if written:
        print
    else:
        write("! " + line)
