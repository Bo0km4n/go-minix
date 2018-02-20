package operations

var reg16 = map[byte]string{
	0x00: "ax",
	0x01: "cx",
	0x02: "dx",
	0x03: "bx",
	0x04: "sp",
	0x05: "bp",
	0x06: "si",
	0x07: "di",
}

var reg8 = map[byte]string{
	0x00: "al",
	0x01: "cl",
	0x02: "dl",
	0x03: "bl",
	0x04: "ah",
	0x05: "ch",
	0x06: "dh",
	0x07: "bh",
}

var ea = map[byte]string{
	0x00: "[bx+si]",
	0x01: "[bx+di]",
	0x02: "[bp+si]",
	0x03: "[bp+di]",
	0x04: "[si]",
	0x05: "[di]",
	0x06: "[bp]",
	0x07: "[bx]",
}

// Reg16b get register(16bit) name
func Reg16b(code byte) string {
	return reg16[code]
}

// Reg8b get register(8bit) name
func Reg8b(code byte) string {
	return reg8[code]
}

// EA get string
func EA(code byte) string {
	return ea[code]
}
