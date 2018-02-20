#pragma once
#include "../VMBase.h"
#include "OpCode.h"

namespace i8086 {
    extern const char *header;

    struct VM : public VMBase {
        uint16_t IP, r[8];
        uint8_t * r8[8];
        bool OF, DF, SF, ZF, AF, PF, CF;
        uint16_t start_sp;

        static bool ptable[256];
        void init();

        VM();
        VM(const VM &vm);
        virtual ~VM();

        virtual bool load(const std::string &fn, FILE *f, size_t size);
        virtual void disasm();
        virtual void showHeader();
        virtual void run2();

        std::string disstr(const OpCode &op);
        void run1(uint8_t rep = 0);
        void debug(uint16_t ip, const OpCode &op);
        int addr(const Operand &opr);
        void shift(Operand *opr, int c, uint8_t *p);

        inline int setf8(int value) {
            int8_t v = value;
            OF = value != v;
            SF = v < 0;
            ZF = v == 0;
            PF = ptable[uint8_t(value)];
            return value;
        }

        inline int setf16(int value) {
            int16_t v = value;
            OF = value != v;
            SF = v < 0;
            ZF = v == 0;
            PF = ptable[uint8_t(value)];
            return value;
        }

        inline uint16_t getf() {
            return 0xf002 | (OF << 11) | (DF << 10) | (SF << 7) |
                    (ZF << 6) | (AF << 4) | (PF << 2) | CF;
        }

        inline void setf(uint16_t flags) {
            CF = flags & 0x001;
            PF = flags & 0x004;
            AF = flags & 0x010;
            ZF = flags & 0x040;
            SF = flags & 0x080;
            DF = flags & 0x400;
            OF = flags & 0x800;
        }

        inline void jumpif(int8_t offset, bool c) {
            if (c) {
                IP += 2 + offset;
            } else {
                IP += 2;
            }
        }

        inline void push(uint16_t val) {
            write16(r[4] -= 2, val);
        }

        inline uint16_t pop() {
            uint16_t val = read16(r[4]);
            r[4] += 2;
            return val;
        }
    };
}
