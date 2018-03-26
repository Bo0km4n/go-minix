package operations

import (
	"bytes"
	"encoding/binary"
	"fmt"
)

// SUB model
type SUB struct{}

// Analyze sub analyze
func (sub *SUB) Analyze(ctx *Context, inst byte) (int, string) {
	switch inst {
	case 0x28:
		d := 0x00
		w := 0x01
		opt := ctx.Body[ctx.Idx+1]
		mod := opt & maskTop2 >> 6
		reg := opt & maskMid3 >> 3
		rm := opt & maskLow3
		regFunc := getRegFunc(byte(w))
		regStr := regFunc(reg)

		return getModRegRM(ctx, mod, rm, getFromOrTo(byte(d)), regStr, "sub", regFunc)
	case 0x29:
		d := 0x00
		w := 0x01
		opt := ctx.Body[ctx.Idx+1]
		mod := opt & maskTop2 >> 6
		reg := opt & maskMid3 >> 3
		rm := opt & maskLow3
		regFunc := getRegFunc(byte(w))
		regStr := regFunc(reg)

		return getModRegRM(ctx, mod, rm, getFromOrTo(byte(d)), regStr, "sub", regFunc)
	case 0x2a:
		d := 0x00
		w := 0x01
		opt := ctx.Body[ctx.Idx+1]
		mod := opt & maskTop2 >> 6
		reg := opt & maskMid3 >> 3
		rm := opt & maskLow3
		regFunc := getRegFunc(byte(w))
		regStr := regFunc(reg)

		return getModRegRM(ctx, mod, rm, getFromOrTo(byte(d)), regStr, "sub", regFunc)
	case 0x2b:
		d := 0x00
		w := 0x01
		opt := ctx.Body[ctx.Idx+1]
		mod := opt & maskTop2 >> 6
		reg := opt & maskMid3 >> 3
		rm := opt & maskLow3
		regFunc := getRegFunc(byte(w))
		regStr := regFunc(reg)

		return getModRegRM(ctx, mod, rm, getFromOrTo(byte(d)), regStr, "sub", regFunc)
	case 0x2c:
		data := uint16(ctx.Body[ctx.Idx+1])
		dataStr := fmt.Sprintf("%04x", data)
		return 2, getResult(ctx.Idx, getOrgOpe(ctx.Body[ctx.Idx:ctx.Idx+2]), getOpeString("sub", "al", dataStr))
	case 0x2d:
		var data uint16
		binary.Read(bytes.NewBuffer([]byte{ctx.Body[ctx.Idx+1], ctx.Body[ctx.Idx+2]}), binary.LittleEndian, &data)
		dataStr := fmt.Sprintf("%04x", data)
		return 3, getResult(ctx.Idx, getOrgOpe(ctx.Body[ctx.Idx:ctx.Idx+3]), getOpeString("sub", "ax", dataStr))
	}
	return 999, ""
}
