package core

import (
	"log"

	"github.com/Bo0km4n/go-minix/kernel/pkg/core/cpu"
	"github.com/Bo0km4n/go-minix/kernel/pkg/core/memory"
)

var K *Kernel

// Kernel is core structure
type Kernel struct {
	Memory *memory.Memory
	CPU    *cpu.CPU
}

func (k *Kernel) exec() {
	for {
		if err := k.CPU.Exec(); err != nil {
			log.Fatalf("\n%v", err)
		}
	}
}
