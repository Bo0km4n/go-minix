package core

import (
	"os"
)

func Run(filename string) error {
	if err := loadBin(filename); err != nil {
		return err
	}
	K.PrintParams(os.Stdin)
	K.PrintRegs(os.Stdin)
	return nil
}
