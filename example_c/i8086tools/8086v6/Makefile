include ../Makefile.inc
TARGET = nm strip cc

all: $(TARGET)

nm: nm.c
	v6cc $<
	mv a.out $@
	v6strip $@

strip: strip.s
	v6as $<
	mv a.out $@
	v6strip $@

cc: cc.c
	m2cc -s -o $@ $<

install: $(TARGET)
	mkdir -p $(V6BIN) $(ACKBIN)
	install -c nm strip $(V6BIN)
	install -c cc $(ACKBIN)

clean:
	rm -f $(TARGET)
