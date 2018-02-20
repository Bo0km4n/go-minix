include ../Makefile.inc
TARGET = crt0.o libc.a
CC = 8086v6-cc
AS = m2as
AR = m2aal

all: $(TARGET)

crt0.o: crt0.s
	$(CC) -c $<

libc.a: cret.o write.o exit.o strlen.o puts.o putc.o printf.o
	$(AR) r $@ $^

install: $(TARGET)
	install -c $(TARGET) $(ACKLIB)

clean:
	rm -f *.o $(TARGET)
