package operations

import (
	"fmt"
)

// GRP model
type GRP struct{}

// Analyze grp analyze
func (grp *GRP) Analyze(ctx *Context, inst byte) (int, string) {
	switch inst {
	case 0x80:
		mode := (ctx.Body[ctx.Idx+1] & maskMid3) >> 3
		return grp.matchOpe1B(ctx, inst, mode)
	case 0x81:
		mode := (ctx.Body[ctx.Idx+1] & maskMid3) >> 3
		return grp.matchOpe1W(ctx, inst, mode)
	case 0x83:
		mode := (ctx.Body[ctx.Idx+1] & maskMid3) >> 3
		return grp.matchOpe1WB(ctx, inst, mode)
	case 0xf7:
		mode := (ctx.Body[ctx.Idx+1] & maskMid3) >> 3
		return grp.matchOpe3W(ctx, inst, mode)
	case 0xf6:
		mode := (ctx.Body[ctx.Idx+1] & maskMid3) >> 3
		return grp.matchOpe3B(ctx, inst, mode)
	case 0xd1:
		mode := (ctx.Body[ctx.Idx+1] & maskMid3) >> 3
		return grp.matchOpe2(ctx, inst, mode)
	// case 0xd2:
	// 	mode := (ctx.Body[ctx.Idx+1] & maskMid3) >> 3
	// 	return grp.matchOpe2(ctx, inst, mode)
	case 0xff:
		mode := (ctx.Body[ctx.Idx+1] & maskMid3) >> 3
		return grp.matchOpe5(ctx, inst, mode)
	default:
		return NOT_FOUND, ""
	}
}

func (grp *GRP) matchOpe1B(ctx *Context, inst byte, mode byte) (int, string) {
	switch mode {
	case 0x00:
		regMode := inst & 0x01
		regFunc := getRegFunc(regMode)
		regAddr := ctx.Body[ctx.Idx+1] & maskLow3
		regStr := regFunc(regAddr)
		im := fmt.Sprintf("%02x", ctx.Body[ctx.Idx+2])
		return 3, getResult(ctx.Idx, getOrgOpe(ctx.Body[ctx.Idx:ctx.Idx+3]), getOpeString("add", regStr, im))
	case 0x05:
		s := inst & 0x02 >> 1
		w := inst & 0x01
		opt := ctx.Body[ctx.Idx+1]
		mod := opt & maskTop2 >> 6
		rm := opt & maskLow3
		data := ctx.Body[ctx.Idx+2]

		if s == 0x00 || w == 0x01 {
			addtionalData := ctx.Body[ctx.Idx+3]
			ea := ""
			if mod == 0x01 {
				disp := signExtend(data)
				ea = getRM(mod, rm, int(int16(disp)))
			} else {
				ea = getRM(mod, rm, int(data))
			}
			dataStr := fmt.Sprintf("%d", signExtend(addtionalData))

			return 4, getResult(ctx.Idx, getOrgOpe(ctx.Body[ctx.Idx:ctx.Idx+4]), getOpeString("sub", ea, dataStr))
		}
		ea := getRM(mod, rm, int(data))
		dataStr := fmt.Sprintf("%02x", data)

		return 3, getResult(ctx.Idx, getOrgOpe(ctx.Body[ctx.Idx:ctx.Idx+3]), getOpeString("cmp", ea, dataStr))
	case 0x07: // CMP: Immediate with Register/Memory s = 0, w = 0
		opt := ctx.Body[ctx.Idx+1]
		mod := opt & maskTop2 >> 6
		rm := opt & maskLow3
		var disp int
		switch mod {
		case 0x00:
			disp = 0
			ea := getRM(mod, rm, disp)
			dataStr := fmt.Sprintf("%x", ctx.Body[ctx.Idx+2])
			return 3, getResult(ctx.Idx, getOrgOpe(ctx.Body[ctx.Idx:ctx.Idx+3]), getOpeString("cmp byte", ea, dataStr))
		case 0x01:
			disp = int(int16(signExtend(ctx.Body[ctx.Idx+2])))
			ea := getRM(mod, rm, disp)
			dataStr := fmt.Sprintf("%x", ctx.Body[ctx.Idx+3])
			return 4, getResult(ctx.Idx, getOrgOpe(ctx.Body[ctx.Idx:ctx.Idx+4]), getOpeString("cmp byte", ea, dataStr))
		case 0x02:
			disp = joinDispHighAndLow(ctx.Body[ctx.Idx+3], ctx.Body[ctx.Idx+2])
			ea := getRM(mod, rm, disp)
			return 4, getResult(ctx.Idx, getOrgOpe(ctx.Body[ctx.Idx:ctx.Idx+4]), getOpeString("cmp", ea))
		case 0x03:
			regStr := Reg16b(rm)
			dataStr := fmt.Sprintf("%x", ctx.Body[ctx.Idx+2])
			return 3, getResult(ctx.Idx, getOrgOpe(ctx.Body[ctx.Idx:ctx.Idx+3]), getOpeString("cmp", regStr, dataStr))
		}
	default:
		return NOT_FOUND, ""
	}
	return NOT_FOUND, ""
}

