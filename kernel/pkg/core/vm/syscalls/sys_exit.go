package syscalls

import (
	"fmt"

	"github.com/Bo0km4n/go-minix/kernel/pkg/core/config"
	"github.com/Bo0km4n/go-minix/kernel/pkg/core/vm/state"
)

func sysExit(s *state.State) error {
	if config.Trace {
		fmt.Println("<exit(0)>")
	}
	s.HasExit = true
	return nil
}
