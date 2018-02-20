/* This file is in the public domain. */

#include <stdio.h>
#include <string.h>
#include <fcntl.h>

#ifdef WIN32
#  ifdef DEBUG_LD
#    define EBSS 0x421c
static char mem[0xe000];
#  else
#    define EBSS 0
static char mem[0x10000];
#  endif
static char *cur   = mem + EBSS;
static char *start = mem + EBSS;
static char *end   = mem + sizeof(mem);

int brk(char *p) {
#  ifdef DEBUG
	fprintf(stderr, "<brk(0x%04x)", p - mem);
#  endif
	if (p < start || p >= end) {
#  ifdef DEBUG
		fprintf(stderr, " => ENOMEM>\n");
#  endif
		return -1;
	}
	cur = p;
#  ifdef DEBUG
	fprintf(stderr, " => 0>\n");
#  endif
	return 0;
}

char *sbrk(int d) {
	char *old = cur;
	if (!brk(cur + d)) return old;
	return (char *)-1;
}
#endif

int myopen(const char *path, int flag, int mode) {
#ifdef WIN32
	flag |= O_BINARY;
#endif
	int result = open(path, flag, mode);
#ifdef DEBUG
	fprintf(stderr, "<open(%s, %d, %d) => %d>\n", path, flag, mode, result);
#endif
	return result;
}

off_t mylseek(int fd, off_t o, int w) {
	off_t result = lseek(fd, o, w);
#ifdef DEBUG
	fprintf(stderr, "<lseek(%d, %ld, %d) => %ld>\n", fd, o, w, result);
#endif
	return result;
}
