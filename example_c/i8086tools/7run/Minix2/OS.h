#pragma once
#include "../UnixBase.h"
#include "../i8086/VM.h"

namespace Minix2 {

    class OS : public UnixBase {
    private:
        i8086::VM cpu;

    public:
        OS();
        OS(const OS &os);
        virtual ~OS();

        virtual void disasm();
        virtual bool syscall(int n);

    protected:
        virtual void setArgs(
                const std::vector<std::string> &args,
                const std::vector<std::string> &envs);
        virtual bool load2(const std::string &fn, FILE *f, size_t size);
        virtual void setsig(int sig, int h);
        virtual void setstat(uint16_t addr, struct stat *st);
        virtual void swtch(bool reset = false);

    private:
        bool syscall(int *result, uint8_t *m);
        int minix_fork(); //  2
        int minix_brk(int nd); // 17
        int minix_signal(int sig, int h); // 48
        int minix_exec(const char *path, int frame, int fsize); // 59
        int minix_sigaction(int sig, int act, int oact); // 71

        static void sighandler(int sig);
        void sighandler2(int sig);
        virtual int convsig(int sig);
        void resetsig();

        struct sigact {
            uint16_t handler;
            uint16_t mask;
            uint16_t flags;
        };
        static const int nsig = 12;
        sigact sigacts[nsig];
    };
}
