package core

import "github.com/Bo0km4n/go-minix/kernel/pkg/core/kernel"

func Run(filename string) error {
	if err := loadBin(filename); err != nil {
		return err
	}
	kernel.K.Run()
	return nil
}
