#include "File.h"
#include <fcntl.h>
#include <unistd.h>

FileBase::FileBase(int fd, const std::string &path) : fd(fd), count(1), path(path) {
}

FileBase::~FileBase() {
}

File::File(int fd, const std::string &path) : FileBase(fd, path) {
}

File::File(const std::string &path, int flag, int mode)
: FileBase(open(path.c_str(), flag, mode), path) {
}

File::~File() {
    if (fd > 2) close(fd);
}

int File::read(void *buf, int len) {
    return ::read(fd, buf, len);
}

int File::write(void *buf, int len) {
    return ::write(fd, buf, len);
}

off_t File::lseek(off_t o, int w) {
    return ::lseek(fd, o, w);
}

FileBase *File::dup() {
    return new File(::dup(fd), path);
}
