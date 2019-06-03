package syscalls

import (
	"fmt"
	"syscall"

	"github.com/Bo0km4n/go-minix/kernel/pkg/core/config"
	"github.com/Bo0km4n/go-minix/kernel/pkg/core/vm/state"
)

func sysExit(s *state.State) error {
	if config.Trace {
		fmt.Println("<exit(0)>")
	}
	syscall.Exit(0)
	return nil
}