func (grp *GRP) matchOpe1WB(ctx *Context, inst byte, mode byte) (int, string) {
	switch mode {
	case 0x00: // ADD:  s = 1, w = 1
		opt := ctx.Body[ctx.Idx+1]
		mod := opt & maskTop2 >> 6
		rm := opt & maskLow3

		switch mod {
		case 0x00:
			disp := 0
			ea := getRM(mod, rm, disp)
			dataStr := fmt.Sprintf("%x", int(int16(signExtend(ctx.Body[ctx.Idx+2]))))
			return 3, getResult(ctx.Idx, getOrgOpe(ctx.Body[ctx.Idx:ctx.Idx+3]), getOpeString("add", ea, dataStr))
		case 0x01:
			disp := int(int16(signExtend(ctx.Body[ctx.Idx+2])))
			ea := getRM(mod, rm, disp)
			dataStr := fmt.Sprintf("%x", int(int16(signExtend(ctx.Body[ctx.Idx+3]))))
			return 4, getResult(ctx.Idx, getOrgOpe(ctx.Body[ctx.Idx:ctx.Idx+4]), getOpeString("add", ea, dataStr))
		case 0x03:
			regStr := Reg16b(rm)
			dataStr := fmt.Sprintf("%x", int(int16(signExtend(ctx.Body[ctx.Idx+2]))))
			return 3, getResult(ctx.Idx, getOrgOpe(ctx.Body[ctx.Idx:ctx.Idx+3]), getOpeString("add", regStr, dataStr))
		}
	case 0x03: // SBB: Immediate from Register/Memory. s = 1, w = 1
		opt := ctx.Body[ctx.Idx+1]
		mod := opt & maskTop2 >> 6
		rm := opt & maskLow3
		return buildGRPOpeStringWithSW(ctx, 0x01, 0x01, mod, rm, "", "sbb")
	case 0x05: // SUB: s = 1, w = 1
		opt := ctx.Body[ctx.Idx+1]
		mod := opt & maskTop2 >> 6
		rm := opt & maskLow3

		switch mod {
		case 0x01:
			disp := int(int16(signExtend(ctx.Body[ctx.Idx+2])))
			ea := getRM(mod, rm, disp)
			dataStr := fmt.Sprintf("%x", int(int16(signExtend(ctx.Body[ctx.Idx+3]))))
			return 4, getResult(ctx.Idx, getOrgOpe(ctx.Body[ctx.Idx:ctx.Idx+4]), getOpeString("sub", ea, dataStr))
		case 0x03:
			regStr := Reg16b(rm)
			dataStr := fmt.Sprintf("%x", int(int16(signExtend(ctx.Body[ctx.Idx+2]))))
			return 3, getResult(ctx.Idx, getOrgOpe(ctx.Body[ctx.Idx:ctx.Idx+3]), getOpeString("sub", regStr, dataStr))
		}

	case 0x07: // CMP Immediate with Register/Memory s = 1, w = 1
		opt := ctx.Body[ctx.Idx+1]
		mod := opt & maskTop2 >> 6
		rm := opt & maskLow3

		switch mod {
		case 0x00:
			disp := ctx.Body[ctx.Idx+2]
			data := ctx.Body[ctx.Idx+3]

			ea := getRM(mod, rm, int(int16(signExtend(disp))))
			dataExtended := fmt.Sprintf("%x", signExtend(data))
			return 4, getResult(ctx.Idx, getOrgOpe(ctx.Body[ctx.Idx:ctx.Idx+4]), getOpeString("cmp", ea, dataExtended))
		case 0x01:
			disp := ctx.Body[ctx.Idx+2]
			data := ctx.Body[ctx.Idx+3]

			ea := getRM(mod, rm, int(int16(signExtend(disp))))
			dataExtended := fmt.Sprintf("%x", signExtend(data))
			return 4, getResult(ctx.Idx, getOrgOpe(ctx.Body[ctx.Idx:ctx.Idx+4]), getOpeString("cmp", ea, dataExtended))
		case 0x02:
			disp := ctx.Body[ctx.Idx+2]
			data := ctx.Body[ctx.Idx+3]

			ea := getRM(mod, rm, int(int16(signExtend(disp))))
			dataExtended := fmt.Sprintf("%x", signExtend(data))
			return 4, getResult(ctx.Idx, getOrgOpe(ctx.Body[ctx.Idx:ctx.Idx+4]), getOpeString("cmp", ea, dataExtended))
		case 0x03:
			regStr := Reg16b(rm)
			dataStr := fmt.Sprintf("%x", int(int16(signExtend(ctx.Body[ctx.Idx+2]))))
			return 3, getResult(ctx.Idx, getOrgOpe(ctx.Body[ctx.Idx:ctx.Idx+3]), getOpeString("cmp", regStr, dataStr))
		}
	}

	return NOT_FOUND, ""
}

