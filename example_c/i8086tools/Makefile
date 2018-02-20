include Makefile.inc

all install uninstall:
	$(MAKE) $@ -C 7run
	$(MAKE) $@ -C distrib

depend:
	$(MAKE) $@ -C 7run

clean:
	$(MAKE) $@ -C 7run
	$(MAKE) $@ -C tools
	$(MAKE) $@ -C tests
	$(MAKE) $@ -C Ack-5.5
	$(MAKE) $@ -C 8086v6
	$(MAKE) $@ -C libc
	$(MAKE) $@ -C trans
