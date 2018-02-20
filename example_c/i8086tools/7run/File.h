#pragma once
#include <stdint.h>
#include <string>
#include <sys/types.h>

struct FileBase {
    int fd;
    int count;
    std::string path;

    FileBase(int fd, const std::string &path);
    virtual ~FileBase();

    virtual int read(void *buf, int len) = 0;
    virtual int write(void *buf, int len) = 0;
    virtual off_t lseek(off_t o, int w) = 0;
    virtual FileBase *dup() = 0;
};

struct File : public FileBase {
    File(int fd, const std::string &path);
    File(const std::string &path, int flag, int mode);
    virtual ~File();

    virtual int read(void *buf, int len);
    virtual int write(void *buf, int len);
    virtual off_t lseek(off_t o, int w);
    virtual FileBase *dup();
};
