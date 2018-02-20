#include "disasm.h"
#include <stdio.h>
#include <string.h>

using namespace PDP11;

std::string PDP11::regs[] = {"r0", "r1", "r2", "r3", "r4", "r5", "sp", "pc"};

static inline OpCode srcdst(uint8_t *mem, uint16_t addr, int w, const char *mne) {
    Operand opr1(mem + 2, addr + 2, w >> 6);
    int offset = 2 + opr1.len;
    Operand opr2(mem + offset, addr + offset, w);
    return OpCode(offset + opr2.len, mne, opr1, opr2);
}

static inline OpCode modr(uint8_t *mem, uint16_t addr, int w, const char *mne) {
    Operand opr(mem + 2, addr + 2, w);
    return OpCode(2 + opr.len, mne, opr);
}

static inline OpCode rdst(uint8_t *mem, uint16_t addr, int w, const char *mne) {
    Operand opr(mem + 2, addr + 2, w);
    return OpCode(2 + opr.len, mne, reg(w >> 6), opr);
}

static inline OpCode srcr(uint8_t *mem, uint16_t addr, int w, const char *mne) {
    Operand opr(mem + 2, addr + 2, w);
    return OpCode(2 + opr.len, mne, opr, reg(w >> 6));
}

static inline OpCode branch(uint16_t addr, int w, const char *mne) {
    return OpCode(2, mne, address(addr + 2 + ((int8_t) (w & 255)) * 2));
}

static inline OpCode srcdstb(uint8_t *mem, uint16_t addr, int w, const char *mne) {
    OpCode ret = srcdst(mem, addr, w, mne);
    ret.opr1.w = ret.opr2.w = false;
    return ret;
}

static inline OpCode modrb(uint8_t *mem, uint16_t addr, int w, const char *mne) {
    OpCode ret = modr(mem, addr, w, mne);
    ret.opr1.w = false;
    return ret;
}

