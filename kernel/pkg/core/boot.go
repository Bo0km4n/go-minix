package core

import (
	"github.com/Bo0km4n/go-minix/kernel/pkg/core/config"
	"github.com/Bo0km4n/go-minix/kernel/pkg/core/kernel"
)

func Boot(filename string, trace bool) error {
	config.Trace = trace
	if err := initKernel(filename); err != nil {
		return err
	}
	kernel.K.Run()
	return nil
}
