#pragma once
#include "Operand.h"
#include <stdint.h>

namespace PDP11 {

    struct OpCode {
        size_t len;
        const char *mne;
        Operand opr1, opr2;

        OpCode();
        OpCode(int len, const char *mne,
                const Operand &opr1 = Operand(),
                const Operand &opr2 = Operand());

        inline bool empty() const {
            return len == 0;
        }

        bool undef() const;
        std::string str() const;
    };

    extern OpCode undefop;
}
