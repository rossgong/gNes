package g6502

import (
	"gongaware.org/gNES/memory"
)

const (
	//Bit 5 is not used and always 1
	//Even though the InterruptFlag is set on reset, keep that in the reset
	initialStatus = 0b0010_0000
)

type Processor struct {
	registers [numRegisters]byte
	pc        memory.Address

	memory *memory.CPUMap
}

func (proc *Processor) InitializeRegisters() {
	proc.registers[Status] = initialStatus
}