func (grp *GRP) matchOpe3W(ctx *Context, inst, mode byte) (int, string) {
	switch mode {
	case 0x00: // TEST: Immediate Data and Register/Memory w = 0
		opt := ctx.Body[ctx.Idx+1]
		mod := opt & maskTop2 >> 6
		rm := opt & maskLow3
		switch mod {
		case 0x01:
			disp := int(int16(signExtend(ctx.Body[ctx.Idx+2])))
			ea := getRM(mod, rm, disp)
			dataStr := fmt.Sprintf("%04x", joinDispHighAndLow(ctx.Body[ctx.Idx+3], ctx.Body[ctx.Idx+4]))
			return 5, getResult(ctx.Idx, getOrgOpe(ctx.Body[ctx.Idx:ctx.Idx+5]), getOpeString("test", ea, dataStr))
		case 0x03:
			regStr := Reg16b(rm)
			dataStr := fmt.Sprintf("%04x", joinDispHighAndLow(ctx.Body[ctx.Idx+2], ctx.Body[ctx.Idx+3]))
			return 4, getResult(ctx.Idx, getOrgOpe(ctx.Body[ctx.Idx:ctx.Idx+4]), getOpeString("or", regStr, dataStr))
		}
	case 0x01: // TEST: Immediate Data and Register/Memory w = 1
	case 0x03: // NEG: Change sign w = 1
		opt := ctx.Body[ctx.Idx+1]
		w := inst & 0x01
		mod := opt & maskTop2 >> 6
		rm := opt & maskLow3
		regFunc := getRegFunc(w)
		return getModRegRM(ctx, mod, rm, false, "", "neg", regFunc)
	case 0x04:
		opt := ctx.Body[ctx.Idx+1]
		mod := opt & maskTop2 >> 6
		rm := opt & maskLow3
		ea := getRM(mod, rm, 0)
		return 2, getResult(ctx.Idx, getOrgOpe(ctx.Body[ctx.Idx:ctx.Idx+2]), getOpeString("mul", ea))
	case 0x06: // DIV: Dicvide(Unsigned)
		opt := ctx.Body[ctx.Idx+1]
		mod := opt & maskTop2 >> 6
		rm := opt & maskLow3
		switch mod {
		case 0x03:
			regStr := Reg16b(rm)
			return 2, getResult(ctx.Idx, getOrgOpe(ctx.Body[ctx.Idx:ctx.Idx+2]), getOpeString("div", regStr))
		}
	}
	return NOT_FOUND, ""
}

