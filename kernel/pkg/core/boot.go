package core

import (
	"github.com/Bo0km4n/go-minix/kernel/pkg/core/kernel"
)

func Boot(filename string) error {
	if err := initKernel(filename); err != nil {
		return err
	}
	kernel.K.Run()
	return nil
}
