package task

import (
	"fmt"

	"github.com/Bo0km4n/go-minix/kernel/pkg/core/config"
	"github.com/Bo0km4n/go-minix/kernel/pkg/core/vm/asem"
	"github.com/Bo0km4n/go-minix/kernel/pkg/core/vm/state"
	"github.com/Bo0km4n/go-minix/kernel/pkg/core/vm/syscalls"
)

type Task struct {
	state *state.State
}

func NewTask(s *state.State) *Task {
	return &Task{
		state: s,
	}
}

func (t *Task) Exec() error {
	if t.state.IP == 0 && config.Trace {
		t.state.PrintParams()
	}
	if config.Trace {
		t.state.PrintRegs()
	}
	t.fetch()
	return t.execAsem()
}

func (t *Task) fetch() {
	t.state.CurInst = t.state.Mem.Text[t.state.IP]
}

func (t *Task) execAsem() error {
	switch t.state.CurInst {
	case 0xbb: // mov
		n, err := asem.MovImmToReg(t.state, t.state.CurInst)
		if err != nil {
			return err
		}
		t.state.IP += int32(n)
		return nil
	case 0x88, 0x89, 0x8a, 0x8b: // MOV Register/Memory to /from Register
		n, err := asem.MovRmToRm(t.state, t.state.CurInst)
		if err != nil {
			return err
		}
		t.state.IP += int32(n)
		return nil

	case 0x50, 0x51, 0x52, 0x53, 0x54, 0x55, 0x56, 0x57: // PUSH Register
		n, err := asem.PushReg(t.state, t.state.CurInst)
		if err != nil {
			return err
		}
		t.state.IP += int32(n)
		return nil
	case 0xcd: // int
		if config.Trace {
			t.state.Display.Write(
				[]byte(fmt.Sprintf("%02x%02x\n", t.state.Mem.Text[t.state.IP], t.state.Mem.Text[t.state.IP+1])),
			)
		}
		ope := t.state.Mem.Text[t.state.IP+1]
		if ope != 0x20 {
			return fmt.Errorf("Not matched operand: %02x", ope)
		}
		if err := syscalls.Invoke(t.state); err != nil {
			return err
		}

		t.state.IP += 2
		return nil
	}
	return fmt.Errorf("Not implemented instruction: %02x", t.state.CurInst)
}

func (t *Task) SetArgs(args, envs []string) {
	t.state.SetArgs(args, envs)
}
