package cpu

import (
	"log"

	"gongaware.org/gNES/memory"
)

const stackPage = 0x0100

func (proc *Processor) pushByte(data byte) {
	//Put the SP register ontoan address from the stack page
	pointer := proc.makeStackPointer()
	proc.memory.Write(pointer, data)
	proc.registers[StackPointer]--
}

func (proc *Processor) pushAddress(addr memory.Address) {
	pointer := proc.makeStackPointer()
	proc.memory.WriteAddress(pointer-1, addr)
	proc.registers[StackPointer] -= 2
}

func (proc *Processor) popByte() byte {
	pointer := proc.makeStackPointer()
	//Pointer points to last FREE byte so +1 to get correct location
	result, err := proc.memory.Read(pointer + 1)
	if err != nil {
		log.Fatal(err)
	}
	proc.registers[StackPointer]++
	return result
}

func (proc *Processor) popAddress() memory.Address {
	pointer := proc.makeStackPointer()
	//Pointer points to last FREE byte so +1 to get correct location
	result, err := proc.memory.ReadAddress(pointer + 1)
	if err != nil {
		log.Fatal(err)
	}
	proc.registers[StackPointer] += 2
	return result
}

func (proc Processor) makeStackPointer() memory.Address {
	return stackPage | memory.Address(proc.registers[StackPointer])
}
