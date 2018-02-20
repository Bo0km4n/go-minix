TARGET = libobject.a
SOURCES = rd.c rd_arhdr.c rd_bytes.c rd_int2.c rd_long.c rd_ranlib.c rd_unsig2.c wr.c wr_arhdr.c wr_bytes.c wr_int2.c wr_long.c wr_putc.c wr_ranlib.c mylib.c
OBJECTS = $(SOURCES:.c=.o)
CFLAGS = -I../include

all: $(TARGET)

$(TARGET): $(OBJECTS)
	$(AR) cvru $@ $^

clean:
	rm -f $(TARGET) $(OBJECTS)
