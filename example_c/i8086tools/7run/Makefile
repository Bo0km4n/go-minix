include ../Makefile.inc

TARGET   = 7run
CXX      = g++
CXXFLAGS = -Wall -O2 -g
LDFLAGS  =
OBJECTS  = $(SOURCES:%.cpp=%.o)
SOURCES  = main.cpp utils.cpp File.cpp \
	   VMBase.cpp UnixBase.cpp UnixBase.sys.cpp \
	   i8086/OpCode.cpp i8086/Operand.cpp \
	   i8086/VM.cpp i8086/VM.inst.cpp i8086/disasm.cpp \
	   Minix2/OS.cpp Minix2/OS.sys.cpp Minix2/OS.signal.cpp \
	   PDP11/OpCode.cpp PDP11/Operand.cpp \
	   PDP11/VM.cpp PDP11/VM.inst.cpp PDP11/disasm.cpp \
	   UnixV6/OS.cpp UnixV6/OS.sys.cpp \
	   UnixV6/OSPDP11.cpp UnixV6/OSi8086.cpp

all: $(TARGET)

.SUFFIXES: .cpp .o
.cpp.o:
	$(CXX) $(CXXFLAGS) -o $@ -c $<

$(TARGET): $(OBJECTS)
	$(CXX) $(CXXFLAGS) -o $@ $(OBJECTS) $(LDFLAGS)

clean:
	rm -f $(TARGET) $(TARGET).exe $(OBJECTS) *core

install: $(TARGET)
	mkdir -p $(PREFIX)/bin
	install -cs $(TARGET) $(PREFIX)/bin

uninstall:
	cd $(PREFIX)/bin && rm -f $(TARGET) $(TARGET).exe

depend:
	rm -f dependencies
	for cpp in $(SOURCES); do \
	  (echo -n `dirname $$cpp`/; \
	   g++ -MM $(CXXFLAGS) $$cpp) >> dependencies; \
	done

-include dependencies
