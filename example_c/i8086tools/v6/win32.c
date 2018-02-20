#ifdef _WIN32
#include <windows.h>
#include <stdio.h>
#include <string.h>
#include <fcntl.h>
#include <time.h>
#include <errno.h>
#include <limits.h>
#include <sys/stat.h>

static char path[PATH_MAX];

const char *convtemp(const char src[PATH_MAX])
{
    if (!strncmp(src, "/tmp/", 5)) {
        GetTempPath(sizeof(path), path);
        strncat(path, &src[5], PATH_MAX);
        path[sizeof(path) - 1] = 0;
    } else {
        strcpy(path, src);
    }
    return path;
}

int mkstemp(char fn[PATH_MAX])
{
    int ret, len;
    convtemp(fn);
    len = strlen(path);
    if (len < 6 || strcmp(&path[len - 6], "XXXXXX")) return -1;
    srand(time(NULL));
    do {
        int i;
        for (i = len - 6; i < len; ++i) {
            path[i] = '0' + ((rand() >> 3) % 10);
        }
        ret = open(path, O_CREAT | O_EXCL | O_BINARY | O_WRONLY, S_IWRITE);
    } while (ret == -1 && errno == EEXIST);
    return ret;
}

void link(char *src, char *dst)
{
    FILE *s, *d;
    int len;
    char buf[512];
    s = fopen(src, "rb");
    d = fopen(dst, "wb");
    while ((len = fread(buf, 1, sizeof(buf), s)) > 0) {
        fwrite(buf, 1, len, d);
    }
    fclose(d);
    fclose(s);
}

static char buf[256 * 1024];
static char *ps = &buf[0], *pe = &buf[sizeof(buf)];

void *sbrk(int size)
{
    char *ret = ps;
    if (ps + size > pe) return (void *)-1;
    ps += size;
    return ret;
}
#endif
