#pragma once
#include "../utils.h"
#include <string>

namespace i8086 {

    enum OperandType {
        Reg, SReg, Imm, Addr, Far, Ptr, ModRM
    };

    enum Direction {
        RmReg, RegRm
    };

    struct VM;

    struct Operand {
        VM *vm;
        int len;
        bool w;
        int type, value, addr, seg;

        inline bool empty() const {
            return len < 0;
        }

        std::string str() const;

        inline Operand(int len, bool w, int type, int value, int seg = -1) {
            this->len = len;
            this->w = w;
            this->type = type;
            this->value = value;
            this->seg = seg;
        }

        inline Operand(VM *vm) {
            this->vm = vm;
        }

        void set(int type, bool w, int v);
        uint8_t * ptr() const;
        int u() const;
        int setf(int val);
        void operator =(int val);
        
        size_t modrm(uint8_t *p, bool w);
        size_t regrm(Operand *opr, uint8_t *p, bool dir, bool w);
        size_t aimm(Operand *opr, bool w, uint8_t *p);
        size_t getopr(Operand *opr, uint8_t b, uint8_t *p);

        inline int operator *() const {
            int ret = u();
            return w ? int16_t(ret) : int8_t(ret);
        }

        inline bool operator>(int val) {
            return u() > (w ? uint16_t(val) : uint8_t(val));
        }

        inline bool operator<(int val) {
            return u() < (w ? uint16_t(val) : uint8_t(val));
        }

    };

    extern Operand noopr, dx, cl, es, cs, ss, ds;
}
