package core

var K *Kernel

// Kernel is core structure
type Kernel struct {
	Memory *Memory
	CPU    *CPU
	PC     int // Program Counter
}

// Memory has some byte areas
type Memory struct {
	Text []byte
	Data []byte
}

// CPU has registers
type CPU struct {
	GeneralReg16 map[string]uint16 // AX, CX, DX, BX
	GeneralReg8  map[string]uint8  // AL, CL, DL, BL, AH, CH, DH, BH
	SpecialReg16 map[string]uint16 // SP, BP, SI, DI
	Flag         map[string]bool   // OF, DF, IF, TF
}

func newMemory(textArea, dataArea []byte) *Memory {
	return &Memory{
		Text: textArea,
		Data: dataArea,
	}
}
