include ../Makefile.inc
TARGET = pe2aout aout-tool strip-cr
CFLAGS = -g
CROSSLDFLAGS = -s -Ttext=0 --image-base=0 --section-alignment=0 --file-alignment=0

all: $(TARGET)

pe2aout: pe2aout.c pe.h a.out.h
aout-tool: aout-tool.c a.out.h
strip-cr: strip-cr.c

$(TARGET):
	$(CC) $(CFLAGS) -o $@ $<

test: exit-32.out write-16.out aout-tool
	./aout-tool *.out

exit-32.out: exit-32.exe pe2aout
	./pe2aout $<

write-16.out: write-16.exe pe2aout
	./pe2aout -16 $<

exit-32.exe: exit-32.o
write-16.exe: write-16.o

exit-32.exe write-16.exe:
	$(MINGW)ld -o $@ $(CROSSLDFLAGS) $<

exit-32.o: exit-32.c
	$(MINGW)gcc -c -o $@ -masm=intel $<

write-16.o: write-16.s
	$(MINGW)as -o $@ $<

install: $(TARGET)
	install -cs $(TARGET) $(PREFIX)/bin
	install -c extract-text.py $(PREFIX)/bin/extract-text

uninstall:
	cd $(PREFIX)/bin && rm -f $(TARGET) extract-text

clean:
	rm -f $(TARGET) *.exe *.o *.out
