package syscalls

import (
	"fmt"

	"github.com/Bo0km4n/go-minix/kernel/pkg/core/vm/state"
)

func Invoke(s *state.State) error {
	offset := s.BX()
	syscallNum := s.Read16(offset + 2)
	switch syscallNum {
	case 0x0001:
		return nil
	case 0x0004:
		return sysWrite(s, s.Read16(offset+4), s.Read16(offset+10), s.Read16(offset+6))
	}
	return fmt.Errorf("Not found systemcall number: %04x", syscallNum)
}
