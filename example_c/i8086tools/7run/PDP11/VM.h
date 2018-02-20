#pragma once
#include "../VMBase.h"
#include "OpCode.h"

namespace PDP11 {
    extern const char *header;

    struct VM : public VMBase {
        uint16_t r[8];
        bool Z, N, C, V;
        uint16_t start_sp;
        std::vector<OpCode> cache;

        VM();
        VM(const VM &vm);
        virtual ~VM();

        virtual bool load(const std::string &fn, FILE *f, size_t size);
        virtual void disasm();
        virtual void showHeader();
        virtual void run2();

        std::string disstr(const OpCode &op);
        void run1();
        void debug(uint16_t pc, const OpCode &op);
        int addr(const Operand &opr, bool nomobe = false);

        inline uint32_t getReg32(int reg) {
            return (r[reg] << 16) | r[(reg + 1) & 7];
        }

        inline void setReg32(int reg, uint32_t v) {
            r[reg] = v >> 16;
            r[(reg + 1) & 7] = v;
        }

        inline void setZNCV(bool z, bool n, bool c, bool v) {
            Z = z;
            N = n;
            C = c;
            V = v;
        }

        inline uint16_t getInc(const Operand &opr) {
            uint16_t ret = r[opr.reg];
            r[opr.reg] += opr.diff();
            return ret;
        }

        inline uint16_t getDec(const Operand &opr) {
            r[opr.reg] -= opr.diff();
            return r[opr.reg];
        }

        inline uint8_t get8(const Operand &opr, bool nomove = false) {
            if (opr.mode == 0 && opr.reg != 7) return r[opr.reg];
            int ad = addr(opr, nomove);
            return ad < 0 ? opr.value : read8(ad);
        }

        inline uint16_t get16(const Operand &opr, bool nomove = false) {
            if (opr.mode == 0 && opr.reg != 7) return r[opr.reg];
            int ad = addr(opr, nomove);
            return ad < 0 ? opr.value : read16(ad);
        }

        inline void set8(const Operand &opr, uint8_t value, bool sx = false) {
            if (opr.mode == 0) {
                if (sx) {
                    r[opr.reg] = (int16_t) (int8_t) value;
                } else {
                    r[opr.reg] = (r[opr.reg] & 0xff00) | value;
                }
            } else {
                int ad = addr(opr);
                if (ad >= 0) write8(ad, value);
            }
        }

        inline void set16(const Operand &opr, uint16_t value) {
            if (opr.mode == 0) {
                r[opr.reg] = value;
            } else {
                int ad = addr(opr);
                if (ad >= 0) write16(ad, value);
            }
        }
    };
}
