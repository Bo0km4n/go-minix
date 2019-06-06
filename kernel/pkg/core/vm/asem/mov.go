package asem

import (
	"encoding/binary"
	"errors"

	"github.com/Bo0km4n/go-minix/kernel/pkg/core/vm/state"
)

func MOV_RM_Reg(c *state.State, inst, ope byte) error {
	// d := inst & 0x02 >> 1
	// w := inst & 0x01
	// mod := ope & maskTop2 >> 6
	// reg := ope & maskMid3 >> 3
	// rm := ope & maskLow3

	// var regKeyFunc regKeyFunc
	// if w == 0x00 {
	// 	regKeyFunc = c.getReg8Key
	// } else if w == 0x01 {
	//     regKeyFunc = c.getReg16Key
	// }

	// if d == 0x00 { // to reg

	// } else if d == 0x01 { // from reg

	// }
	return nil
}

// 1 0 1 1 w reg
func MovImmToReg(s *state.State, inst byte) (int, error) {
	w := inst & 0x08 >> 3
	regb := inst & maskLow3

	if w == 0x00 {
		data := s.Mem.Text[s.IP+1]
		regKey := s.GetReg8Key(regb)
		reg := s.GeneralReg8[regKey]
		reg.SetVal(data)
		printInstBytes(s, s.Mem.Text[s.IP:s.IP+2])
		return 2, nil
	} else {
		regKey := s.GetReg16Key(regb)
		reg := s.GeneralReg16[regKey]
		data := binary.LittleEndian.Uint16([]byte{s.Mem.Text[s.IP+1], s.Mem.Text[s.IP+2]})
		reg.SetVal(data)
		printInstBytes(s, s.Mem.Text[s.IP:s.IP+3])
		return 3, nil
	}
}

// 1 0 0 0 1 0 d w
func MovRmToRm(s *state.State, inst byte) (int, error) {
	d := inst & 0x02 >> 1
	w := inst & 0x01

	op := s.Mem.Text[s.IP+1]
	mod := op & maskTop2 >> 6
	reg := op & maskMid3 >> 3
	rm := op & maskLow3
	if w == 0x00 {
		reg8 := s.GeneralReg8[s.GetReg8Key(reg)]
		switch mod {
		case 0x00:
			if rm == 0x06 {
				disp := binary.LittleEndian.Uint16([]byte{s.Mem.Text[s.IP+2], s.Mem.Text[s.IP+3]})
				if d == 0x00 {
					s.Write8(disp, reg8.GetVal())
				} else {
					reg8.SetVal(s.Read8(disp))
				}
				printInstBytes(s, s.Mem.Text[s.IP:s.IP+4])
				return 4, nil
			} else {
				ea := uint16(reg8.GetVal())
				if d == 0x00 {
					s.Write8(ea, reg8.GetVal())
				} else {
					reg8.SetVal(s.Read8(ea))
				}
				printInstBytes(s, s.Mem.Text[s.IP:s.IP+3])
				return 3, nil
			}
		case 0x01:
			disp := int16(uint16(s.Mem.Text[s.IP+2])<<8) >> 8
			var ea uint16
			if disp >= 0 {
				ea = uint16(reg8.GetVal()) + uint16(disp)
			} else {
				ea = uint16(reg8.GetVal()) - uint16(disp)
			}
			if d == 0x00 {
				s.Write8(ea, reg8.GetVal())
			} else {
				reg8.SetVal(s.Read8(ea))
			}
			printInstBytes(s, s.Mem.Text[s.IP:s.IP+3])
			return 3, nil
		case 0x02:
			disp := int16(
				binary.LittleEndian.Uint16([]byte{s.Mem.Text[s.IP+2], s.Mem.Text[s.IP+3]}),
			)
			var ea uint16
			if disp >= 0 {
				ea = uint16(reg8.GetVal()) + uint16(disp)
			} else {
				ea = uint16(reg8.GetVal()) - uint16(disp)
			}
			if d == 0x00 {
				s.Write8(ea, reg8.GetVal())
			} else {
				reg8.SetVal(s.Read8(ea))
			}
			printInstBytes(s, s.Mem.Text[s.IP:s.IP+4])
			return 4, nil

		case 0x03:
			rmReg := s.GeneralReg8[s.GetReg8Key(rm)]
			if d == 0x00 { // from reg
				rmReg.SetVal(reg8.GetVal())
			} else { // to reg
				reg8.SetVal(rmReg.GetVal())
			}
			printInstBytes(s, s.Mem.Text[s.IP:s.IP+2])
			return 2, nil
		}
	} else {
		reg16 := s.GeneralReg16[s.GetReg16Key(reg)]
		switch mod {
		case 0x00:
			if rm == 0x06 {
				disp := binary.LittleEndian.Uint16([]byte{s.Mem.Text[s.IP+2], s.Mem.Text[s.IP+3]})
				if d == 0x00 {
					s.Write16(disp, reg16.GetVal())
				} else {
					reg16.SetVal(s.Read16(disp))
				}
				printInstBytes(s, s.Mem.Text[s.IP:s.IP+4])
				return 4, nil
			} else {
				ea := reg16.GetVal()
				if d == 0x00 {
					s.Write16(ea, reg16.GetVal())
				} else {
					reg16.SetVal(s.Read16(ea))
				}
				printInstBytes(s, s.Mem.Text[s.IP:s.IP+3])
				return 3, nil
			}
		case 0x01:
			disp := int16(uint16(s.Mem.Text[s.IP+2])<<8) >> 8
			var ea uint16
			if disp >= 0 {
				ea = reg16.GetVal() + uint16(disp)
			} else {
				ea = reg16.GetVal() - uint16(disp)
			}
			if d == 0x00 {
				s.Write16(ea, reg16.GetVal())
			} else {
				reg16.SetVal(s.Read16(ea))
			}

			printInstBytes(s, s.Mem.Text[s.IP:s.IP+3])
			return 3, nil
		case 0x02:
			disp := int16(
				binary.LittleEndian.Uint16([]byte{s.Mem.Text[s.IP+2], s.Mem.Text[s.IP+3]}),
			)
			var ea uint16
			if disp >= 0 {
				ea = reg16.GetVal() + uint16(disp)
			} else {
				ea = reg16.GetVal() - uint16(disp)
			}
			if d == 0x00 {
				s.Write16(ea, reg16.GetVal())
			} else {
				reg16.SetVal(s.Read16(ea))
			}
			return 4, nil

		case 0x03:
			rmReg := s.GeneralReg16[s.GetReg16Key(rm)]
			if d == 0x00 { // from reg
				rmReg.SetVal(reg16.GetVal())
			} else { // to reg
				reg16.SetVal(rmReg.GetVal())
			}
			return 2, nil
		}
	}

	return 0, errors.New("Not found MovRmToRm case")
}
