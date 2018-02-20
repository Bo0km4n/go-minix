TARGET = fatfile
M2ROOT = .

all: crc

crc: $(TARGET).o
	m2crc $(TARGET).*

$(TARGET).o: $(TARGET).s
	7run $(M2ROOT)/usr/lib/as - -o $@ $<

$(TARGET).s: $(TARGET).m
	7run $(M2ROOT)/usr/lib/ncg $< $@

$(TARGET).m: $(TARGET).k
	7run k2m $(M2ROOT)/usr/lib/nopt $< $@

$(TARGET).k: $(TARGET).i
	7run $(M2ROOT)/usr/lib/ncem -L $< $@

$(TARGET).i: $(TARGET).c
	7run -r . $(M2ROOT)/usr/lib/ncpp -D_EM_WSIZE=2 -D_EM_PSIZE=2 -D_EM_SSIZE=2 -D_EM_LSIZE=4 -D_EM_FSIZE=4 -D_EM_DSIZE=8 -D__ACK__ -D__minix -D__i86 $< > $@
	strip-cr $@ $@.orig
	cp $@.orig $@
	7run $(M2ROOT)/usr/lib/irrel $@

clean:
	rm -f $(TARGET).o $(TARGET).s $(TARGET).m $(TARGET).k $(TARGET).i $(TARGET).i.orig
