package task

import (
	"fmt"

	"github.com/Bo0km4n/go-minix/kernel/pkg/core/cpu/asem"
	"github.com/Bo0km4n/go-minix/kernel/pkg/core/cpu/state"
	"github.com/Bo0km4n/go-minix/kernel/pkg/core/cpu/syscalls"
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
	if t.state.IP == 0 {
		t.state.PrintParams()
	}
	t.state.PrintRegs()
	t.fetch()
	return t.execAsem()
}

func (t *Task) fetch() {
	t.state.CurInst = t.state.Mem.Text[t.state.IP]
}

func (t *Task) execAsem() error {
	switch t.state.CurInst {
	case 0xbb: // mov
		if err := asem.MOV_Imm_To_Reg(t.state, t.state.CurInst); err != nil {
			return err
		} else {
			t.state.Display.Write(
				[]byte(fmt.Sprintf("%02x%02x%02x\n", t.state.Mem.Text[t.state.IP], t.state.Mem.Text[t.state.IP+1], t.state.Mem.Text[t.state.IP+2])),
			)
			t.state.IP += 3
		}
		return nil
	case 0xcd: // int
		ope := t.state.Mem.Text[t.state.IP+1]
		if ope != 0x20 {
			return fmt.Errorf("Not matched operand: %02x", ope)
		}
		bx := t.state.BX()
		if err := syscalls.Invoke(t.state, bx); err != nil {
			return err
		}
		t.state.Display.Write(
			[]byte(fmt.Sprintf("%02x%02x\n", t.state.Mem.Text[t.state.IP], t.state.Mem.Text[t.state.IP+1])),
		)
		t.state.IP += 2
		return nil
	}
	return fmt.Errorf("Not implemented instruction: %02x", t.state.CurInst)
}
