#include "Operand.h"
#include "disasm.h"

using namespace PDP11;

Operand::Operand()
: len(-1), mode(0), reg(0), value(0), w(true) {
}

Operand::Operand(int len, int mode, int reg, int value)
: len(len), mode(mode), reg(reg), value(value), w(true) {
}

Operand::Operand(uint8_t *mem, int pc, int modr)
: w(true) {
    mode = (modr >> 3) & 7;
    reg = modr & 7;
    if (reg == 7) {
        switch (mode) {
            case 0:
                len = 0;
                value = pc;
                return;
            case 1:
                len = 0;
                value = read16(mem);
                return;
            case 2:
            case 3:
                len = 2;
                value = read16(mem);
                return;
            case 6:
            case 7:
                len = 2;
                value = uint16_t(pc + 2 + read16(mem));
                return;
        }
    }
    if (mode < 6) {
        len = value = 0;
    } else {
        len = 2;
        value = (int16_t) read16(mem);
    }
}

std::string Operand::str() const {
    if (reg == 7) {
        switch (mode) {
            case 2: return "$" + hex(value);
            case 3: return "*$" + hex(value, 4);
            case 6: return hex(value, 4);
            case 7: return "*" + hex(value, 4);
            case 8: return hex(value); // imm
            case 9: return hex(value, 4); // address
        }
    }
    const std::string &rn = regs[reg];
    switch (mode) {
        case 0: return rn;
        case 1: return "(" + rn + ")";
        case 2: return "(" + rn + ")+";
        case 3: return "*(" + rn + ")+";
        case 4: return "-(" + rn + ")";
        case 5: return "*-(" + rn + ")";
        case 6: return hex(value) + "(" + rn + ")";
        case 7: return "*" + hex(value) + "(" + rn + ")";
    }
    return "?";
}
