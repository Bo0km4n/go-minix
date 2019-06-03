package memory

func NewMemory(textArea, dataArea []byte) *Memory {
	data := make([]byte, 0x10000)
	copy(data[0:len(dataArea)], dataArea[:])
	return &Memory{
		Text: textArea,
		Data: data,
	}
}

// Memory has some byte areas
type Memory struct {
	Text []byte
	Data []byte
}
