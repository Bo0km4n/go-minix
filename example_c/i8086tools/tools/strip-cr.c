#include <stdio.h>

int main(int argc, char *argv[]) {
	int ch;
	FILE *r, *w;
	if (argc != 3) {
		fprintf(stderr, "usage: %s in out\n", argv[0]);
		return 1;
	}
	r = fopen(argv[1], "rb");
	if (!r) {
		fprintf(stderr, "can not open: %s\n", argv[1]);
		return 1;
	}
	w = fopen(argv[2], "wb");
	if (!w) {
		fprintf(stderr, "can not open: %s\n", argv[2]);
		fclose(r);
		return 1;
	}
	while ((ch = fgetc(r)) != -1) {
		if (ch != '\r') fputc(ch, w);
	}
	fclose(w);
	fclose(r);
	return 0;
}
