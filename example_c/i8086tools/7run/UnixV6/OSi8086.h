#pragma once
#include "OS.h"
#include "../i8086/VM.h"

namespace UnixV6 {

    class OSi8086 : public OS {
    public:
        static bool check(uint8_t h[2]);

    private:
        i8086::VM cpu;

    public:
        OSi8086();
        OSi8086(const OSi8086 &os);
        virtual ~OSi8086();

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
