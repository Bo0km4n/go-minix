TARGET = ld
SOURCES = archive.c error.c extract.c finish.c main.c memory.c output.c read.c relocate.c save.c scan.c sym.c write.c
OBJECTS = $(SOURCES:.c=.o)
CFLAGS  = -I../include
LDFLAGS = -L../lib -lobject

all: $(TARGET)

$(TARGET): $(OBJECTS)
	$(CC) -o $@ $^ $(LDFLAGS)

clean:
	rm -f $(TARGET) $(TARGET).exe $(OBJECTS)
