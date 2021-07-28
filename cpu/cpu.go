package cpu

import (
	"gongaware.org/gNES/memory"
)

type Processor struct {
	registers [numRegisters]byte
	pc        memory.Address

	memory *memory.CPUMap
}
