package cpu

import (
	"fmt"
	"io"
	"os"

	"github.com/Bo0km4n/go-minix/kernel/pkg/core/memory"
)

// CPU is 8086 cpu
type CPU struct {
	mem          *memory.Memory
	generalReg8  map[string]reg8  // AL, CL, DL, BL, AH, CH, DH, BH
	generalReg16 map[string]reg16 // SP, BP, SI, DI
	flag         map[string]bool  // OF, SF, ZF, CF
	ip           int32            // Program Counter
	display      io.Writer
	curInst      byte
}

func NewCPU(mem *memory.Memory) *CPU {
	al := &AL{val: 0x00}
	ah := &AH{val: 0x00}
	ax := &AX{al: al, ah: ah}
	bl := &BL{val: 0x00}
	bh := &BH{val: 0x00}
	bx := &BX{bl: bl, bh: bh}
	cl := &CL{val: 0x00}
	ch := &CH{val: 0x00}
	cx := &CX{cl: cl, ch: ch}
	dl := &DL{val: 0x00}
	dh := &DH{val: 0x00}
	dx := &DX{dl: dl, dh: dh}
	sp := &SP{val: 0x0000}
	bp := &BP{val: 0x0000}
	si := &SI{val: 0x0000}
	di := &DI{val: 0x0000}
	return &CPU{
		mem: mem,
		generalReg8: map[string]reg8{
			"AH": ah,
			"AL": al,
			"BH": bh,
			"BL": bl,
			"CH": ch,
			"CL": cl,
			"DH": dh,
			"DL": dl,
		},
		generalReg16: map[string]reg16{
			"AX": ax,
			"BX": bx,
			"CX": cx,
			"DX": dx,
			"SP": sp,
			"BP": bp,
			"SI": si,
			"DI": di,
		},
		flag: map[string]bool{
			"OF": false,
			"SF": false,
			"ZF": false,
			"CF": false,
		},
		display: os.Stderr,
	}
}

func (c *CPU) AX() uint16 {
	return c.generalReg16["AX"].GetVal()
}
func (c *CPU) BX() uint16 {
	return c.generalReg16["BX"].GetVal()
}
func (c *CPU) CX() uint16 {
	return c.generalReg16["CX"].GetVal()
}
func (c *CPU) DX() uint16 {
	return c.generalReg16["DX"].GetVal()
}

func (c *CPU) SP() uint16 {
	return c.generalReg16["SP"].GetVal()
}
func (c *CPU) BP() uint16 {
	return c.generalReg16["BP"].GetVal()
}
func (c *CPU) SI() uint16 {
	return c.generalReg16["SI"].GetVal()
}
func (c *CPU) DI() uint16 {
	return c.generalReg16["DI"].GetVal()
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

func (c *CPU) Exec() error {
	if c.ip == 0 {
		c.printParams()
	}
	c.printRegs()
	c.fetch()
	return c.execAsem()
}

func (c *CPU) getReg16Key(reg byte) string {
	switch reg {
	case 0x00:
		return "AX"
	case 0x01:
		return "CX"
	case 0x02:
		return "DX"
	case 0x03:
		return "BX"
	case 0x04:
		return "SP"
	case 0x05:
		return "BP"
	case 0x06:
		return "SI"
	case 0x07:
		return "DI"
	}
	return ""
}

func (c *CPU) getReg8Key(reg byte) string {
	switch reg {
	case 0x00:
		return "AL"
	case 0x01:
		return "CL"
	case 0x02:
		return "DL"
	case 0x03:
		return "BL"
	case 0x04:
		return "AH"
	case 0x05:
		return "CH"
	case 0x06:
		return "DH"
	case 0x07:
		return "BH"
	}
	return ""
}
