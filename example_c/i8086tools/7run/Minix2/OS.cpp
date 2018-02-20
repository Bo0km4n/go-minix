#include "OS.h"
#include "../i8086/regs.h"
#include <string.h>

using namespace Minix2;

OS::OS() {
    vm = &cpu;
    cpu.unix = this;
    memset(sigacts, 0, sizeof (sigacts));
}

OS::OS(const OS &os) : UnixBase(os), cpu(os.cpu) {
    vm = &cpu;
    cpu.unix = this;
    memcpy(sigacts, os.sigacts, sizeof (sigacts));
}

OS::~OS() {
}

void OS::disasm() {
    vm->disasm();
}

void OS::setArgs(
        const std::vector<std::string> &args,
        const std::vector<std::string> &envs) {
    int slen = 0;
    for (int i = 0; i < (int) args.size(); i++) {
        slen += args[i].size() + 1;
    }
    for (int i = 0; i < (int) envs.size(); i++) {
        slen += envs[i].size() + 1;
    }
    cpu.SP -= (slen + 1) & ~1;
    uint16_t ad1 = cpu.SP;
    cpu.SP -= (1 + args.size() + 1 + envs.size() + 1) * 2;
    uint16_t ad2 = cpu.start_sp = cpu.SP;
    vm->write16(cpu.SP, args.size()); // argc
    for (int i = 0; i < (int) args.size(); i++) {
        vm->write16(ad2 += 2, ad1);
        strcpy((char *) vm->data + ad1, args[i].c_str());
        ad1 += args[i].size() + 1;
    }
    vm->write16(ad2 += 2, 0); // argv[argc]
    for (int i = 0; i < (int) envs.size(); i++) {
        vm->write16(ad2 += 2, ad1);
        strcpy((char *) vm->data + ad1, envs[i].c_str());
        ad1 += envs[i].size() + 1;
    }
    vm->write16(ad2 += 2, 0); // envp (last)
}

bool OS::load2(const std::string &fn, FILE *f, size_t size) {
    uint8_t h[0x20];
    if (!fread(h, sizeof (h), 1, f) || !(h[0] == 1 && h[1] == 3)) {
        return vm->load(fn, f, size);
    }
    if (h[3] != 4) {
        fprintf(stderr, "unknown cpu id: %d\n", h[3]);
        return false;
    }
    vm->release();
    vm->text = new uint8_t[0x10000];
    memset(vm->text, 0, 0x10000);
    fseek(f, h[4], SEEK_SET);
    vm->tsize = ::read32(h + 8);
    vm->dsize = ::read32(h + 12);
    uint16_t bss = ::read32(h + 16);
    cpu.IP = ::read32(h + 20);
    if (h[2] & 0x20) {
        vm->data = new uint8_t[0x10000];
        memset(vm->data, 0, 0x10000);
        fread(vm->text, 1, vm->tsize, f);
        fread(vm->data, 1, vm->dsize, f);
        vm->brksize = vm->dsize + bss;
    } else {
        vm->data = vm->text;
        fread(vm->text, 1, vm->tsize + vm->dsize, f);
        vm->brksize = vm->tsize + vm->dsize + bss;
    }
    return true;
}

void OS::setstat(uint16_t addr, struct stat *st) {
    vm->write16(addr, st->st_dev);
    vm->write16(addr + 2, st->st_ino);
    vm->write16(addr + 4, st->st_mode);
    vm->write16(addr + 6, st->st_nlink);
    vm->write16(addr + 8, st->st_uid);
    vm->write16(addr + 10, st->st_gid);
    vm->write16(addr + 12, st->st_rdev);
    vm->write32(addr + 14, st->st_size);
    vm->write32(addr + 18, st->st_atime);
    vm->write32(addr + 22, st->st_mtime);
    vm->write32(addr + 26, st->st_ctime);
}
