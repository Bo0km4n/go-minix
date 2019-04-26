package operations

import (
	"bytes"
	"encoding/binary"
	"fmt"
)

// MOV model
type MOV struct{}

// Analyze mov analyze
func (mov *MOV) Analyze(ctx *Context, inst byte) (int, string) {
	switch inst {
	case 0x88: // Register/Memory to /from Register d = 0, w = 0
		opt := ctx.Body[ctx.Idx+1]
		regFunc := Reg8b
		mod := opt & maskTop2 >> 6
		reg := opt & maskMid3 >> 3
		rm := opt & maskLow3

		fromOrTo := false
		regStr := regFunc(reg)
		return getModRegRM(ctx, mod, rm, fromOrTo, regStr, "mov", regFunc)
	case 0x89: // Register/Memory to /from Register d = 0, w = 1
		opt := ctx.Body[ctx.Idx+1]
		regFunc := Reg16b
		mod := opt & maskTop2 >> 6
		reg := opt & maskMid3 >> 3
		rm := opt & maskLow3

		fromOrTo := false
		regStr := regFunc(reg)
		return getModRegRM(ctx, mod, rm, fromOrTo, regStr, "mov", regFunc)
	case 0x8a: // Register/Memory to /from Register d = 1, w = 0
		opt := ctx.Body[ctx.Idx+1]
		d := 0x01
		w := 0x00
		regFunc := getRegFunc(byte(w))
		mod := opt & maskTop2 >> 6
		reg := opt & maskMid3 >> 3
		rm := opt & maskLow3

		fromOrTo := d == 0x01
		regStr := regFunc(reg)
		return getModRegRM(ctx, mod, rm, fromOrTo, regStr, "mov", regFunc)
	case 0x8b:
		opt := ctx.Body[ctx.Idx+1]
		d := 0x01
		w := 0x01
		regFunc := getRegFunc(byte(w))
		mod := opt & maskTop2 >> 6
		reg := opt & maskMid3 >> 3
		rm := opt & maskLow3

		fromOrTo := d == 0x01
		regStr := regFunc(reg)
		return getModRegRM(ctx, mod, rm, fromOrTo, regStr, "mov", regFunc)
	case 0xb8:
		regCode := inst & maskLow3
		reg := Reg16b(regCode)
		im := getOrgOpe([]byte{ctx.Body[ctx.Idx+2], ctx.Body[ctx.Idx+1]})
		return 3, getResult(ctx.Idx, getOrgOpe(ctx.Body[ctx.Idx:ctx.Idx+3]), getOpeString("mov", reg, im))
	case 0xb9:
		var data uint16
		reg := inst & maskLow3
		binary.Read(bytes.NewBuffer([]byte{ctx.Body[ctx.Idx+1], ctx.Body[ctx.Idx+2]}), binary.LittleEndian, &data)
		regFunc := getRegFunc(0x01)
		regStr := regFunc(reg)
		dataStr := fmt.Sprintf("%04x", data)
		return 3, getResult(ctx.Idx, getOrgOpe(ctx.Body[ctx.Idx:ctx.Idx+3]), getOpeString("mov", regStr, dataStr))
	case 0xba: // Immediate to register w = 1 reg = 010
		reg := inst & maskLow3
		rw := Reg16b(reg)
		iw := fmt.Sprintf("%02x%02x", ctx.Body[ctx.Idx+2], ctx.Body[ctx.Idx+1])
		return 3, getResult(ctx.Idx, getOrgOpe(ctx.Body[ctx.Idx:ctx.Idx+3]), getOpeString("mov", rw, iw))

	case 0xbb: // Immediate to register w = 1 reg = 011
		reg := inst & maskLow3
		rw := Reg16b(reg)
		iw := fmt.Sprintf("%02x%02x", ctx.Body[ctx.Idx+2], ctx.Body[ctx.Idx+1])
		return 3, getResult(ctx.Idx, getOrgOpe(ctx.Body[ctx.Idx:ctx.Idx+3]), getOpeString("mov", rw, iw))

	case 0xc6: // Immediate to Register/Memory w = 0
		opt := ctx.Body[ctx.Idx+1]
		mod := opt & maskTop2 >> 6
		rm := opt & maskLow3
		disp := signExtend(ctx.Body[ctx.Idx+2])
		ea := getRM(mod, rm, int(int16(disp)))
		err, dataStr := getDispStr(mod, int(int16(disp)))
		if err != nil {
			panic(err)
		}

		return 3, getResult(ctx.Idx, getOrgOpe(ctx.Body[ctx.Idx:ctx.Idx+3]), getOpeString("mov", "byte "+ea, dataStr))

	case 0xc7: // Immediate to Register/Memory w = 1
		opt := ctx.Body[ctx.Idx+1]
		mod := opt & maskTop2 >> 6
		rm := opt & maskLow3

		switch mod {
		case 0x00:
			if rm == 0x06 { // exception
				disp := int((uint16(ctx.Body[ctx.Idx+3]) << 8) + uint16(ctx.Body[ctx.Idx+2]))
				ea := getRM(mod, rm, int(int16(disp)))
				dataStr := fmt.Sprintf("%04x", joinDispHighAndLow(ctx.Body[ctx.Idx+4], ctx.Body[ctx.Idx+5]))
				return 6, getResult(ctx.Idx, getOrgOpe(ctx.Body[ctx.Idx:ctx.Idx+5]), getOpeString("mov", ea, dataStr))
			}
		case 0x01:
			disp := signExtend(ctx.Body[ctx.Idx+2])
			ea := getRM(mod, rm, int(int16(disp)))
			dataStr := fmt.Sprintf("%04x", joinDispHighAndLow(ctx.Body[ctx.Idx+3], ctx.Body[ctx.Idx+4]))
			return 5, getResult(ctx.Idx, getOrgOpe(ctx.Body[ctx.Idx:ctx.Idx+5]), getOpeString("mov", ea, dataStr))
		}
	}
	return NOT_FOUND, ""
}