func (grp *GRP) matchOpe1W(ctx *Context, inst, mode byte) (int, string) {
	switch mode {
	case 0x00: // OR: Immediate to Register/Memory w = 0
		opt := ctx.Body[ctx.Idx+1]
		mod := opt & maskTop2 >> 6
		rm := opt & maskLow3
		switch mod {
		case 0x03:
			regStr := Reg16b(rm)
			dataStr := fmt.Sprintf("%04x", joinDispHighAndLow(ctx.Body[ctx.Idx+2], ctx.Body[ctx.Idx+3]))
			return 4, getResult(ctx.Idx, getOrgOpe(ctx.Body[ctx.Idx:ctx.Idx+4]), getOpeString("or", regStr, dataStr))
		}
	case 0x01: // OR: Immediate to Register/Memory w = 1
		opt := ctx.Body[ctx.Idx+1]
		mod := opt & maskTop2 >> 6
		rm := opt & maskLow3
		switch mod {
		case 0x01:
			disp := int(int16(signExtend(ctx.Body[ctx.Idx+2])))
			dataStr := fmt.Sprintf("%04x", joinDispHighAndLow(ctx.Body[ctx.Idx+3], ctx.Body[ctx.Idx+4]))
			ea := getRM(mod, rm, disp)
			return 5, getResult(ctx.Idx, getOrgOpe(ctx.Body[ctx.Idx:ctx.Idx+5]), getOpeString("or", ea, dataStr))
		case 0x03:
			regStr := Reg16b(rm)
			dataStr := fmt.Sprintf("%04x", joinDispHighAndLow(ctx.Body[ctx.Idx+2], ctx.Body[ctx.Idx+3]))
			return 4, getResult(ctx.Idx, getOrgOpe(ctx.Body[ctx.Idx:ctx.Idx+4]), getOpeString("or", regStr, dataStr))
		}
	case 0x04: // AND: Immediate to Register/Memory w = 1
		opt := ctx.Body[ctx.Idx+1]
		mod := opt & maskTop2 >> 6
		rm := opt & maskLow3
		switch mod {
		case 0x01:
			disp := int(int16(signExtend(ctx.Body[ctx.Idx+2])))
			ea := getRM(mod, rm, disp)
			dataStr := fmt.Sprintf("%04x", joinDispHighAndLow(ctx.Body[ctx.Idx+3], ctx.Body[ctx.Idx+4]))
			return 5, getResult(ctx.Idx, getOrgOpe(ctx.Body[ctx.Idx:ctx.Idx+5]), getOpeString("and", ea, dataStr))
		case 0x03:
			regStr := Reg16b(rm)
			dataStr := fmt.Sprintf("%04x", joinDispHighAndLow(ctx.Body[ctx.Idx+2], ctx.Body[ctx.Idx+3]))
			return 4, getResult(ctx.Idx, getOrgOpe(ctx.Body[ctx.Idx:ctx.Idx+4]), getOpeString("and", regStr, dataStr))
		}
	case 0x05: // SUB: Immediate from Register/Memory s = 0, w = 1
		opt := ctx.Body[ctx.Idx+1]
		mod := opt & maskTop2 >> 6
		rm := opt & maskLow3
		switch mod {
		case 0x03:
			regStr := Reg16b(rm)
			dataStr := fmt.Sprintf("%04x", joinDispHighAndLow(ctx.Body[ctx.Idx+2], ctx.Body[ctx.Idx+3]))
			return 4, getResult(ctx.Idx, getOrgOpe(ctx.Body[ctx.Idx:ctx.Idx+4]), getOpeString("sub", regStr, dataStr))
		}
	case 0x07: // CMP: s = 0, w = 1
		opt := ctx.Body[ctx.Idx+1]
		mod := opt & maskTop2 >> 6
		rm := opt & maskLow3

		switch mod {
		case 0x00:
			dataHigh8bit := ctx.Body[ctx.Idx+2]
			dataLow8bit := ctx.Body[ctx.Idx+3]
			dataStr := fmt.Sprintf("%02x%02x", dataLow8bit, dataHigh8bit)
			return 4, getResult(ctx.Idx, getOrgOpe(ctx.Body[ctx.Idx:ctx.Idx+4]), getOpeString("cmp", getRM(mod, rm, 0), dataStr))
		case 0x01:
			disp := int(int16(signExtend(ctx.Body[ctx.Idx+2])))
			ea := getRM(mod, rm, disp)
			dataStr := fmt.Sprintf("%04x", joinDispHighAndLow(ctx.Body[ctx.Idx+3], ctx.Body[ctx.Idx+4]))
			return 5, getResult(ctx.Idx, getOrgOpe(ctx.Body[ctx.Idx:ctx.Idx+5]), getOpeString("cmp", ea, dataStr))
		case 0x02:
			dataHigh8bit := ctx.Body[ctx.Idx+2]
			dataLow8bit := ctx.Body[ctx.Idx+3]
			dataStr := fmt.Sprintf("%02x%02x", dataLow8bit, dataHigh8bit)
			return 4, getResult(ctx.Idx, getOrgOpe(ctx.Body[ctx.Idx:ctx.Idx+4]), getOpeString("cmp", getRM(mod, rm, 0), dataStr))
		case 0x03:
			dataHigh8bit := ctx.Body[ctx.Idx+2]
			dataLow8bit := ctx.Body[ctx.Idx+3]
			dataStr := fmt.Sprintf("%02x%02x", dataLow8bit, dataHigh8bit)
			return 4, getResult(ctx.Idx, getOrgOpe(ctx.Body[ctx.Idx:ctx.Idx+4]), getOpeString("cmp", getRM(mod, rm, 0), dataStr))
		}
	}
	return NOT_FOUND, ""
}

