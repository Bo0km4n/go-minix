#include "OS.h"
#include "../i8086/regs.h"
#include <string.h>
#include <errno.h>
#include <unistd.h>

using namespace Minix2;

bool OS::syscall(int n) {
    if (n != 0x20) return false;
    int result;
    if (syscall(&result, vm->data + cpu.BX)) {
        vm->write16(cpu.BX + 2, result == -1 ? -errno : result);
        cpu.AX = 0;
    }
    return true;
}

bool OS::syscall(int *result, uint8_t *m) {
    *result = 0;
    int n = read16(m + 2);
    switch (n) {
        case 1:
            sys_exit((int16_t) read16(m + 4));
            return false;
        case 2:
            *result = minix_fork();
            return true;
        case 3:
            *result = sys_read(read16(m + 4), read16(m + 10), read16(m + 6));
            return true;
        case 4:
            *result = sys_write(read16(m + 4), read16(m + 10), read16(m + 6));
            return true;
        case 5:
        {
            int flag = read16(m + 6);
            if (flag & 64 /*O_CREAT*/) {
                *result = sys_open(vm->str16(m + 10), flag, read16(m + 8));
            } else {
                *result = sys_open(vm->str16(m + 8), flag);
            }
            return true;
        }
        case 6:
            *result = sys_close(read16(m + 4));
            return true;
        case 7:
        {
            int status;
            *result = sys_wait(&status);
            write16(m + 4, status);
            return true;
        }
        case 8:
            *result = sys_creat(vm->str16(m + 8), read16(m + 6));
            return true;
        case 9:
            *result = sys_link(vm->str16(m + 10), vm->str16(m + 12));
            return true;
        case 10:
            *result = sys_unlink(vm->str16(m + 8));
            return true;
        case 11:
            fprintf(stderr, "<waitpid: not implemented>\n");
            break;
        case 12:
            *result = sys_chdir(vm->str16(m + 8));
            return true;
        case 13:
            *result = sys_time();
            if (*result >= 0) {
                write32(m + 10, *result);
                *result = 0;
            }
            return true;
        case 14:
            fprintf(stderr, "<mknod: not implemented>\n");
            break;
        case 15:
            *result = sys_chmod(vm->str16(m + 8), read16(m + 6));
            return true;
        case 16:
            fprintf(stderr, "<chown: not implemented>\n");
            break;
        case 17:
            *result = minix_brk(read16(m + 10));
            if (!*result) write16(m + 18, vm->brksize);
            return true;
        case 18:
            *result = sys_stat(vm->str16(m + 10), read16(m + 12));
            return true;
        case 19:
        {
            off_t o = sys_lseek(read16(m + 4), read32(m + 10), read16(m + 6));
            if (o == -1) {
                *result = -1;
            } else {
                write32(m + 10, o);
                *result = 0;
            }
            return true;
        }
        case 20:
            *result = sys_getpid();
            return true;
        case 21:
            fprintf(stderr, "<mount: not implemented>\n");
            break;
        case 22:
            fprintf(stderr, "<umount: not implemented>\n");
            break;
        case 23:
            fprintf(stderr, "<setuid: not implemented>\n");
            break;
        case 24:
            *result = sys_getuid();
            return true;
        case 25:
            fprintf(stderr, "<stime: not implemented>\n");
            break;
        case 26:
            fprintf(stderr, "<ptrace: not implemented>\n");
            break;
        case 27:
            fprintf(stderr, "<alarm: not implemented>\n");
            break;
        case 28:
            *result = sys_fstat(read16(m + 4), read16(m + 10));
            return true;
        case 29:
            fprintf(stderr, "<pause: not implemented>\n");
            break;
        case 30:
            fprintf(stderr, "<utime: not implemented>\n");
            break;
        case 33:
            *result = sys_access(vm->str16(m + 8), read16(m + 6));
            return true;
        case 36:
            fprintf(stderr, "<sync: not implemented>\n");
            break;
        case 37:
            fprintf(stderr, "<kill: not implemented>\n");
            break;
        case 38:
            fprintf(stderr, "<rename: not implemented>\n");
            break;
        case 39:
            fprintf(stderr, "<mkdir: not implemented>\n");
            break;
        case 40:
            fprintf(stderr, "<rmdir: not implemented>\n");
            break;
        case 41:
            fprintf(stderr, "<dup: not implemented>\n");
            break;
        case 42:
            fprintf(stderr, "<pipe: not implemented>\n");
            break;
        case 43:
            fprintf(stderr, "<times: not implemented>\n");
            break;
        case 46:
            fprintf(stderr, "<setgid: not implemented>\n");
            break;
        case 47:
            *result = sys_getgid();
            return true;
        case 48:
            *result = minix_signal(read16(m + 4), read16(m + 14));
            return true;
        case 54:
            *result = sys_ioctl(read16(m + 4), read16(m + 8), read16(m + 18));
            return true;
        case 55:
            fprintf(stderr, "<fcntl: not implemented>\n");
            break;
        case 59:
            *result = minix_exec(vm->str16(m + 10), read16(m + 12), read16(m + 6));
            return *result;
        case 60:
            *result = sys_umask(read16(m + 4));
            return true;
        case 61:
            fprintf(stderr, "<chroot: not implemented>\n");
            break;
        case 62:
            fprintf(stderr, "<setsid: not implemented>\n");
            break;
        case 63:
            fprintf(stderr, "<getpgrp: not implemented>\n");
            break;
        case 64:
            fprintf(stderr, "<ksig: not implemented>\n");
            break;
        case 65:
            fprintf(stderr, "<unpause: not implemented>\n");
            break;
        case 67:
            fprintf(stderr, "<revive: not implemented>\n");
            break;
        case 68:
            fprintf(stderr, "<task_reply: not implemented>\n");
            break;
        case 71:
            *result = minix_sigaction(read16(m + 6), read16(m + 10), read16(m + 12));
            return true;
        case 72:
            fprintf(stderr, "<sigsuspend: not implemented>\n");
            break;
        case 73:
            fprintf(stderr, "<sigpending: not implemented>\n");
            break;
        case 74:
            fprintf(stderr, "<sigprocmask: not implemented>\n");
            break;
        case 75:
            fprintf(stderr, "<sigreturn: not implemented>\n");
            break;
        case 76:
            fprintf(stderr, "<reboot: not implemented>\n");
            break;
        case 77:
            fprintf(stderr, "<svrctl: not implemented>\n");
            break;
        default:
            fprintf(stderr, "<%d: unknown syscall>\n", n);
            break;
    }
    sys_exit(-1);
    return false;
}

