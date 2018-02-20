static int print10(int v) {
	char buf[16], *end = buf + sizeof(buf), *p = end;
	int len, minus = v < 0;
	if (minus) v = -v; else if (v == 0) *(--p) = '0';
	while (v) {
		*(--p) = '0' + (v % 10);
		v /= 10;
	}
	if (minus) *(--p) = '-';
	len = end - p;
	write(1, p, len);
	return len;
}

int printf(const char *fmt, ...) {
	int ret = 0, *p = &fmt + 1, len;
	char ch;
	while (ch = *(fmt++)) {
		if (ch != '%') {
			putc(ch);
			++ret;
			continue;
		}
		switch (ch = *(fmt++)) {
		case 'd':
			ret += print10(*(p++));
			break;
		case 's':
			len = strlen(*p);
			write(1, *(p++), len);
			ret += len;
			break;
		case '%':
			putc('%');
			++ret;
			break;
		default:
			write(1, "???", 3);
			if (!ch) --fmt;
			break;
		}
	}
	return ret;
}
