package syscalls

import (
	"fmt"
	"syscall"

	"github.com/Bo0km4n/go-minix/kernel/pkg/core/config"
	"github.com/Bo0km4n/go-minix/kernel/pkg/core/vm/state"
)

func sysWrite(s *state.State, fd, buf, length uint16) error {
	if config.Trace {
		fmt.Printf("<write>(%d, %04x, %d) => %s", int(fd), buf, int(length), string(s.Mem.Data[buf:buf+length]))
		return nil
	}
	n, err := syscall.Write(int(fd), s.Mem.Data[buf:buf+length])
	if err != nil {
		return err
	}
	if n != int(length) {
		fmt.Errorf("syscall write: length(%d) != n(%d)", length, n)
	}
	return nil
}
