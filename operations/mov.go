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
	case 0x89:
		opt := ctx.Body[ctx.Idx+1]
		d := 0x00
		w := 0x01
		regFunc := getRegFunc(byte(w))
		mod := opt & maskTop2 >> 6
		reg := opt & maskMid3 >> 3
		rm := opt & maskLow3

		fromOrTo := d == 0x01
		regStr := regFunc(reg)
		return getModRegRM(ctx, mod, rm, fromOrTo, regStr, "mov", regFunc)
	case 0x8a:
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
	case 0xba:
	case 0xbb:
		reg := inst & maskLow3
		rw := Reg16b(reg)
		iw := fmt.Sprintf("%02x%02x", ctx.Body[ctx.Idx+1], ctx.Body[ctx.Idx+2])
		return 3, getResult(ctx.Idx, getOrgOpe(ctx.Body[ctx.Idx:ctx.Idx+3]), getOpeString("mov", rw, iw))
	default:
		return 0, ""
	}
	return 0, ""
}
