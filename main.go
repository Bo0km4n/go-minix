package main

import (
	"fmt"
	"io/ioutil"
	"os"

	"github.com/Bo0km4n/go-minix/operations"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("not enough fila name arguments")
		return
	}
	contents, err := ioutil.ReadFile(os.Args[1])
	if err != nil {
		fmt.Println("can not find file")
		return
	}
	b := operations.NewBinary(contents)
	ctx := operations.Context{}
	ctx.Disassemble(b.Body)
}