func (grp *GRP) matchOpe3B(ctx *Context, inst, mode byte) (int, string) {
	switch mode {
	case 0x00: // test reg data w = 0
		opt := ctx.Body[ctx.Idx+1]
		mod := opt & maskTop2 >> 6
		rm := opt & maskLow3
		var disp int
		switch mod {
		case 0x00:
			disp = 0
			ea := getRM(mod, rm, disp)
			return 3, getResult(ctx.Idx, getOrgOpe(ctx.Body[ctx.Idx:ctx.Idx+4]), getOpeString("test", ea))
		case 0x01:
			disp = int(int16(signExtend(ctx.Body[ctx.Idx+2])))
			ea := getRM(mod, rm, disp)
			dataStr := fmt.Sprintf("%x", ctx.Body[ctx.Idx+3])
			return 4, getResult(ctx.Idx, getOrgOpe(ctx.Body[ctx.Idx:ctx.Idx+4]), getOpeString("test byte", ea, dataStr))
		case 0x02:
			disp = joinDispHighAndLow(ctx.Body[ctx.Idx+3], ctx.Body[ctx.Idx+2])
			ea := getRM(mod, rm, disp)
			return 4, getResult(ctx.Idx, getOrgOpe(ctx.Body[ctx.Idx:ctx.Idx+4]), getOpeString("test", ea))
		case 0x03:
			regStr := Reg8b(rm)
			dataStr := fmt.Sprintf("%x", ctx.Body[ctx.Idx+2])
			return 3, getResult(ctx.Idx, getOrgOpe(ctx.Body[ctx.Idx:ctx.Idx+3]), getOpeString("test", regStr, dataStr))
		}
	default:
		return NOT_FOUND, ""
	}
	return NOT_FOUND, ""
}

func (grp *GRP) matchOpe2(ctx *Context, inst, mode byte) (int, string) {
	switch mode {
	case 0x04: // shl reg data
		var countStr string
		opt := ctx.Body[ctx.Idx+1]
		// mod := opt & maskTop2 >> 6
		rm := opt & maskLow3
		w := inst & 0x01
		v := inst & 0x02
		if v == 0x00 {
			countStr = fmt.Sprintf("%d", 1)
		} else {
			countStr = "cl"
		}
		regStr := getRegFunc(w)(rm)
		return 2, getResult(ctx.Idx, getOrgOpe(ctx.Body[ctx.Idx:ctx.Idx+2]), getOpeString("shl", regStr, countStr))
	case 0x02: // RCL: v = 0, w = 1
		opt := ctx.Body[ctx.Idx+1]
		mod := opt & maskTop2 >> 6
		rm := opt & maskLow3

		switch mod {
		case 0x03:
			regStr := Reg16b(rm)
			countStr := fmt.Sprintf("%x", 1)
			return 2, getResult(ctx.Idx, getOrgOpe(ctx.Body[ctx.Idx:ctx.Idx+2]), getOpeString("rcl", regStr, countStr))
		}

	case 0x05: // SHR: v = 0, w = 1
		opt := ctx.Body[ctx.Idx+1]
		mod := opt & maskTop2 >> 6
		rm := opt & maskLow3

		switch mod {
		case 0x03:
			regStr := Reg16b(rm)
			countStr := fmt.Sprintf("%x", 1)
			return 2, getResult(ctx.Idx, getOrgOpe(ctx.Body[ctx.Idx:ctx.Idx+2]), getOpeString("shr", regStr, countStr))
		}
	case 0x07: // SAR: v = 0, w = 1
		opt := ctx.Body[ctx.Idx+1]
		mod := opt & maskTop2 >> 6
		rm := opt & maskLow3

		switch mod {
		case 0x03:
			regStr := Reg16b(rm)
			countStr := fmt.Sprintf("%x", 1)
			return 2, getResult(ctx.Idx, getOrgOpe(ctx.Body[ctx.Idx:ctx.Idx+2]), getOpeString("sar", regStr, countStr))
		}
	}
	return NOT_FOUND, ""
}

