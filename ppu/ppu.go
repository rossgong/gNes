package ppu

const (
	Ctrl = iota
	Mask
	Status
	OamAddr
	OamData
	Scroll
	Address
	Data
	OamDMA
	lengthOfRegisters
)

type Processor struct {
	Registers [lengthOfRegisters]byte
}
