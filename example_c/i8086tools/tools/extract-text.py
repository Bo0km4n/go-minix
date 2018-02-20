#!/bin/env python
import sys, struct

if len(sys.argv) != 3:
    print "usage: %s input output" % sys.argv[0]
    sys.exit(1)

with open(sys.argv[1], "rb") as f:
    assert f.read(2) == "\x01\x03", "not a.out"
    f.seek(4)
    a_hdrlen = ord(f.read(1))
    f.seek(8)
    a_text = struct.unpack("<L", f.read(4))[0]
    f.seek(a_hdrlen)
    text = f.read(a_text)

with open(sys.argv[2], "wb") as f:
    f.write(text)
