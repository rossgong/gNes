package g6502

import (
	"log"

	"gongaware.org/gNES/memory"
)

//Address of interrupt handlers
const (
	NMIVector =  0xFFFA
	ResetVector = 0xFFFC
	IRQVector = 0xFFFE
	BreakVector = IRQVector
)
const (
	NMI = iota
	Reset
	IRQ
	Break
)

/*
BRK

Software interrupt.

Push PC+2 to the stack (instruction + interrupt mark). Set break status to 1 and then push to stack
*/
func (proc *Processor) Break() {
	if proc.interruptsMasked() {
		return
	}

	proc.setStatusFlags(Break)
}

func (proc *Processor) Reset() {
	//pc handled here
	proc.interrupt(ResetVector)
}

func (proc *Processor) ReturnFromInterrupt() {
	proc.registers[Status] = proc.popByte()
	proc.pc = proc.popAddress()
}

//Utility Functions
func (proc Processor) interruptsMasked() bool {
	return proc.registers[Status] & InterruptDisableFlag > 0
}

func (proc *Processor) interrupt(interruptType int) {
	var vector memory.Address
	switch interruptType {
	case NMI:
		vector = NMIVector
	case Reset:
		vector = ResetVector
	case IRQ:
		vector = IRQVector
	case Break:
		vector = BreakVector
	default:
		log.Fatal("interrupts: Illegal interrupt type")
	}

	if (interruptType != ResetVector) {
		proc.pushAddress(proc.pc+2)
		proc.pushByte(proc.registers[Status])
		proc.setStatusFlags(InterruptDisableFlag)
	}

	handler, err := proc.memory.ReadAddress(vector)
	if err != nil {
		log.Fatal(err)
	}
	proc.pc = handler
}