#include "OS.h"
#include <string.h>
#include <signal.h>

#define V6_SIGINT   2
#define V6_SIGINS   4
#define V6_SIGFPT   8
#define V6_SIGSEG  11

using namespace UnixV6;

OS::OS(int ver) {
    memset(sighandlers, 0, sizeof (sighandlers));
    textbase = ver <= 2 ? 0x4000 : 0;
    this->ver = ver;
}

OS::OS(const OS &os) : UnixBase(os) {
    memset(sighandlers, 0, sizeof (sighandlers));
    textbase = os.textbase;
    ver = os.ver;
}

OS::~OS() {
}

void OS::readsym(FILE *f, int ssize) {
    if (!ssize) return;

    uint8_t buf[12];
    for (int i = 0; i < ssize; i += 12) {
        fread(buf, sizeof (buf), 1, f);
        Symbol sym = {
            readstr(buf, 8), ::read16(buf + 8), ::read16(buf + 10) + textbase
        };
        int t = sym.type;
        if (t < 6) {
            t = "uatdbc"[t];
        } else if (0x20 <= t && t < 0x26) {
            t = "UATDBC"[t - 0x20];
        }
        switch (t) {
            case 't':
            case 'T':
                if (!sym.name.empty() && !startsWith(sym.name, "~")) {
                    vm->syms[1][sym.addr] = sym;
                }
                break;
            case 0x1f:
                vm->syms[0][sym.addr] = sym;
                break;
        }
    }
}

void OS::setstat(uint16_t addr, struct stat * st) {
    if (ver >= 7) {
        memset(vm->data + addr, 0, 30);
        vm->write16(addr, st->st_dev);
        vm->write16(addr + 2, st->st_ino);
        vm->write16(addr + 4, st->st_mode);
        vm->write16(addr + 6, st->st_nlink);
        vm->write16(addr + 8, st->st_uid);
        vm->write16(addr + 10, st->st_gid);
        vm->write16(addr + 12, st->st_rdev);
        vm->write32pdp(addr + 14, st->st_size);
        vm->write32pdp(addr + 18, st->st_atime);
        vm->write32pdp(addr + 22, st->st_mtime);
        vm->write32pdp(addr + 26, st->st_ctime);
    } else {
        memset(vm->data + addr, 0, 36);
        vm->write16(addr, st->st_dev);
        vm->write16(addr + 2, st->st_ino);
        vm->write16(addr + 4, st->st_mode);
        vm->write8(addr + 6, st->st_nlink);
        vm->write8(addr + 7, st->st_uid);
        vm->write8(addr + 8, st->st_gid);
        vm->write8(addr + 9, st->st_size >> 16);
        vm->write16(addr + 10, st->st_size);
        vm->write32pdp(addr + 28, st->st_atime);
        vm->write32pdp(addr + 32, st->st_mtime);
    }
}

void OS::sighandler(int sig) {
    OS *cur = dynamic_cast<OS *> (current);
    if (cur) cur->sighandler2(sig);
}

int OS::convsig(int sig) {
    switch (sig) {
        case V6_SIGINT: return SIGINT;
        case V6_SIGINS: return SIGILL;
        case V6_SIGFPT: return SIGFPE;
        case V6_SIGSEG: return SIGSEGV;
    }
    return -1;
}

void OS::setsig(int sig, int h) {
    if (h == 0) {
        signal(sig, SIG_DFL);
    } else if (h & 1) {
        signal(sig, SIG_IGN);
    } else {
        signal(sig, &sighandler);
    }
}

void OS::resetsig() {
    for (int i = 0; i < nsig; i++) {
        uint16_t &sgh = sighandlers[i];
        if (sgh && !(sgh & 1)) {
            sighandlers[i] = 0;
            int s = convsig(i);
            if (s >= 0) signal(s, SIG_DFL);
        }
    }
}

void OS::swtch(bool reset) {
    for (int i = 0; i < nsig; i++) {
        int s = convsig(i);
        if (s >= 0) {
            if (reset) {
                signal(s, SIG_DFL);
            } else {
                setsig(s, sighandlers[i]);
            }
        }
    }
}

void OS::coredump(const char *path) {
    FILE *f = fopen(path, "wb");
    fwrite(vm->data, 1, 0x10000, f);
    fclose(f);
}
