package cpu

import "encoding/binary"

func MOV_RM_Reg(c *CPU, inst, ope byte) error {
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

func MOV_Imm_To_Reg(c *CPU, inst byte) error {
	w := inst & 0x08 >> 3
	regb := inst & maskLow3

	if w == 0x00 {
		data := c.mem.Text[c.ip+1]
		regKey := c.getReg8Key(regb)
		reg := c.generalReg8[regKey]
		reg.SetVal(data)
	} else if w == 0x01 {
		regKey := c.getReg16Key(regb)
		reg := c.generalReg16[regKey]
		data := binary.LittleEndian.Uint16([]byte{c.mem.Text[c.ip+1], c.mem.Text[c.ip+2]})
		reg.SetVal(data)
	}

	return nil
}