func (grp *GRP) matchOpe5(ctx *Context, inst, mode byte) (int, string) {
	switch mode {
	case 0x00: // INC: Register/Memory w = 1
		opt := ctx.Body[ctx.Idx+1]
		mod := opt & maskTop2 >> 6
		rm := opt & maskLow3
		var disp int
		switch mod {
		case 0x00:
			disp = 0
			ea := getRM(mod, rm, disp)
			return 2, getResult(ctx.Idx, getOrgOpe(ctx.Body[ctx.Idx:ctx.Idx+2]), getOpeString("inc", ea))
		case 0x01:
			disp = int(int16(signExtend(ctx.Body[ctx.Idx+2])))
			ea := getRM(mod, rm, disp)
			return 3, getResult(ctx.Idx, getOrgOpe(ctx.Body[ctx.Idx:ctx.Idx+3]), getOpeString("inc", ea))
		case 0x02:
			disp = joinDispHighAndLow(ctx.Body[ctx.Idx+3], ctx.Body[ctx.Idx+2])
			ea := getRM(mod, rm, disp)
			return 3, getResult(ctx.Idx, getOrgOpe(ctx.Body[ctx.Idx:ctx.Idx+3]), getOpeString("inc", ea))
		case 0x03:
			regStr := Reg16b(rm)
			return 2, getResult(ctx.Idx, getOrgOpe(ctx.Body[ctx.Idx:ctx.Idx+2]), getOpeString("inc", regStr))
		}
	case 0x01: // DEC: Register/memory w = 1
		opt := ctx.Body[ctx.Idx+1]
		mod := opt & maskTop2 >> 6
		rm := opt & maskLow3
		switch mod {
		case 0x01:
			disp := int(int16(signExtend(ctx.Body[ctx.Idx+2])))
			ea := getRM(mod, rm, disp)
			return 3, getResult(ctx.Idx, getOrgOpe(ctx.Body[ctx.Idx:ctx.Idx+3]), getOpeString("dec", ea))
		}
	case 0x04: // JMP: Indirect within Segment
		opt := ctx.Body[ctx.Idx+1]
		mod := opt & maskTop2 >> 6
		rm := opt & maskLow3
		switch mod {
		case 0x03:
			regStr := Reg16b(rm)
			return 2, getResult(ctx.Idx, getOrgOpe(ctx.Body[ctx.Idx:ctx.Idx+2]), getOpeString("jmp", regStr))
		}
	case 0x06: // PUSH: Register/Memory
		opt := ctx.Body[ctx.Idx+1]
		mod := opt & maskTop2 >> 6
		rm := opt & maskLow3

		switch mod {
		case 0x01:
			disp := ctx.Body[ctx.Idx+2]
			ea := getRM(mod, rm, int(disp))
			return 3, getResult(ctx.Idx, getOrgOpe(ctx.Body[ctx.Idx:ctx.Idx+3]), getOpeString("push", ea))
		case 0x00:
			if rm == 0x06 { // exception
				disp := int(uint16(ctx.Body[ctx.Idx+3])<<8 + uint16(ctx.Body[ctx.Idx+2]))
				ea := getRM(mod, rm, disp)
				return 4, getResult(ctx.Idx, getOrgOpe(ctx.Body[ctx.Idx:ctx.Idx+4]), getOpeString("push", ea))
			}
			disp := 0
			ea := getRM(mod, rm, disp)
			return 2, getResult(ctx.Idx, getOrgOpe(ctx.Body[ctx.Idx:ctx.Idx+4]), getOpeString("push", ea))
		case 0x02:
			disp := joinDispHighAndLow(ctx.Body[ctx.Idx+2], ctx.Body[ctx.Idx+3])
			ea := getRM(mod, rm, disp)
			return 4, getResult(ctx.Idx, getOrgOpe(ctx.Body[ctx.Idx:ctx.Idx+4]), getOpeString("push", ea))
		}

	case 0x02:
		opt := ctx.Body[ctx.Idx+1]
		mod := opt & maskTop2 >> 6
		rm := opt & maskLow3
		ea := getRM(mod, rm, 0)
		return 2, getResult(ctx.Idx, getOrgOpe(ctx.Body[ctx.Idx:ctx.Idx+2]), getOpeString("call", ea))
	}

	return NOT_FOUND, ""
}
