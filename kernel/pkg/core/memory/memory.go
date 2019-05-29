package memory

func NewMemory(textArea, dataArea []byte) *Memory {
	return &Memory{
		Text: textArea,
		Data: dataArea,
	}
}

// Memory has some byte areas
type Memory struct {
	Text []byte
	Data []byte
}
