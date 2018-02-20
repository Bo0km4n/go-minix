TARGET = ar
SOURCES = archiver.c
OBJECTS = $(SOURCES:.c=.o)
CFLAGS  = -I../include
LDFLAGS = -L../lib -lobject

all: $(TARGET)

$(TARGET): $(OBJECTS)
	$(CC) -o $@ $^ $(LDFLAGS)

clean:
	rm -f $(TARGET) $(TARGET).exe $(OBJECTS)
