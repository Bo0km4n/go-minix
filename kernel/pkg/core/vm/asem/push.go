package asem

import (
	"github.com/Bo0km4n/go-minix/kernel/pkg/core/vm/state"
)

func PushReg(s *state.State, inst byte) (int, error) {
	reg := s.GeneralReg16[s.GetReg16Key(inst&maskLow3)]
	addr := s.SP() - 2
	value := reg.GetVal()
	s.Write16(addr, value)
	s.GeneralReg16["SP"].SetVal(addr)
	printInstBytes(s, []byte{inst})
	return 1, nil
}
