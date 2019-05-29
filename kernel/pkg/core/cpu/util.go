package cpu

import (
	"fmt"
	"strings"
)

func (c *CPU) printParams() {
	s := fmt.Sprintf("%s %s %s %s %s %s %s %s %s %s\n",
		centering("AX", 4),
		centering("BX", 4),
		centering("CX", 4),
		centering("DX", 4),
		centering("SP", 4),
		centering("BP", 4),
		centering("SI", 4),
		centering("DI", 4),
		centering("FLAGS", 5),
		centering("IP", 4),
	)
	c.display.Write([]byte(s))
}

func centering(s string, l int) string {
	ls := (l - len(s)) / 2
	cs := strings.Repeat(" ", ls) + s + strings.Repeat(" ", l-(ls+len(s)))
	return cs
}

func (c *CPU) printRegs() {
	s := fmt.Sprintf(
		"%04x %04x %04x %04x %04x %04x %04x %04x %s %04x:",
		c.AX(),
		c.BX(),
		c.CX(),
		c.DX(),
		c.SP(),
		c.BP(),
		c.SI(),
		c.DI(),
		c.DumpFlag(),
		c.IP(),
	)
	c.display.Write([]byte(s))
}

var (
	maskMid3 = byte(0x38)
	maskLow3 = byte(0x07)
	maskTop2 = byte(0xc0)
)

type regKeyFunc func(byte) string
