package operations

import (
	"bytes"
	"encoding/binary"
	"errors"
	"fmt"
)

func getResult(idx int, ope, opeString string) string {
	return fmt.Sprintf("%04x: %-10s %s", idx, ope, opeString)
}

func getOrgOpe(args []byte) string {
	var buffer bytes.Buffer
	for _, v := range args {
		s := fmt.Sprintf("%02x", v)
		buffer.WriteString(s)
	}
	return buffer.String()
}

func getOpeString(prefix string, args ...string) string {
	var buffer bytes.Buffer
	buffer.WriteString(prefix)
	buffer.WriteString(" ")

	for i, v := range args {
		buffer.WriteString(v)
		if (i != len(args)-1) && args[i+1] != "" {
			buffer.WriteString(", ")
		}
	}
	return buffer.String()
}

func getRegFunc(mode byte) func(byte) string {
	switch mode {
	case 0x00:
		return Reg8b
	case 0x01:
		return Reg16b
	default:
		return nil
	}
}

func getRM(mod, rm byte, disp int) string {
	var dispStr string
	if mod == 0x03 {
		return Reg16b(rm)
	}
	if mod == 0x00 && rm == 0x06 {
		return fmt.Sprintf("[%04x]", disp)
	}
	if disp > 0 && mod != 0x00 {
		dispStr = fmt.Sprintf("+%x", disp)
	} else if disp < 0 && mod != 0x00 {
		dispStr = fmt.Sprintf("%x", disp)
	} else {
		dispStr = fmt.Sprintf("")
	}
	switch rm {
	case 0x00:
		return fmt.Sprintf("[bx+si%s]", dispStr)
	case 0x01:
		return fmt.Sprintf("[bx+di%s]", dispStr)
	case 0x02:
		return fmt.Sprintf("[bp+si%s]", dispStr)
	case 0x03:
		return fmt.Sprintf("[bp+di%s]", dispStr)
	case 0x04:
		return fmt.Sprintf("[si%s]", dispStr)
	case 0x05:
		return fmt.Sprintf("[di%s]", dispStr)
	case 0x06:
		return fmt.Sprintf("[bp%s]", dispStr)
	case 0x07:
		return fmt.Sprintf("[bx%s]", dispStr)
	default:
		return ""
	}
}

func getDispStr(mod byte, disp int) (error, string) {
	switch mod {
	case 0x00:
		return nil, fmt.Sprintf("%02x", disp)
	case 0x01:
		return nil, fmt.Sprintf("%04x", disp)
	case 0x02:
		return nil, fmt.Sprintf("%04x", disp)
	default:
		return errors.New("Not found mod in disp"), ""
	}
}

func joinDispHighAndLow(high, low byte) int {
	var disp int
	var disp16 uint16
	binary.Read(bytes.NewBuffer([]byte{high, low}), binary.LittleEndian, &disp16)
	disp = int(disp16)
	return disp
}

func signExtend(disp byte) uint16 {
	sign := (disp & 0x80) >> 7

	switch sign {
	case 0x00:
		result := uint16(0x0000)
		result = result + uint16(disp)
		return result
	case 0x01:
		result := uint16(0xff00)
		result = result + uint16(disp)
		return result
	}
	return uint16(0)
}

func getModRegRM(ctx *Context, mod, rm byte, fromOrTo bool, regStr, inst string, regFunc func(byte) string) (int, string) {
	switch {
	case mod == 0x00:
		disp := 0
		if rm == 0x06 {
			disp = joinDispHighAndLow(ctx.Body[ctx.Idx+2], ctx.Body[ctx.Idx+3])
			ea := getRM(mod, rm, disp)

			if fromOrTo {
				return 4, getResult(ctx.Idx, getOrgOpe(ctx.Body[ctx.Idx:ctx.Idx+4]), getOpeString(inst, regStr, ea))
			}
			return 4, getResult(ctx.Idx, getOrgOpe(ctx.Body[ctx.Idx:ctx.Idx+4]), getOpeString(inst, ea, regStr))
		}
		ea := getRM(mod, rm, disp)

		if fromOrTo {
			return 2, getResult(ctx.Idx, getOrgOpe(ctx.Body[ctx.Idx:ctx.Idx+2]), getOpeString(inst, regStr, ea))
		}
		return 2, getResult(ctx.Idx, getOrgOpe(ctx.Body[ctx.Idx:ctx.Idx+2]), getOpeString(inst, ea, regStr))

	case mod == 0x01:
		disp := signExtend(ctx.Body[ctx.Idx+2])
		ea := getRM(mod, rm, int(int16(disp)))
		if fromOrTo {
			return 3, getResult(ctx.Idx, getOrgOpe(ctx.Body[ctx.Idx:ctx.Idx+3]), getOpeString(inst, regStr, ea))
		}
		return 3, getResult(ctx.Idx, getOrgOpe(ctx.Body[ctx.Idx:ctx.Idx+3]), getOpeString(inst, ea, regStr))

	case mod == 0x02:
		disp := joinDispHighAndLow(ctx.Body[ctx.Idx+2], ctx.Body[ctx.Idx+3])
		ea := getRM(mod, rm, disp)

		if fromOrTo {
			return 4, getResult(ctx.Idx, getOrgOpe(ctx.Body[ctx.Idx:ctx.Idx+4]), getOpeString(inst, regStr, ea))
		}
		return 4, getResult(ctx.Idx, getOrgOpe(ctx.Body[ctx.Idx:ctx.Idx+4]), getOpeString(inst, ea, regStr))

	case mod == 0x03:
		rmReg := regFunc(rm)

		if fromOrTo {
			return 2, getResult(ctx.Idx, getOrgOpe(ctx.Body[ctx.Idx:ctx.Idx+2]), getOpeString(inst, regStr, rmReg))
		}
		return 2, getResult(ctx.Idx, getOrgOpe(ctx.Body[ctx.Idx:ctx.Idx+2]), getOpeString(inst, rmReg, regStr))
	}
	return NOT_FOUND, ""
}

func getFromOrTo(d byte) bool {
	return d == 0x01
}
