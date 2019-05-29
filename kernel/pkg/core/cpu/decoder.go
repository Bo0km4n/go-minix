package cpu

import "fmt"

func (c *CPU) fetch() {
	c.curInst = c.mem.Text[c.IP()]
}

func (c *CPU) execAsem() error {
	switch c.curInst {
	case 0xbb:
		if err := MOV_Imm_To_Reg(c, c.curInst); err != nil {
			return err
		} else {
			c.display.Write(
				[]byte(fmt.Sprintf("%02x%02x%02x\n", c.mem.Text[c.ip], c.mem.Text[c.ip+1], c.mem.Text[c.ip+2])),
			)
			c.ip += 3
		}
		return nil
	}
	return fmt.Errorf("Not implemented instruction: %02x", c.curInst)
}
