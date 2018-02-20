#include "VM.h"
#include "disasm.h"
#include "regs.h"
#include <stdio.h>
#include <string.h>
#include <sys/stat.h>

using namespace PDP11;

const char *PDP11::header = " r0   r1   r2   r3   r4   r5   sp  flags pc\n";

void VM::showHeader() {
    fprintf(stderr, header);
}

void VM::debug(uint16_t pc, const OpCode &op) {
    debugsym(pc);
    fprintf(stderr,
            "%04x %04x %04x %04x %04x %04x %04x %c%c%c%c %04x:%-14s %s",
            r[0], r[1], r[2], r[3], r[4], r[5], r[6],
            "-Z"[Z], "-N"[N], "-C"[C], "-V"[V],
            pc, hexdump2(text + pc, op.len).c_str(), op.str().c_str());
    if (trace >= 3) {
        uint16_t r[8];
        memcpy(r, this->r, sizeof (r));
        int ad1 = addr(op.opr1);
        int ad2 = addr(op.opr2);
        if (ad1 >= 0) {
            if (op.opr1.w)
                fprintf(stderr, " ;[%04x]%04x", ad1, read16(ad1));
            else
                fprintf(stderr, " ;[%04x]%02x", ad1, read8(ad1));
        }
        if (ad2 >= 0) {
            if (op.opr2.w)
                fprintf(stderr, " ;[%04x]%04x", ad2, read16(ad2));
            else
                fprintf(stderr, " ;[%04x]%02x", ad2, read8(ad2));
        }
        memcpy(this->r, r, sizeof (r));
    }
    fprintf(stderr, "\n");
}

VM::VM() : start_sp(0) {
    memset(r, 0, sizeof (r));
    Z = N = C = V = false;
}

VM::VM(const VM &vm) : VMBase(vm) {
    memcpy(r, vm.r, sizeof (r));
    Z = vm.Z;
    N = vm.N;
    C = vm.C;
    V = vm.V;
    start_sp = vm.start_sp;
}

VM::~VM() {
}

bool VM::load(const std::string& fn, FILE* f, size_t size) {
    if (!VMBase::load(fn, f, size)) return false;
    PC = 0;
    cache.clear();
    return true;
}

int VM::addr(const Operand &opr, bool nomove) {
    if (opr.reg == 7) {
        switch (opr.mode) {
            case 3:
            case 6: return opr.value;
            case 7: return read16(opr.value);
        }
        return -1;
    }
    if (nomove) {
        switch (opr.mode) {
            case 1: return r[opr.reg];
            case 2: return r[opr.reg];
            case 3: return read16(r[opr.reg]);
            case 4: return uint16_t(r[opr.reg] - opr.diff());
            case 5: return uint16_t(read16(r[opr.reg] - opr.diff()));
            case 6: return uint16_t(r[opr.reg] + opr.value);
            case 7: return uint16_t(read16(r[opr.reg] + opr.value));
        }
    }
    switch (opr.mode) {
        case 1: return r[opr.reg];
        case 2: return getInc(opr);
        case 3: return read16(getInc(opr));
        case 4: return getDec(opr);
        case 5: return read16(getDec(opr));
        case 6: return uint16_t(r[opr.reg] + opr.value);
        case 7: return uint16_t(read16(r[opr.reg] + opr.value));
    }
    return -1;
}

void VM::run2() {
    while (!hasExited) run1();
}

void VM::disasm() {
    int addr = 0, undef = 0;
    while (addr < (int) tsize) {
        showsym(addr);
        OpCode op = disasm1(text, addr);
        disout(text, addr, op.len, disstr(op));
        if (op.undef()) undef++;
        addr += op.len;
    }
    if (undef) printf("undefined: %d\n", undef);
}

std::string VM::disstr(const OpCode &op) {
    std::string ret = op.str();
    if (op.opr1.isaddr()) {
        std::map<int, Symbol>::iterator it = syms[1].find(op.opr1.value);
        if (it != syms[1].end()) {
            ret += " ; " + it->second.name;
        }
    }
    if (op.opr2.isaddr()) {
        std::map<int, Symbol>::iterator it = syms[1].find(op.opr2.value);
        if (it != syms[1].end()) {
            ret += " ; " + it->second.name;
        }
    }
    return ret;
}
