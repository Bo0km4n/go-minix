package asem

import (
	"encoding/binary"

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

func MovImmToReg(s *state.State, inst byte) error {
	w := inst & 0x08 >> 3
	regb := inst & maskLow3

	if w == 0x00 {
		data := s.Mem.Text[s.IP+1]
		regKey := s.GetReg8Key(regb)
		reg := s.GeneralReg8[regKey]
		reg.SetVal(data)
	} else if w == 0x01 {
		regKey := s.GetReg16Key(regb)
		reg := s.GeneralReg16[regKey]
		data := binary.LittleEndian.Uint16([]byte{s.Mem.Text[s.IP+1], s.Mem.Text[s.IP+2]})
		reg.SetVal(data)
	}

	return nil
}
