#pragma once
#include "utils.h"
#include "File.h"
#include "VMBase.h"
#include <stdio.h>
#include <sys/stat.h>
#include <vector>
#include <list>
#ifdef WIN32
#define NO_FORK
#endif

class UnixBase {
protected:
    static UnixBase *current;
#ifdef NO_FORK
    static std::list<UnixBase *> forks;
#endif
    VMBase *vm;
    int exitcode, pid;
    uint16_t umask;
    std::vector<FileBase *> files;

public:
    UnixBase();
    UnixBase(const UnixBase &os);
    virtual ~UnixBase();

    virtual void disasm() = 0;
    virtual bool syscall(int n) = 0;
    bool load(const std::string &fn);
    int run(
            const std::vector<std::string> &args,
            const std::vector<std::string> &envs);
    int run();
    void swtch(UnixBase *to);

protected:
    virtual bool load2(const std::string &fn, FILE *f, size_t size) = 0;
    virtual void setArgs(
            const std::vector<std::string> &args,
            const std::vector<std::string> &envs) = 0;
    virtual void setstat(uint16_t addr, struct stat *st) = 0;
    virtual int convsig(int sig) = 0;
    virtual void setsig(int sig, int h) = 0;
    virtual void swtch(bool reset = false) = 0;

public:
    int getfd();
    int open(const std::string &path, int flag, int mode);
    int close(int fd);
    int dup(int fd);
    FileBase *file(int fd);

    void sys_exit(int code); // 1
    //int sys_fork(); // 2
    int sys_read(int fd, int buf, int len); // 3
    int sys_write(int fd, int buf, int len); // 4
    int sys_open(const char *path, int flag, mode_t mode = 0); // 5
    int sys_close(int fd); // 6
    int sys_wait(int *status); // 7
    int sys_creat(const char *path, mode_t mode); // 8
    int sys_link(const char *src, const char *dst); // 9
    int sys_unlink(const char *path); // 10
    int sys_chdir(const char *path); // 12
    int sys_time(); // 13
    int sys_chmod(const char *path, mode_t mode); // 15
    int sys_brk(int nd, int sp); // 17
    int sys_stat(const char *path, int p); // 18
    off_t sys_lseek(int fd, off_t o, int w); // 19
    int sys_getpid(); // 20
    int sys_getuid(); // 24
    int sys_fstat(int fd, int p); // 28
    int sys_access(const char *path, mode_t mode); // 33
    int sys_dup(int fd); // 41
    int sys_getgid(); // 47
    //void sys_signal(); // 48
    int sys_ioctl(int fd, int rq, int d); // 54
    //int sys_exec(const char *path, int frame, int fsize); // 59
    int sys_umask(mode_t mask); // 60
    //void sys_sigaction(); // 71
};
