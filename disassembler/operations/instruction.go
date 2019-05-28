package operations

// Instruction inst asset
type Instruction interface {
	Analyze(*Context, byte) (int, string)
}

var (
	// Instructions map
	Instructions = map[byte]Instruction{}
)
