package asem

import "github.com/Bo0km4n/go-minix/kernel/pkg/core/vm/state"

func CallDisp(s *state.State, inst byte) (int, error) {
	// push return address to stack
	retAddr := s.IP + 3
	s.PushToStack(uint16(retAddr))
	next := s.ReadTextU16(uint16(s.IP + 1))
	printInstBytes(s, s.Mem.Text[s.IP:s.IP+3])
	return int(3 + next), nil
}
