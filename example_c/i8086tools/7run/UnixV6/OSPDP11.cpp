#include "OSPDP11.h"
#include "../PDP11/regs.h"
#include "../PDP11/disasm.h"
#include <string.h>
#include <errno.h>
#include <unistd.h>

using namespace UnixV6;
using namespace PDP11;

bool OSPDP11::check(uint8_t h[2]) {
    int magic = ::read16(h);
    return magic == 0407 || magic == 0410 || magic == 0411;
}

OSPDP11::OSPDP11(int ver) : OS(ver) {
    vm = &cpu;
    cpu.unix = this;
}

OSPDP11::OSPDP11(const OSPDP11 &os) : OS(os), cpu(os.cpu) {
    vm = &cpu;
    cpu.unix = this;
}

OSPDP11::~OSPDP11() {
}

void OSPDP11::disasm() {
    int addr = textbase, undef = 0, end = textbase + vm->tsize;
    while (addr < end) {
        vm->showsym(addr);
        OpCode op = disasm1(vm->text, addr);
        std::string ops = cpu.disstr(op);
        int argc = 0;
        if (vm->text[addr + 1] == 0x89) {
            int n = vm->text[addr];
            if (n < nsys && !sysargs[n].name.empty()) {
                argc = sysargs[n].argc;
                if (ver >= 7 && n == 19) {
                    ++argc;
                    ops += " ; lseek";
                } else {
                    ops += " ; " + sysargs[n].name;
                }
            }
        }
        disout(vm->text, addr, op.len, ops);
        if (op.undef()) undef++;
        addr += op.len;
        for (int i = 0; i < argc; i++, addr += 2) {
            ::disout(vm->text, addr, 2, "; arg");
        }
    }
    if (undef) printf("undefined: %d\n", undef);
}

void OSPDP11::setArgs(
        const std::vector<std::string> &args,
        const std::vector<std::string> &) {
    int slen = 0;
    for (int i = 0; i < (int) args.size(); i++) {
        slen += args[i].size() + 1;
    }
    cpu.SP -= (slen + 1) & ~1;
    uint16_t ad1 = cpu.SP;
    cpu.SP -= (1 + args.size() + 1) * 2;
    uint16_t ad2 = cpu.start_sp = cpu.SP;
    vm->write16(cpu.SP, args.size()); // argc
    for (int i = 0; i < (int) args.size(); i++) {
        vm->write16(ad2 += 2, ad1);
        strcpy((char *) vm->data + ad1, args[i].c_str());
        ad1 += args[i].size() + 1;
    }
}

bool OSPDP11::load2(const std::string &fn, FILE *f, size_t size) {
    uint8_t h[0x10];
    if (!fread(h, sizeof (h), 1, f) || !check(h)) {
        return vm->load(fn, f, size);
    }

    vm->release();
    vm->text = new uint8_t[0x10000];
    memset(vm->text, 0, 0x10000);
    vm->tsize = ::read16(h + 2);
    vm->dsize = ::read16(h + 4);
    uint16_t bss = ::read16(h + 6);
    memset(cpu.r, 0, sizeof (cpu.r));
    cpu.PC = ::read16(h + 10);
    cpu.cache.clear();
    cpu.cache.resize(0x10000);
    uint16_t magic = read16(h);
    if (magic == 0411) {
        vm->data = new uint8_t[0x10000];
        memset(vm->data, 0, 0x10000);
        fread(vm->text, 1, vm->tsize, f);
        fread(vm->data, 1, vm->dsize, f);
        vm->brksize = vm->dsize + bss;
    } else if (magic == 0410) {
        vm->data = vm->text;
        fread(vm->text, 1, vm->tsize, f);
        uint16_t doff = (vm->tsize + 0x1fff) & ~0x1fff;
        fread(vm->text + doff, 1, vm->dsize, f);
        vm->brksize = doff + vm->dsize + bss;
    } else { // 0407
        vm->data = vm->text;
        int len = vm->tsize + vm->dsize; // for as
        if (textbase + len > 0x10000) {
            len = 0x10000 - textbase;
        }
        fread(vm->text + textbase, 1, len, f);
        vm->brksize = textbase + len + bss;
        if (textbase) cpu.PC = textbase;
    }

    uint16_t ssize = ::read16(h + 8);
    if (ssize) {
        if (!::read16(h + 14)) {
            fseek(f, vm->tsize + vm->dsize, SEEK_CUR);
        }
        readsym(f, ssize);
    }
    if (read16(vm->text + 2) == 0x1d80) {
        ver = 7;
    }
    return true;
}

bool OSPDP11::syscall(int n) {
    int result, ret;
    if (n == 0) {
        int p = read16(vm->text + cpu.PC);
        int nn = vm->read8(p);
        OS::syscall(&result, nn, cpu.r[0], vm->data + p + 2);
        ret = (nn == 11/*exec*/ || nn == 59/*exece*/) && !result ? 0 : 2;
    } else {
        ret = OS::syscall(&result, n, cpu.r[0], vm->text + cpu.PC);
    }
    if (ret >= 0) {
        cpu.PC += ret;
        cpu.r[0] = (cpu.C = (result == -1)) ? errno : result;
        if (ver >= 7 && n == 19/*lseek*/) {
            cpu.r[1] = result >> 16;
        }
    }
    return true;
}

int OSPDP11::v6_fork() { // 2
#ifdef NO_FORK
    OSPDP11 *ub = new OSPDP11(*this);
    ub->cpu.r[0] = pid;
    forks.push_back(ub);
    int result = ub->pid;
#else
    int result = fork();
    if (result > 0) result = (result % 30000) + 1;
#endif
    if (trace) fprintf(stderr, "<fork() => %d>\n", result);
    return result;
}

int OSPDP11::v6_wait() { // 7
    int status, result = sys_wait(&status);
    cpu.r[1] = status;
    return result;
}

int OSPDP11::v6_exec(const char *path, int argp) { // 11
    if (trace) fprintf(stderr, "<exec(\"%s\"", path);
    std::vector<std::string> args, envs;
    int slen = 0, p;
    while ((p = vm->read16(argp + args.size() * 2))) {
        std::string arg = vm->str(p);
        if (trace && !args.empty()) fprintf(stderr, ", \"%s\"", arg.c_str());
        slen += arg.size() + 1;
        args.push_back(arg);
    }
    if (!load(path)) {
        if (trace) fprintf(stderr, ") => EINVAL>\n");
        errno = EINVAL;
        return -1;
    }
    resetsig();
    setArgs(args, envs);
    if (trace) fprintf(stderr, ") => 0>\n");
    return 0;
}

int OSPDP11::v6_brk(int nd) { // 17
    return sys_brk(nd, cpu.SP);
}

void OSPDP11::sighandler2(int sig) {
    uint16_t r[8];
    memcpy(r, cpu.r, sizeof (r));
    bool Z = cpu.Z, N = cpu.N, C = cpu.C, V = cpu.V;
    cpu.write16((cpu.SP -= 2), cpu.PC);
    cpu.PC = sighandlers[sig];
    while (!cpu.hasExited && !(cpu.PC == PC && cpu.SP == SP)) {
        cpu.run1();
    }
    if (!cpu.hasExited) {
        memcpy(cpu.r, r, sizeof (r));
        cpu.Z = Z;
        cpu.N = N;
        cpu.C = C;
        cpu.V = V;
    }
}
