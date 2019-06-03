package task

import (
	"fmt"
	"os"

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
	if t.state.IP == 0 && config.Trace && !t.state.HasExit {
		t.state.PrintParams()
	}
	if config.Trace && !t.state.HasExit {
		t.state.PrintRegs()
	}
	t.fetch()
	return t.execAsem()
}

func (t *Task) fetch() {
	t.state.CurInst = t.state.Mem.Text[t.state.IP]
}

func (t *Task) execAsem() error {
	if t.state.HasExit {
		os.Exit(0)
	}
	switch t.state.CurInst {
	case 0xbb: // mov
		if config.Trace {
			t.state.Display.Write(
				[]byte(fmt.Sprintf("%02x%02x%02x\n", t.state.Mem.Text[t.state.IP], t.state.Mem.Text[t.state.IP+1], t.state.Mem.Text[t.state.IP+2])),
			)
		}
		if err := asem.MOV_Imm_To_Reg(t.state, t.state.CurInst); err != nil {
			return err
		}
		t.state.IP += 3
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
