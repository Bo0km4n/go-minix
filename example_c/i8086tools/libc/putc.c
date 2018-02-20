int putc(char ch) {
	write(1, &ch, 1);
	return ch;
}
