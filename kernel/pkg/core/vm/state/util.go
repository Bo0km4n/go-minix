package state

import (
	"encoding/binary"
	"fmt"
	"strings"
)

func (c *State) PrintParams() {
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
	c.Display.Write([]byte(s))
}

func centering(s string, l int) string {
	ls := (l - len(s)) / 2
	cs := strings.Repeat(" ", ls) + s + strings.Repeat(" ", l-(ls+len(s)))
	return cs
}

func (c *State) PrintRegs() {
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
		c.IP,
	)
	c.Display.Write([]byte(s))
}

var (
	maskMid3 = byte(0x38)
	maskLow3 = byte(0x07)
	maskTop2 = byte(0xc0)
)

type regKeyFunc func(byte) string

func (s *State) Write16(p uint16, d uint16) {
	s.Mem.Data[p] = uint8(d & 0x00ff)
	s.Mem.Data[p+1] = uint8((d & 0xff00) >> 8)
}

func (s *State) Write8(p uint16, d byte) {
	s.Mem.Data[p] = d
}

func (s *State) Read8(p uint16) uint8 {
	return s.Mem.Data[p]
}

func (s *State) Read16(p uint16) uint16 {
	return binary.LittleEndian.Uint16(s.Mem.Data[p : p+2])
}