OpCode PDP11::disasm1(uint8_t *text, uint16_t addr) {
    uint8_t *mem = text + addr;
    uint16_t w = ::read16(mem);
    switch (w >> 12) {
        case 000:
            switch ((w >> 6) & 077) {
                case 000:
                    switch (w & 077) {
                        case 0: return OpCode(2, "halt");
                        case 1: return OpCode(2, "wait");
                        case 2: return OpCode(2, "rti");
                        case 3: return OpCode(2, "bpt");
                        case 4: return OpCode(2, "iot");
                        case 5: return OpCode(2, "reset");
                        case 6: return OpCode(2, "rtt");
                    }
                    break;
                case 001: return modr(mem, addr, w, "jmp");
                case 002:
                    switch ((w >> 3) & 7) {
                        case 0: return OpCode(2, "rts", reg(w & 7));
                        case 3: return OpCode(2, "spl", imm(w & 7));
                        case 4:
                        case 5:
                        case 6:
                        case 7:
                            switch (w) {
                                case 0240: return OpCode(2, "nop");
                                case 0241: return OpCode(2, "clc");
                                case 0242: return OpCode(2, "clv");
                                case 0243: return OpCode(2, "clvc");
                                case 0244: return OpCode(2, "clz");
                                case 0245: return OpCode(2, "clzc");
                                case 0246: return OpCode(2, "clzv");
                                case 0247: return OpCode(2, "clzvc");
                                case 0250: return OpCode(2, "cln");
                                case 0251: return OpCode(2, "clnc");
                                case 0252: return OpCode(2, "clnv");
                                case 0253: return OpCode(2, "clnvc");
                                case 0254: return OpCode(2, "clnz");
                                case 0255: return OpCode(2, "clnzc");
                                case 0256: return OpCode(2, "clnzv");
                                case 0257: return OpCode(2, "ccc");
                                case 0261: return OpCode(2, "sec");
                                case 0262: return OpCode(2, "sev");
                                case 0263: return OpCode(2, "sevc");
                                case 0264: return OpCode(2, "sez");
                                case 0265: return OpCode(2, "sezc");
                                case 0266: return OpCode(2, "sezv");
                                case 0267: return OpCode(2, "sezvc");
                                case 0270: return OpCode(2, "sen");
                                case 0271: return OpCode(2, "senc");
                                case 0272: return OpCode(2, "senv");
                                case 0273: return OpCode(2, "senvc");
                                case 0274: return OpCode(2, "senz");
                                case 0275: return OpCode(2, "senzc");
                                case 0276: return OpCode(2, "senzv");
                                case 0277: return OpCode(2, "scc");
                            }
                    }
                    break;
                case 003: return modr(mem, addr, w, "swab");
                case 004:
                case 005:
                case 006:
                case 007: return branch(addr, w, "br");
                case 010:
                case 011:
                case 012:
                case 013: return branch(addr, w, "bne");
                case 014:
                case 015:
                case 016:
                case 017: return branch(addr, w, "beq");
                case 020:
                case 021:
                case 022:
                case 023: return branch(addr, w, "bge");
                case 024:
                case 025:
                case 026:
                case 027: return branch(addr, w, "blt");
                case 030:
                case 031:
                case 032:
                case 033: return branch(addr, w, "bgt");
                case 034:
                case 035:
                case 036:
                case 037: return branch(addr, w, "ble");
                case 040:
                case 041:
                case 042:
                case 043:
                case 044:
                case 045:
                case 046:
                case 047: return rdst(mem, addr, w, "jsr");
                case 050: return modr(mem, addr, w, "clr");
                case 051: return modr(mem, addr, w, "com");
                case 052: return modr(mem, addr, w, "inc");
                case 053: return modr(mem, addr, w, "dec");
                case 054: return modr(mem, addr, w, "neg");
                case 055: return modr(mem, addr, w, "adc");
                case 056: return modr(mem, addr, w, "sbc");
                case 057: return modr(mem, addr, w, "tst");
                case 060: return modr(mem, addr, w, "ror");
                case 061: return modr(mem, addr, w, "rol");
                case 062: return modr(mem, addr, w, "asr");
                case 063: return modr(mem, addr, w, "asl");
                case 064: return OpCode(2, "mark", imm(w & 077));
                case 065: return modr(mem, addr, w, "mfpi");
                case 066: return modr(mem, addr, w, "mtpi");
                case 067: return modr(mem, addr, w, "sxt");
            }
            break;
        case 001: return srcdst(mem, addr, w, "mov");
        case 002: return srcdst(mem, addr, w, "cmp");
        case 003: return srcdst(mem, addr, w, "bit");
        case 004: return srcdst(mem, addr, w, "bic");
        case 005: return srcdst(mem, addr, w, "bis");
        case 006: return srcdst(mem, addr, w, "add");
        case 007:
            switch ((w >> 9) & 7) {
                case 0: return srcr(mem, addr, w, "mul");
                case 1: return srcr(mem, addr, w, "div");
                case 2: return srcr(mem, addr, w, "ash");
                case 3: return srcr(mem, addr, w, "ashc");
                case 4: return rdst(mem, addr, w, "xor");
                case 7: return OpCode(2, "sob", reg(w >> 6), imm(w & 077));
            }
            break;
        case 010:
            switch ((w >> 6) & 077) {
                case 000:
                case 001:
                case 002:
                case 003: return branch(addr, w, "bpl");
                case 004:
                case 005:
                case 006:
                case 007: return branch(addr, w, "bmi");
                case 010:
                case 011:
                case 012:
                case 013: return branch(addr, w, "bhi");
                case 014:
                case 015:
                case 016:
                case 017: return branch(addr, w, "blos");
                case 020:
                case 021:
                case 022:
                case 023: return branch(addr, w, "bvc");
                case 024:
                case 025:
                case 026:
                case 027: return branch(addr, w, "bvs");
                case 030:
                case 031:
                case 032:
                case 033: return branch(addr, w, "bcc");
                case 034:
                case 035:
                case 036:
                case 037: return branch(addr, w, "bcs");
                case 040:
                case 041:
                case 042:
                case 043: return OpCode(2, "emt", imm(w & 255));
                case 044:
                case 045:
                case 046:
                case 047: return OpCode(2, "sys", imm(w & 255));
                case 050: return modrb(mem, addr, w, "clrb");
                case 051: return modrb(mem, addr, w, "comb");
                case 052: return modrb(mem, addr, w, "incb");
                case 053: return modrb(mem, addr, w, "decb");
                case 054: return modrb(mem, addr, w, "negb");
                case 055: return modrb(mem, addr, w, "adcb");
                case 056: return modrb(mem, addr, w, "sbcb");
                case 057: return modrb(mem, addr, w, "tstb");
                case 060: return modrb(mem, addr, w, "rorb");
                case 061: return modrb(mem, addr, w, "rolb");
                case 062: return modrb(mem, addr, w, "asrb");
                case 063: return modrb(mem, addr, w, "aslb");
                case 064: break;
                case 065: return modr(mem, addr, w, "mfpd");
                case 066: return modr(mem, addr, w, "mtpd");
            }
            break;
        case 011: return srcdstb(mem, addr, w, "movb");
        case 012: return srcdstb(mem, addr, w, "cmpb");
        case 013: return srcdstb(mem, addr, w, "bitb");
        case 014: return srcdstb(mem, addr, w, "bicb");
        case 015: return srcdstb(mem, addr, w, "bisb");
        case 016: return srcdst(mem, addr, w, "sub");
        case 017:
            switch (w) {
                case 0170011: return OpCode(2, "setd");
            }
            break;
    }
    return undefop;
}

void PDP11::disasm(uint8_t *text, size_t size) {
    int addr = 0, undef = 0;
    while (addr < (int) size) {
        OpCode op = disasm1(text, addr);
        disout(text, addr, op.len, op.str());
        if (op.undef()) undef++;
        addr += op.len;
    }
    if (undef) printf("undefined: %d\n", undef);
}

void PDP11::disout(uint8_t *text, uint16_t addr, int len, const std::string &ops) {
    for (int i = 0; i < len; i += 6) {
        int left = len - i;
        if (left > 6) left = 6;
        std::string hex = hexdump2(text + addr + i, left);
        if (i == 0) {
            printf("%04x: %-14s  %s\n", addr, hex.c_str(), ops.c_str());
        } else {
            printf("      %-14s\n", hex.c_str());
        }
    }
}
