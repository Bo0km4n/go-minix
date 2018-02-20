int puts(const char *s) {
	write(1, s, strlen(s));
	putc('\n');
	return 0;
}
