package cpu

import (
	"log"

	"gongaware.org/gNES/memory"
)

//Interrupt Vector Table
//These are the memory addresses of the handler addresses
const (
	NMIVector    = 0xFFFA //+0xFFFB
	ResetVector  = 0xFFFC //+0xFFFD
	IRQBRKVector = 0xFFFE //+0xFFFF
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

func (proc *Processor) Reset() {
	proc.registers[Status] = setBits(proc.registers[Status], InterruptDisableFlag)
	addr, err := proc.memory.ReadAddress(ResetVector)
	if err != nil {
		log.Fatal(err)
	}
	proc.pc = addr
}
