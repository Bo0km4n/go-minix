package kernel

import (
	"log"

	"github.com/Bo0km4n/go-minix/kernel/pkg/core/vm/task"
	"github.com/Bo0km4n/go-minix/kernel/pkg/core/memory"
)

var K *Kernel

// Kernel is core structure
type Kernel struct {
	Memory *memory.Memory
	Task   *task.Task
}

func (k *Kernel) Run() {
	for {
		if err := k.Task.Exec(); err != nil {
			log.Fatalf("\n%v", err)
		}
	}
}
