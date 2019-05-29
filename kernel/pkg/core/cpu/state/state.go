package state

import (
	"fmt"
	"io"
	"os"

	"github.com/Bo0km4n/go-minix/kernel/pkg/core/memory"
)

// State is 8086 cpu
type State struct {
	Mem          *memory.Memory
	GeneralReg8  map[string]reg8  // AL, CL, DL, BL, AH, CH, DH, BH
	GeneralReg16 map[string]reg16 // SP, BP, SI, DI
	Flag         map[string]bool  // OF, SF, ZF, CF
	IP           int32            // Program Counter
	Display      io.Writer
	CurInst      byte
}

func NewState(mem *memory.Memory) *State {
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
	return &State{
		Mem: mem,
		GeneralReg8: map[string]reg8{
			"AH": ah,
			"AL": al,
			"BH": bh,
			"BL": bl,
			"CH": ch,
			"CL": cl,
			"DH": dh,
			"DL": dl,
		},
		GeneralReg16: map[string]reg16{
			"AX": ax,
			"BX": bx,
			"CX": cx,
			"DX": dx,
			"SP": sp,
			"BP": bp,
			"SI": si,
			"DI": di,
		},
		Flag: map[string]bool{
			"OF": false,
			"SF": false,
			"ZF": false,
			"CF": false,
		},
		Display: os.Stderr,
	}
}

func (c *State) AX() uint16 {
	return c.GeneralReg16["AX"].GetVal()
}
func (c *State) BX() uint16 {
	return c.GeneralReg16["BX"].GetVal()
}
func (c *State) CX() uint16 {
	return c.GeneralReg16["CX"].GetVal()
}
func (c *State) DX() uint16 {
	return c.GeneralReg16["DX"].GetVal()
}

func (c *State) SP() uint16 {
	return c.GeneralReg16["SP"].GetVal()
}
func (c *State) BP() uint16 {
	return c.GeneralReg16["BP"].GetVal()
}
func (c *State) SI() uint16 {
	return c.GeneralReg16["SI"].GetVal()
}
func (c *State) DI() uint16 {
	return c.GeneralReg16["DI"].GetVal()
}

func (c *State) DumpFlag() string {
	var of, sf, zf, cf string
	if c.Flag["OF"] {
		of = "O"
	} else {
		of = "-"
	}

	if c.Flag["SF"] {
		sf = "S"
	} else {
		sf = "-"
	}

	if c.Flag["ZF"] {
		zf = "Z"
	} else {
		zf = "-"
	}

	if c.Flag["CF"] {
		cf = "C"
	} else {
		cf = "-"
	}

	return fmt.Sprintf("%s%s%s%s", of, sf, zf, cf)
}

func (c *State) GetReg16Key(reg byte) string {
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

func (c *State) GetReg8Key(reg byte) string {
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
