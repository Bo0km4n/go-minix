#include "VM.h"
#include "../UnixBase.h"
#include "disasm.h"
#include "regs.h"
#include <stdio.h>
#include <stdlib.h>

using namespace PDP11;

void VM::run1() {
    OpCode *op, op1;
    if (cache.empty()) {
        op = &(op1 = disasm1(text, PC));
    } else {
        op = &cache[PC];
        if (op->empty()) *op = disasm1(text, PC);
    }
    if (PC + op->len > 0x10000) {
        fprintf(stderr, "overrun: %04x\n", PC);
        hasExited = true;
        return;
    }
    if (SP < brksize) {
        fprintf(stderr, "stack overflow: %04x\n", SP);
        hasExited = true;
        return;
    }
    if (trace >= 2) debug(PC, *op);
    uint16_t w = ::read16(text + PC);
    uint16_t oldpc = PC;
    int dst, src, val;
    int16_t val16;
    int8_t val8;
    PC += op->len;
    switch (w >> 12) {
        case 000:
            switch ((w >> 6) & 077) {
                case 000:
                    switch (w & 077) {
                        case 0: // halt
                        case 1: // wait
                        case 2: // rti
                        case 3: // bpt
                        case 4: // iot
                        case 5: // reset
                        case 6: // rtt
                            break;
                    }
                    break;
                case 001: // jmp: JuMP
                    r[7] = addr(op->opr1);
                    return;
                case 002:
                    switch ((w >> 3) & 7) {
                        case 0: // rts: ReTurn from Subroutine
                            r[7] = r[op->opr1.reg];
                            r[op->opr1.reg] = read16(SP);
                            SP += 2;
                            return;
                        case 3: // spl
                            break;
                        case 4:
                        case 5:
                        case 6:
                        case 7:
                            // nop/cl*/se*/ccc/scc: CLear/SEt (Condition Codes)
                            bool f = (w & 16) != 0;
                            if (w & 8) N = f;
                            if (w & 4) Z = f;
                            if (w & 2) V = f;
                            if (w & 1) C = f;
                            return;
                    }
                    break;
                case 003: // swab: SWAp Bytes
                {
                    uint16_t val0 = get16(op->opr1, true);
                    uint16_t bh = (val0 >> 8) & 0xff;
                    uint16_t bl = val0 & 0xff;
                    int val1 = (bl << 8) | bh;
                    set16(op->opr1, val1);
                    setZNCV(val1 == 0, (val1 & 0x8000) != 0, false, false);
                    return;
                }
                case 004: // br: BRanch
                case 005:
                case 006:
                case 007:
                    PC = op->opr1.value;
                    return;
                case 010: // bne: Branch if Not Equal
                case 011:
                case 012:
                case 013:
                    if (!Z) PC = op->opr1.value;
                    return;
                case 014: // beq: Branch if EQual
                case 015:
                case 016:
                case 017:
                    if (Z) PC = op->opr1.value;
                    return;
                case 020: // bge: Branch if Greater or Equal
                case 021:
                case 022:
                case 023:
                    if (!(N ^ V)) PC = op->opr1.value;
                    return;
                case 024: // blt: Branch if Less Than
                case 025:
                case 026:
                case 027:
                    if (N ^ V) PC = op->opr1.value;
                    return;
                case 030: // bgt: Branch if Greater Than
                case 031:
                case 032:
                case 033:
                    if (!(Z || (N ^V))) PC = op->opr1.value;
                    return;
                case 034: // ble: Branch if Less or Equal
                case 035:
                case 036:
                case 037:
                    if (Z || (N ^ V)) PC = op->opr1.value;
                    return;
                case 040: // jsr: Jump to SubRoutine
                case 041:
                case 042:
                case 043:
                case 044:
                case 045:
                case 046:
                case 047:
                    val = addr(op->opr2);
                    write16(SP -= 2, r[op->opr1.reg]);
                    r[op->opr1.reg] = PC;
                    PC = val;
                    return;
                case 050: // clr: CLeaR
                    set16(op->opr1, 0);
                    setZNCV(true, false, false, false);
                    return;
                case 051: // com: COMplement
                    val = ~get16(op->opr1, true);
                    set16(op->opr1, val);
                    setZNCV(val == 0, (val & 0x8000) != 0, true, false);
                    return;
                case 052: // inc: INCrement
                    val = int(int16_t(get16(op->opr1, true))) + 1;
                    set16(op->opr1, val);
                    setZNCV(val == 0, val < 0, C, val == 0x8000);
                    return;
                case 053: // dec: DECrement
                    val = int(int16_t(get16(op->opr1, true))) - 1;
                    set16(op->opr1, val);
                    setZNCV(val == 0, val < 0, C, val == -0x8001);
                    return;
                case 054: // neg: NEGate
                    val = -int16_t(get16(op->opr1, true));
                    set16(op->opr1, val);
                    setZNCV(val == 0, val < 0, val != 0, val == 0x8000);
                    return;
                case 055: // adc: ADd Carry
                    val = int(int16_t(get16(op->opr1, true))) + int(C);
                    set16(op->opr1, val);
                    setZNCV(val == 0, val < 0, C && val == 0, val == 0x8000);
                    return;
                case 056: // sbc: SuBtract Carry
                    val = int(int16_t(get16(op->opr1, true))) - int(C);
                    set16(op->opr1, val);
                    setZNCV(val == 0, val < 0, C && val == -1, val == -0x8001);
                    return;
                case 057: // tst: TeST
                    val = int16_t(get16(op->opr1));
                    setZNCV(val == 0, val < 0, false, false);
                    return;
                case 060: // ror: ROtate Right
                {
                    int val0 = get16(op->opr1, true);
                    int val1 = (val0 >> 1) | (C ? 0x8000 : 0);
                    set16(op->opr1, val1);
                    bool lsb0 = (val0 & 1) != 0;
                    bool msb1 = C;
                    setZNCV(val1 == 0, msb1, lsb0, msb1 != lsb0);
                    return;
                }
                case 061: // rol: ROtate Left
                {
                    int val0 = get16(op->opr1, true);
                    int val1 = uint16_t(val0 << 1) | (C ? 1 : 0);
                    set16(op->opr1, val1);
                    bool msb0 = (val0 & 0x8000) != 0;
                    bool msb1 = (val1 & 0x8000) != 0;
                    setZNCV(val1 == 0, msb1, msb0, msb1 != msb0);
                    return;
                }
                case 062: // asr: Arithmetic Shift Right
                {
                    int val0 = get16(op->opr1, true);
                    int val1 = int16_t(val0) >> 1;
                    set16(op->opr1, val1);
                    bool lsb0 = (val0 & 1) != 0;
                    bool msb1 = val1 < 0;
                    setZNCV(val1 == 0, msb1, lsb0, msb1 != lsb0);
                    return;
                }
                case 063: // asl: Arithmetic Shift Left
                {
                    int val0 = get16(op->opr1, true);
                    int val1 = uint16_t((uint32_t(val0) << 1) & 0xffff);
                    set16(op->opr1, val1);
                    bool msb0 = (val0 & 0x8000) != 0;
                    bool msb1 = val1 < 0;
                    setZNCV(val1 == 0, msb1, msb0, msb1 != msb0);
                    return;
                }
                case 064: // mark: MARK
                    val = w & 077;
                    r[6] = uint16_t((r[6] + 2 * val) & 0xffff);
                    r[7] = r[5];
                    r[5] = read16(r[6]);
                    r[6] += 2;
                    return;
                case 065: // mfpi
                case 066: // mtpi
                    break;
                case 067: // sxt: Sign eXTend
                    set16(op->opr1, -int(N));
                    setZNCV(!N, N, C, V);
                    return;
            }
            break;
        case 001: // mov: MOVe
            src = get16(op->opr1);
            set16(op->opr2, src);
            setZNCV(src == 0, int16_t(src) < 0, C, false);
            return;
        case 002: // cmp: CoMPare
            src = get16(op->opr1);
            dst = get16(op->opr2);
            val16 = val = int16_t(src) - int16_t(dst);
            setZNCV(val16 == 0, val16 < 0, src < dst, val != val16);
            return;
        case 003: // bit: BIt Test
            val = get16(op->opr1) & get16(op->opr2);
            setZNCV(val == 0, (val & 0x8000) != 0, C, false);
            return;
        case 004: // bic: BIt Clear
            val = (~get16(op->opr1)) & get16(op->opr2, true);
            set16(op->opr2, val);
            setZNCV(val == 0, (val & 0x8000) != 0, C, false);
            return;
        case 005: // bis: BIt Set
            val = get16(op->opr1) | get16(op->opr2, true);
            set16(op->opr2, val);
            setZNCV(val == 0, (val & 0x8000) != 0, C, false);
            return;
        case 006: // add: ADD
            src = get16(op->opr1);
            dst = get16(op->opr2, true);
            val16 = val = int16_t(src) + int16_t(dst);
            set16(op->opr2, val16);
            setZNCV(val16 == 0, val16 < 0, src + dst >= 0x10000, val != val16);
            return;
        case 007:
            switch ((w >> 9) & 7) {
                case 0: // mul:MULtiply
                {
                    int src = int16_t(get16(op->opr1));
                    int reg = op->opr2.reg;
                    val = int(r[reg]) * src;
                    if ((reg & 1) == 0) {
                        setReg32(reg, val);
                    } else {
                        r[reg] = val;
                    }
                    setZNCV(val == 0, val < 0, val < -0x8000 || val >= 0x8000, false);
                    return;
                }
                case 1: // div: DIVide
                {
                    int src = int16_t(get16(op->opr1));
                    int reg = op->opr2.reg;
                    if (src == 0 || abs(int16_t(r[reg])) > abs(src)) {
                        setZNCV(false, false, src == 0, true);
                    } else {
                        val = getReg32(reg);
                        int r1 = val / src;
                        r[reg] = r1;
                        r[(reg + 1) & 7] = val % src;
                        setZNCV(r1 == 0, r1 < 0, false, false);
                    }
                    return;
                }
                case 2: // ash: Arithmetic SHift
                {
                    int src = get16(op->opr1) & 077;
                    int reg = op->opr2.reg;
                    int16_t val0 = r[reg];
                    if (src == 0)
                        setZNCV(val0 == 0, val0 < 0, C, false);
                    else if ((src & 040) == 0) {
                        int16_t val1 = val0 << (src - 1);
                        int16_t val2 = val1 << 1;
                        r[reg] = val2;
                        setZNCV(val2 == 0, val2 < 0, val1 < 0, (val0 < 0) != (val2 < 0));
                    } else {
                        int16_t val1 = val0 >> (63 - src);
                        int16_t val2 = val1 >> 1;
                        r[reg] = val2;
                        setZNCV(val2 == 0, val2 < 0, (val1 & 1) != 0, (val0 < 0) != (val2 < 0));
                    }
                    return;
                }
                case 3: // ashc: Arithmetic SHift Combined
                {
                    int src = get16(op->opr1) & 077;
                    int reg = op->opr2.reg;
                    int32_t val0 = getReg32(reg);
                    if (src == 0)
                        setZNCV(val0 == 0, val0 < 0, C, false);
                    else if ((src & 040) == 0) {
                        int32_t val1 = val0 << (src - 1);
                        int32_t val2 = val1 << 1;
                        setReg32(reg, val2);
                        setZNCV(val2 == 0, val2 < 0, val1 < 0, (val0 < 0) != (val2 < 0));
                    } else {
                        int32_t val1 = val0 >> (63 - src);
                        int32_t val2 = val1 >> 1;
                        setReg32(reg, val2);
                        setZNCV(val2 == 0, val2 < 0, (val1 & 1) != 0, (val0 < 0) != (val2 < 0));
                    }
                    return;
                }
                case 4: // xor: eXclusive OR
                    val = r[op->opr1.reg] ^ get16(op->opr2, true);
                    set16(op->opr2, val);
                    setZNCV(val == 0, (val & 0x8000) != 0, C, false);
                    return;
                case 5: // fadd/fsub/fmul/fdiv
                    break;
                case 7: // sob: Subtract One from register, Branch if not zero
                    r[op->opr1.reg]--;
                    if (r[op->opr1.reg] != 0) PC -= op->opr2.value * 2;
                    return;
            }
            break;
        case 010:
            switch ((w >> 6) & 077) {
                case 000: // bpl: Branch if PLus
                case 001:
                case 002:
                case 003:
                    if (!N) PC = op->opr1.value;
                    return;
                case 004: // bmi: Branch if MInus
                case 005:
                case 006:
                case 007:
                    if (N) PC = op->opr1.value;
                    return;
                case 010: // bhi: Branch if HIgher
                case 011:
                case 012:
                case 013:
                    if (!(C | Z)) PC = op->opr1.value;
                    return;
                case 014: // blos: Branch if LOwer or Same
                case 015:
                case 016:
                case 017:
                    if (C | Z) PC = op->opr1.value;
                    return;
                case 020: // bvc: Branch if oVerflow Clear
                case 021:
                case 022:
                case 023:
                    if (!V) PC = op->opr1.value;
                    return;
                case 024:
                case 025:
                case 026:
                case 027: // bvs
                case 0x85: // bvs: Branch if oVerflow Set
                    if (V) PC = op->opr1.value;
                    return;
                case 030: // bcc: Branch if Carry Clear
                case 031:
                case 032:
                case 033:
                    if (!C) PC = op->opr1.value;
                    return;
                case 034: // bcs: Branch if Carry Set
                case 035:
                case 036:
                case 037:
                    if (C) PC = op->opr1.value;
                    return;
                case 040: // emt
                case 041:
                case 042:
                case 043:
                    break;
                case 044: // sys(trap)
                case 045:
                case 046:
                case 047:
                    unix->syscall(w & 255);
                    return;
                case 050: // clrb: CLeaR Byte
                    set8(op->opr1, 0);
                    setZNCV(true, false, false, false);
                    return;
                case 051: // comb: COMplement Byte
                    val = ~get8(op->opr1, true);
                    set8(op->opr1, val);
                    setZNCV(val == 0, (val & 0x80) != 0, true, false);
                    return;
                case 052: // incb: INCrement Byte
                    val = int(int8_t(get8(op->opr1, true))) + 1;
                    set8(op->opr1, val);
                    setZNCV(val == 0, val < 0, C, val == 0x80);
                    return;
                case 053: // decb: DECrement Byte
                    val = int(int8_t(get8(op->opr1, true))) - 1;
                    set8(op->opr1, val);
                    setZNCV(val == 0, val < 0, C, val == -0x81);
                    return;
                case 054: // negb: NEGate Byte
                {
                    int val0 = get8(op->opr1, true);
                    int val1 = -int8_t(val0);
                    set8(op->opr1, val1);
                    setZNCV(val1 == 0, val1 < 0, val1 != 0, val1 == 0x80);
                    return;
                }
                case 055: // adcb: ADd Carry Byte
                    val = int(int8_t(get8(op->opr1, true))) + (C ? 1 : 0);
                    set8(op->opr1, val);
                    setZNCV(val == 0, val < 0, C && val == 0, val == 0x80);
                    return;
                case 056: // sbcb: SuBtract Carry Byte
                    val = int(int8_t(get8(op->opr1, true))) - (C ? 1 : 0);
                    set8(op->opr1, val);
                    setZNCV(val == 0, val < 0, C && val == -1, val == -0x81);
                    return;
                case 057: // tstb: TeST Byte
                    val = int8_t(get8(op->opr1));
                    setZNCV(val == 0, val < 0, false, false);
                    return;
                case 060: // rorb: ROtate Right Byte
                {
                    int val0 = get8(op->opr1, true);
                    int val1 = val0 >> 1;
                    if (C) val1 = uint8_t(val1 + 0x80);
                    set8(op->opr1, val1);
                    bool lsb0 = (val0 & 1) != 0;
                    bool msb1 = C;
                    setZNCV(val1 == 0, msb1, lsb0, msb1 != lsb0);
                    return;
                }
                case 061: // rolb: ROtate Left Byte
                {
                    int val0 = get8(op->opr1, true);
                    int val1 = uint8_t(((uint32_t(val0) << 1) + (C ? 1 : 0)) & 0xff);
                    set8(op->opr1, val1);
                    bool msb0 = (val0 & 0x80) != 0;
                    bool msb1 = (val1 & 0x80) != 0;
                    setZNCV(val1 == 0, msb1, msb0, msb1 != msb0);
                    return;
                }
                case 062: // asrb: Arithmetic Shift Right Byte
                {
                    int val0 = get8(op->opr1, true);
                    int val1 = int8_t(val0) >> 1;
                    set8(op->opr1, val1);
                    bool lsb0 = (val0 & 1) != 0;
                    bool msb1 = val1 < 0;
                    setZNCV(val1 == 0, msb1, lsb0, msb1 != lsb0);
                    return;
                }
                case 063: // aslb: Arithmetic Shift Left Byte
                {
                    int val0 = get8(op->opr1, true);
                    int val1 = uint8_t((uint32_t(val0) << 1) & 0xff);
                    set8(op->opr1, val1);
                    bool msb0 = (val0 & 0x80) != 0;
                    bool msb1 = val1 < 0;
                    setZNCV(val1 == 0, msb1, msb0, msb1 != msb0);
                    return;
                }
                case 065: // mfpd
                case 066: // mtpd
                    break;
            }
            break;
        case 011: // movb: MOVe Byte
            src = get8(op->opr1);
            set8(op->opr2, src, true);
            setZNCV(src == 0, int8_t(src) < 0, C, false);
            return;
        case 012: // cmpb: CoMPare Byte
            src = get8(op->opr1);
            dst = get8(op->opr2);
            val8 = val = int8_t(src) - int8_t(dst);
            setZNCV(val8 == 0, val8 < 0, src < dst, val != val8);
            return;
        case 013: // bitb: BIt Test Byte
            val = get8(op->opr1) & get8(op->opr2);
            setZNCV(val == 0, (val & 0x80) != 0, C, false);
            return;
        case 014: // bicb: BIt Clear Byte
            val = (~get8(op->opr1)) & get8(op->opr2, true);
            set8(op->opr2, val);
            setZNCV(val == 0, (val & 0x80) != 0, C, false);
            return;
        case 015: // bisb: BIt Set Byte
            val = get8(op->opr1) | get8(op->opr2, true);
            set8(op->opr2, val);
            setZNCV(val == 0, (val & 0x80) != 0, C, false);
            return;
        case 016: // sub: SUBtract
            src = get16(op->opr1);
            dst = get16(op->opr2);
            val16 = val = int16_t(dst) - int16_t(src);
            set16(op->opr2, val16);
            setZNCV(val16 == 0, val16 < 0, dst < src, val != val16);
            return;
        case 017:
            switch (w) {
                case 0170011: return; // setd: SET Double
            }
            break;
    }
    if (trace < 2) {
        fprintf(stderr, header);
        debug(oldpc, *op);
    }
    fprintf(stderr, "not implemented\n");
    unix->sys_exit(-1);
}
