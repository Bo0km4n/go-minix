#include "OS.h"
#include <errno.h>
#include <unistd.h>

using namespace UnixV6;

sysarg OS::sysargs[] = {
    {/* 0*/ 1, "indir"},
    {/* 1*/ 0, "exit"},
    {/* 2*/ 0, "fork"},
    {/* 3*/ 2, "read"},
    {/* 4*/ 2, "write"},
    {/* 5*/ 2, "open"},
    {/* 6*/ 0, "close"},
    {/* 7*/ 0, "wait"},
    {/* 8*/ 2, "creat"},
    {/* 9*/ 2, "link"},
    {/*10*/ 1, "unlink"},
    {/*11*/ 2, "exec"},
    {/*12*/ 1, "chdir"},
    {/*13*/ 0, "time"},
    {/*14*/ 3, "mknod"},
    {/*15*/ 2, "chmod"},
    {/*16*/ 2, "chown"},
    {/*17*/ 1, "brk"},
    {/*18*/ 2, "stat"},
    {/*19*/ 2, "seek"}, // lseek for UNIX V7
    {/*20*/ 0, "getpid"},
    {/*21*/ 3, "mount"},
    {/*22*/ 1, "umount"},
    {/*23*/ 0, "setuid"},
    {/*24*/ 0, "getuid"},
    {/*25*/ 0, "stime"},
    {/*26*/ 3, "ptrace"},
    {/*27*/ 1, "intr"}, // for UNIX V2
    {/*28*/ 1, "fstat"},
    {/*29*/ 0, ""},
    {/*30*/ 1, "smdate"},
    {/*31*/ 1, "stty"},
    {/*32*/ 1, "gtty"},
    {/*33*/ 2, "access"}, // for UNIX V7
    {/*34*/ 0, "nice"},
    {/*35*/ 0, "sleep"},
    {/*36*/ 0, "sync"},
    {/*37*/ 1, "kill"},
    {/*38*/ 0, "switch"},
    {/*39*/ 0, ""},
    {/*40*/ 0, ""},
    {/*41*/ 0, "dup"},
    {/*42*/ 0, "pipe"},
    {/*43*/ 1, "times"},
    {/*44*/ 4, "prof"},
    {/*45*/ 0, "tiu"},
    {/*46*/ 0, "setgid"},
    {/*47*/ 0, "getgid"},
    {/*48*/ 2, "signal"},
    {/*49*/ 0, ""},
    {/*50*/ 0, ""},
    {/*51*/ 0, ""},
    {/*52*/ 0, ""},
    {/*53*/ 0, ""},
    {/*54*/ 3, "ioctl"}, // for UNIX V7
    {/*55*/ 0, ""},
    {/*56*/ 0, ""},
    {/*57*/ 0, ""},
    {/*58*/ 0, ""},
    {/*59*/ 3, "exece"}, // for UNIX V7
    {/*60*/ 1, "umask"}, // for UNIX V7
};

int OS::syscall(int *result, int n, int arg0, uint8_t *args) {
    *result = 0;
    switch (n) {
        case 1:
            sys_exit((int16_t) arg0);
            return -1;
        case 2:
        {
#ifdef NO_FORK
            *result = v6_fork();
#else
            int pid = sys_getpid();
            *result = v6_fork();
            if (*result < 0) {
                return 0;
            } else if (*result == 0) {
                *result = pid;
                return 0;
            }
#endif
            return 2;
        }
        case 3:
            *result = sys_read(arg0, read16(args), read16(args + 2));
            return 4;
        case 4:
            *result = sys_write(arg0, read16(args), read16(args + 2));
            return 4;
        case 5:
            *result = sys_open(vm->str16(args), read16(args + 2));
            return 4;
        case 6:
            *result = sys_close(arg0);
            return 0;
        case 7:
            *result = v6_wait();
            return 0;
        case 8:
            *result = sys_creat(vm->str16(args), convmode(read16(args + 2)));
            return 4;
        case 9:
            *result = sys_link(vm->str16(args), vm->str16(args + 2));
            return 4;
        case 10:
            *result = sys_unlink(vm->str16(args));
            return 2;
        case 11:
            //coredump("core");
            *result = v6_exec(vm->str16(args), read16(args + 2));
            //sys_exit(0);
            return *result ? 4 : 0;
        case 12:
            *result = sys_chdir(vm->str16(args));
            return 2;
        case 15:
            *result = sys_chmod(vm->str16(args), convmode(read16(args + 2)));
            return 4;
        case 17:
            *result = v6_brk(read16(args));
            return 2;
        case 18:
            *result = sys_stat(vm->str16(args), read16(args + 2));
            return 4;
        case 19:
            if (ver < 7) {
                *result = v6_seek(arg0, read16(args), read16(args + 2));
                return 4;
            } else {
                off_t o = read32pdp(args);
                *result = sys_lseek(arg0, o, read16(args + 4));
                return 6;
            }
        case 20:
            *result = sys_getpid();
            return 0;
        case 27: // intr (V2)
            *result = v6_signal(2, read16(args));
            return 2;
        case 28:
            *result = sys_fstat(arg0, read16(args));
            return 2;
        case 33:
            *result = sys_access(vm->str16(args), read16(args + 6));
            return 4;
        case 41:
            *result = sys_dup(arg0);
            return 0;
        case 48:
            *result = v6_signal(read16(args), read16(args + 2));
            return 4;
        case 54:
            *result = sys_ioctl(read16(args), read16(args + 2), read16(args + 4));
            return 6;
        case 59:
            //coredump("core");
            *result = v6_exec(vm->str16(args), read16(args + 2));
            //sys_exit(0);
            return *result ? 6 : 0;
        case 60:
            *result = sys_umask(read16(args));
            return 2;
        default:
            if (n < nsys && !sysargs[n].name.empty()) {
                fprintf(stderr, "<%s: not implemented>\n", sysargs[n].name.c_str());
            } else {
                fprintf(stderr, "<%d: unknown syscall>\n", n);
            }
            break;
    }
    sys_exit(-1);
    return -1;
}

int OS::v6_seek(int fd, off_t o, int w) { // 19
    if (trace) fprintf(stderr, "<seek(%d, %ld, %d)", fd, long(o), w);
    FileBase *f = file(fd);
    off_t result = -1;
    switch (w) {
        case 0:
            result = f->lseek(o, SEEK_SET);
            break;
        case 1:
            result = f->lseek(int16_t(uint16_t(o)), SEEK_CUR);
            break;
        case 2:
            result = f->lseek(int16_t(uint16_t(o)), SEEK_END);
            break;
        case 3:
            result = f->lseek(o * 512, SEEK_SET);
            break;
        case 4:
            result = f->lseek(int(int16_t(uint16_t(o))) * 512, SEEK_CUR);
            break;
        case 5:
            result = f->lseek(int(int16_t(uint16_t(o))) * 512, SEEK_END);
            break;
        default:
            errno = EINVAL;
            break;
    }
    if (trace) fprintf(stderr, " => %ld>\n", long(result));
    return result;
}

int OS::v6_signal(int sig, int h) {
    if (trace) fprintf(stderr, "<signal(%d, 0x%04x)>\n", sig, h);
    int s = convsig(sig);
    if (s < 0) {
        errno = EINVAL;
        return -1;
    }
    int oh = sighandlers[sig];
    sighandlers[sig] = h;
    setsig(s, h);
    return oh;
}

int OS::convmode(int mode) {
    if (!textbase) return mode;
    return mode == 037 ? 0755 : 0644;
}
