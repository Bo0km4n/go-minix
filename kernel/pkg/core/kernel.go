package core

var K *Kernel

type Kernel struct {
	Memory    *Memory
	CPU       *CPU
	Registers *Registers
}

type Memory struct {
	Text []byte
	Data []byte
}

type CPU uint16

type Registers struct {
	General map[string]uint16
	Flag    map[string]bool
}
