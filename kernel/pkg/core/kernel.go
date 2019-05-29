package core

import (
	"encoding/binary"
	"fmt"
	"io"
	"strings"
)

var K *Kernel

// Kernel is core structure
type Kernel struct {
	Memory *Memory
	CPU    *CPU
}

// Memory has some byte areas
type Memory struct {
	Text []byte
	Data []byte
}

// CPU has registers
type CPU struct {
	generalReg8  map[string]uint8  // AL, CL, DL, BL, AH, CH, DH, BH
	specialReg16 map[string]uint16 // SP, BP, SI, DI
	flag         map[string]bool   // OF, SF, ZF, CF
	ip           int32             // Program Counter
}

func (k *Kernel) PrintParams(w io.Writer) {
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
	w.Write([]byte(s))
}

func centering(s string, l int) string {
	ls := (l - len(s)) / 2
	cs := strings.Repeat(" ", ls) + s + strings.Repeat(" ", l-(ls+len(s)))
	return cs
}

func (k *Kernel) PrintRegs(w io.Writer) {
	s := fmt.Sprintf(
		"%04x %04x %04x %04x %04x %04x %04x %04x %s %04x:",
		k.CPU.AX(),
		k.CPU.BX(),
		k.CPU.CX(),
		k.CPU.DX(),
		k.CPU.SP(),
		k.CPU.BP(),
		k.CPU.SI(),
		k.CPU.DI(),
		k.CPU.DumpFlag(),
		k.CPU.IP(),
	)
	w.Write([]byte(s))
}

func newMemory(textArea, dataArea []byte) *Memory {
	return &Memory{
		Text: textArea,
		Data: dataArea,
	}
}

func newCPU() *CPU {
	return &CPU{
		generalReg8: map[string]uint8{
			"AL": 0x00,
			"CL": 0x00,
			"DL": 0x00,
			"BL": 0x00,
			"AH": 0x00,
			"CH": 0x00,
			"DH": 0x00,
			"BH": 0x00,
		},
		specialReg16: map[string]uint16{
			"SP": 0x0000,
			"BP": 0x0000,
			"SI": 0x0000,
			"DI": 0x0000,
		},
		flag: map[string]bool{
			"OF": false,
			"SF": false,
			"ZF": false,
			"CF": false,
		},
	}
}

func (c *CPU) AX() uint16 {
	return binary.BigEndian.Uint16([]byte{c.generalReg8["AH"], c.generalReg8["AL"]})
}
func (c *CPU) BX() uint16 {
	return binary.BigEndian.Uint16([]byte{c.generalReg8["BH"], c.generalReg8["BL"]})
}
func (c *CPU) CX() uint16 {
	return binary.BigEndian.Uint16([]byte{c.generalReg8["CH"], c.generalReg8["CL"]})
}
func (c *CPU) DX() uint16 {
	return binary.BigEndian.Uint16([]byte{c.generalReg8["DH"], c.generalReg8["DL"]})
}

func (c *CPU) SP() uint16 {
	return c.specialReg16["SP"]
}
func (c *CPU) BP() uint16 {
	return c.specialReg16["BP"]
}
func (c *CPU) SI() uint16 {
	return c.specialReg16["SI"]
}
func (c *CPU) DI() uint16 {
	return c.specialReg16["DI"]
}

func (c *CPU) DumpFlag() string {
	var of, sf, zf, cf string
	if c.flag["OF"] {
		of = "O"
	} else {
		of = "-"
	}

	if c.flag["SF"] {
		sf = "S"
	} else {
		sf = "-"
	}

	if c.flag["ZF"] {
		zf = "Z"
	} else {
		zf = "-"
	}

	if c.flag["CF"] {
		cf = "C"
	} else {
		cf = "-"
	}

	return fmt.Sprintf("%s%s%s%s", of, sf, zf, cf)
}
func (c *CPU) IP() int32 {
	return c.ip
}
