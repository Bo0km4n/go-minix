include ../Makefile.inc

M2CMDS   = bin/aal bin/ar bin/cc bin/nm bin/strip \
	   bin/crc bin/install bin/m2 \
	   lib/as lib/ld lib/cv
V6CMDS   = ar as cc ld nm strip
V7CMDS   = ar as cc ld
RUN      = 7run

all:

install:
	tar cvf - v6root v7root 8086v6-ack | tar xf - -C $(PREFIX)
	cd $(PREFIX)/v6root/bin && chmod 755 ar as cc ld nm strip
	cd $(PREFIX)/v6root/lib && chmod 755 as2 c0 c1 c2
	cd $(PREFIX)/v7root/bin && chmod 755 ar as cc ld
	cd $(PREFIX)/v7root/lib && chmod 755 as2 cpp c0 c1 c2
	cd $(PREFIX)/8086v6-ack/bin && chmod 755 cc
	cd $(PREFIX)/8086v6-ack/lib && chmod 755 as cv irrel ld ncem ncg ncpp nopt
	mkdir -p $(M2ROOT)/7bin $(V6ROOT)/7bin $(V7ROOT)/7bin
	for cmd in $(M2CMDS); do \
	  sh mkwrap.sh $(PREFIX)/bin/$(RUN) $(M2ROOT) $(PREFIX)/bin/m2 usr/$$cmd; \
	  cp $(PREFIX)/bin/m2`basename $$cmd` $(M2ROOT)/7bin/`basename $$cmd`; \
	done
	for cmd in $(V6CMDS); do \
	  sh mkwrap.sh -x $(PREFIX)/bin/$(RUN) $(V6ROOT) $(PREFIX)/bin/v6 bin/$$cmd; \
	  cp $(PREFIX)/bin/v6`basename $$cmd` $(V6ROOT)/7bin/`basename $$cmd`; \
	done
	for cmd in $(V7CMDS); do \
	  sh mkwrap.sh -x $(PREFIX)/bin/$(RUN) $(V7ROOT) $(PREFIX)/bin/v7 bin/$$cmd -7; \
	  cp $(PREFIX)/bin/v7`basename $$cmd` $(V7ROOT)/7bin/`basename $$cmd`; \
	done
	(echo '#!/bin/sh' && echo 'PATH="$(M2ROOT)/7bin:$$PATH" "$$@"') > $(PREFIX)/bin/m2path
	(echo '#!/bin/sh' && echo 'PATH="$(V6ROOT)/7bin:$$PATH" "$$@"') > $(PREFIX)/bin/v6path
	(echo '#!/bin/sh' && echo 'PATH="$(V7ROOT)/7bin:$$PATH" "$$@"') > $(PREFIX)/bin/v7path
	cd $(PREFIX)/bin && chmod 755 m2path v6path v7path
	sh mkwrap.sh $(PREFIX)/bin/$(RUN) $(ACKROOT) $(PREFIX)/bin/8086v6- bin/cc
	sh mkwrap.sh $(PREFIX)/bin/$(RUN) $(ACKROOT) $(PREFIX)/bin/8086v6- lib/cv

uninstall:
	cd $(PREFIX)/bin && rm -f m2path v6path v7path 8086v6-cc 8086v6-cv
	for f in `find v6root -type f`; do rm -f $(PREFIX)/$$f; done
	for f in `find v7root -type f`; do rm -f $(PREFIX)/$$f; done
	for f in `find 8086v6-ack -type f`; do rm -f $(PREFIX)/$$f; done
	for cmd in $(M2CMDS); do \
	  rm -f $(PREFIX)/bin/m2`basename $$cmd`; \
	  rm -f $(M2ROOT)/7bin/`basename $$cmd`; \
	done
	for cmd in $(V6CMDS); do \
	  rm -f $(PREFIX)/bin/v6`basename $$cmd`; \
	  rm -f $(V6ROOT)/7bin/`basename $$cmd`; \
	done
	for cmd in $(V7CMDS); do \
	  rm -f $(PREFIX)/bin/v7`basename $$cmd`; \
	  rm -f $(V7ROOT)/7bin/`basename $$cmd`; \
	done
	-rmdir $(M2ROOT)/7bin $(V6ROOT)/7bin $(V7ROOT)/7bin
	-rmdir $(V6BIN) $(V6LIB) $(V6ROOT)
	-rmdir $(V7BIN) $(V7LIB) $(V7ROOT)/usr/include/sys $(V7ROOT)/usr/include $(V7ROOT)/usr $(V7ROOT)
	-rmdir $(ACKBIN) $(ACKLIB) $(ACKROOT)
