package core

import (
	"bytes"
	"errors"
	"os"
	"unsafe"
)

type MinixHeader struct {
	A_MAGIC   [2]byte
	A_FLAGS   byte
	A_CPU     byte
	A_HDRLEN  byte
	A_UNUSED  byte
	A_VERSION int16
	A_TEXT    int32
	A_DATA    int32
	A_BSS     int32
	A_ENTRY   int32
	A_TOTAL   int32
	A_SYMS    int32
}

func loadBin(filename string) error {
	f, err := os.Open(filename)
	if err != nil {
		return err
	}
	newKernel := &Kernel{}
	if err := allocate(f, newKernel); err != nil {
		return err
	}
	K = newKernel
	return nil
}

func allocate(f *os.File, kernel *Kernel) error {
	// parse header
	header := &MinixHeader{}
	size := unsafe.Sizeof(*header)
	buf := make([]byte, size)
	if n, err := f.Read(buf); err != nil || n != int(size) {
		return err
	}
	h := *(*MinixHeader)(unsafe.Pointer(&buf[0]))
	if err := assertMagicNumber(&h); err != nil {
		return err
	}
	return nil
}

func assertMagicNumber(h *MinixHeader) error {
	if !bytes.Equal(h.A_MAGIC[:], []byte{0x01, 0x03}) {
		return errors.New("Not matched minix header's magic number")
	}
	return nil
}
