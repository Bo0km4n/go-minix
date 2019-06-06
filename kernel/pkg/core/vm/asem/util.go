package asem

import (
	"bytes"
	"fmt"

	"github.com/Bo0km4n/go-minix/kernel/pkg/core/config"
	"github.com/Bo0km4n/go-minix/kernel/pkg/core/vm/state"
)

var (
	maskMid3 = byte(0x38)
	maskLow3 = byte(0x07)
	maskTop2 = byte(0xc0)
)

type regKeyFunc func(byte) string

func printInstBytes(s *state.State, b []byte) {
	if !config.Trace {
		return
	}
	buf := &bytes.Buffer{}
	for i := range b {
		buf.WriteString(fmt.Sprintf("%02x", b[i]))
	}
	buf.WriteString("\n")
	s.Display.Write(buf.Bytes())
}
