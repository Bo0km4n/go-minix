package operations

import (
	"bytes"
	"encoding/binary"
	"fmt"
)

// Binary model
// minix binary model
// TextSize: sorted littel endian
type Binary struct {
	Header      []byte
	Body        []byte
	TextSize    []byte
	TextSizeInt int32
}

// NewBinary function
// read file binary
func NewBinary(f []byte) *Binary {
	var b Binary
	var s int32
	b.Header = make([]byte, 32)
	b.TextSize = make([]byte, 4)

	// read header
	for i := 0; i < 32; i++ {
		b.Header[i] = f[i]
	}
	// read text size by little endian
	for i, v := range f[8:12] {
		b.TextSize[i] = v
	}
	// convert text size to int
	buf := bytes.NewReader(b.TextSize)
	if err := binary.Read(buf, binary.LittleEndian, &s); err != nil {
		fmt.Println("binary read failed:", err)
		return nil
	}
	b.TextSizeInt = s
	// read body
	for i := 32; i < 32+int(b.TextSizeInt); i++ {
		b.Body = append(b.Body, f[i])
	}
	return &b
}
