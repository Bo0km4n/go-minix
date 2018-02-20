TARGET  = dis88
SOURCES = disfp.c dishand.c dismain.c disrel.c distabs.c
OBJECTS = $(SOURCES:.c=.o)
CFLAGS  = -I../include
LDFLAGS =

all: $(TARGET)

$(TARGET): $(OBJECTS)
	$(CC) -o $@ $^ $(LDFLAGS)

clean:
	rm -f $(TARGET) $(TARGET).exe $(OBJECTS)
