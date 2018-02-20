#pragma once
#include "OpCode.h"
#include "../VMBase.h"

namespace PDP11 {
    extern std::string regs[];

    OpCode disasm1(uint8_t *text, uint16_t addr);
    void disasm(uint8_t *text, size_t size);
    void disout(uint8_t *text, uint16_t addr, int len, const std::string &ops);
}
