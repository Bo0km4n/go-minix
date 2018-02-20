#pragma once
#include "OS.h"
#include "../PDP11/VM.h"

namespace UnixV6 {

    class OSPDP11 : public OS {
    public:
        static bool check(uint8_t h[2]);

    private:
        PDP11::VM cpu;

    public:
        OSPDP11(int ver);
        OSPDP11(const OSPDP11 &os);
        virtual ~OSPDP11();

        virtual void disasm();
        virtual bool syscall(int n);

    protected:
        virtual void setArgs(
                const std::vector<std::string> &args,
                const std::vector<std::string> &envs);
        virtual bool load2(const std::string &fn, FILE *f, size_t size);

    public:
        virtual int v6_fork(); //  2
        virtual int v6_wait(); // 7
        virtual int v6_exec(const char *path, int argp); // 11
        virtual int v6_brk(int nd); // 17

    protected:
        virtual void sighandler2(int sig);
    };
}
