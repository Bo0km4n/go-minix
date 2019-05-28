package core

var K *Kernel

type Kernel struct {
	Memory *Memory
	CPU    *CPU
}

type Memory struct {
	Text []byte
	Data []byte
}

type CPU struct {
	GeneralReg16 map[string]uint16 // AX, CX, DX, BX
	GeneralReg8  map[string]uint8  // AL, CL, DL, BL, AH, CH, DH, BH
	SpecialReg16 map[string]uint16 // SP, BP, SI, DI
	Flag         map[string]bool   // OF, DF, IF, TF
}
