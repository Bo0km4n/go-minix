package core

import (
	"bytes"
	"encoding/binary"
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

	// SHORT FORM ENDS HERE
	A_TRSIZE int32
	A_DRSIZE int32
	A_TBASE  int32
	A_DBASE  int32
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
	bbuf := bytes.NewBuffer(buf)
	if err := binary.Read(bbuf, binary.LittleEndian, header); err != nil {
		return err
	}
	if err := assertMagicNumber(header); err != nil {
		return err
	}
	initRelocationHeader(header)

	// load text area
	f.Seek(int64(header.A_HDRLEN), 0)
	textBuf := make([]byte, header.A_TEXT)
	if _, err := f.Read(textBuf); err != nil {
		return err
	}

	// load data area
	dataBuf := make([]byte, header.A_DATA)
	if _, err := f.Read(dataBuf); err != nil {
		return err
	}

	mem := newMemory(textBuf, dataBuf)
	kernel.Memory = mem
	kernel.CPU = newCPU()
	return nil
}

func assertMagicNumber(h *MinixHeader) error {
	if !bytes.Equal(h.A_MAGIC[:], []byte{0x01, 0x03}) {
		return errors.New("Not matched minix header's magic number")
	}
	return nil
}

func initRelocationHeader(h *MinixHeader) {
	if h.A_HDRLEN <= 0x20 {
		h.A_TRSIZE = 0
		h.A_DRSIZE = 0
		h.A_TBASE = 0
		h.A_DBASE = 0
	}
}
