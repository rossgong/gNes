package cpu

import (
	"gongaware.org/gNES/memory"
)

//Interrupt Vector Table
//These are the memory addresses of the handler addresses
const (
	NMIVector    = 0xFFFA //+0xFFFB
	ResetVector  = 0xFFFC //+0xFFFD
	IRQBRKVector = 0xFFFE //+0xFFFF
)

type Processor struct {
	registers [numRegisters]byte
	pc        memory.Address

	memory *memory.CPUMap
}

func (proc *Processor) InitializeRegisters() {
	//SP initally points to the first free byte on the first page.
	//SP is decrmented so this would be 0xFF (0x01FF)
	proc.registers[StackPointer] = 0xFF

	//Bit 5 is not used and always 1
	proc.registers[Status] = 0b0010_0000
}