int OS::minix_fork() { // 2
#ifdef NO_FORK
    OS *ub = new OS(*this);
    ub->cpu.write16(cpu.BX + 2, 0);
    ub->cpu.AX = 0;
    forks.push_back(ub);
    int result = ub->pid;
#else
    int result = fork();
    if (result > 0) result = (result % 30000) + 1;
#endif
    if (trace) fprintf(stderr, "<fork() => %d>\n", result);
    return result;
}

int OS::minix_brk(int nd) { // 17
    return sys_brk(nd, cpu.SP);
}

int OS::minix_exec(const char *path, int frame, int fsize) { // 59
#if 0
    FILE *f = fopen("core", "wb");
    fwrite(data, 1, 0x10000, f);
    fclose(f);
#endif
    std::vector<uint8_t> f(fsize);
    memcpy(&f[0], vm->data + frame, fsize);
    int argc = read16(&f[0]);
    if (trace) {
        fprintf(stderr, "<exec(\"%s\"", path);
        for (int i = 2; i <= argc; i++) {
            fprintf(stderr, ", \"%s\"", &f[read16(&f[i * 2])]);
        }
        fprintf(stderr, ")>\n");
    }
    if (!load(path)) {
        errno = EINVAL;
        return -1;
    }
    resetsig();
    cpu.start_sp = cpu.SP = 0x10000 - fsize;
    memcpy(vm->data + cpu.start_sp, &f[0], fsize);
    int ad = cpu.start_sp + 2, p;
    for (int i = 0; i < 2; i++, ad += 2) {
        for (; (p = vm->read16(ad)); ad += 2) {
            vm->write16(ad, cpu.start_sp + p);
        }
    }
    return 0;
}
