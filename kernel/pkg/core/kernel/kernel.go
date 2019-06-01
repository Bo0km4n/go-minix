package kernel

import (
	"log"

	"github.com/Bo0km4n/go-minix/kernel/pkg/core/memory"
	"github.com/Bo0km4n/go-minix/kernel/pkg/core/vm/task"
)

var K *Kernel

// Kernel is core structure
type Kernel struct {
	Memory *memory.Memory
	Task   *task.Task
	Args   []string
	Envs   []string
}

func (k *Kernel) Run() {
	k.Task.SetArgs(k.Args, k.Envs)
	for {
		if err := k.Task.Exec(); err != nil {
			log.Fatalf("\n%v", err)
		}
	}
}
