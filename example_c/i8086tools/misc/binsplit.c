#include <stdio.h>

char buf[32];

int main(int argc, char *argv[]) {
	FILE *r, *w;
	int n = 0, len;
	char fn[16];
	if (argc != 2) return 1;
	r = fopen(argv[1], "rb");
	while ((len = fread(buf, 1, sizeof(buf), r)) > 0) {
		sprintf(fn, "%04d", n++);
		w = fopen(fn, "wb");
		fwrite(buf, 1, len, w);
		fclose(w);
	}
	fclose(r);
	return 0;
}
