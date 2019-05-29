package core

import (
	"github.com/Bo0km4n/go-minix/kernel/pkg/core/kernel"
	"github.com/k0kubun/pp"
)

func Run(filename string) error {
	if err := loadBin(filename); err != nil {
		return err
	}
	pp.Println([]byte(filename))
	kernel.K.Run()
	return nil
}
