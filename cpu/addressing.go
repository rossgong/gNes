package cpu

import (
	"gongaware.org/gNES/memory"
)

//operandAddr should be the location of the operand
//modes should return the data being operated on
//Second Returned byte is the amount of bytes the operand takes up

func (cpu Processor) AccumulatorMode(operandAddr memory.Address) (byte, byte) {
	return byte(cpu.registers[A]), 0
}

//Absolute Modes
func (cpu Processor) absoluteOffsetMode(operandAddr memory.Address, offset byte) (byte, byte) {
	address := cpu.memory.ReadAddress(operandAddr + memory.Address(offset))
	val := cpu.memory.Read(address)
	return val, 2
}

func (cpu Processor) AbsoluteMode(operandAddr memory.Address) (byte, byte) {
	return cpu.absoluteOffsetMode(operandAddr, 0)
}

func (cpu Processor) AbsoluteXMode(operandAddr memory.Address) (byte, byte) {
	return cpu.absoluteOffsetMode(operandAddr, cpu.registers[X])
}

func (cpu Processor) AbsoluteYMode(operandAddr memory.Address) (byte, byte) {
	return cpu.absoluteOffsetMode(operandAddr, cpu.registers[Y])
}

//Immediate Mode
func (cpu Processor) ImmediateMode(operandAddr memory.Address) (byte, byte) {
	return cpu.memory.Read(operandAddr), 1
}

//Indirect Mode
func (cpu Processor) IndirectMode(operandAddr memory.Address) (memory.Address, byte) {
	return cpu.memory.ReadAddress(cpu.memory.ReadAddress(operandAddr)), 2
}

//Relative Mode
//Very unsafe
func (cpu Processor) RelativeMode(operandAddr memory.Address) (memory.Address, byte) {
	//Convert to int32 in order to preform addition with negative values
	offset := int32(cpu.memory.ReadSigned(operandAddr))
	return memory.Address(int32(cpu.pc) + offset), 1
}

//Zero Page Mode
func (cpu Processor) zeroModeOffset(operandAddr memory.Address, offset byte) (byte, byte) {
	zeroPageAddress := memory.Address(cpu.memory.Read(operandAddr) + offset)
	return cpu.memory.Read(zeroPageAddress), 1
}

func (cpu Processor) ZeroMode(operandAddr memory.Address) (byte, byte) {
	return cpu.zeroModeOffset(operandAddr, 0)
}

func (cpu Processor) ZeroXMode(operandAddr memory.Address) (byte, byte) {
	return cpu.zeroModeOffset(operandAddr, cpu.registers[X])
}

func (cpu Processor) ZeroYMode(operandAddr memory.Address) (byte, byte) {
	return cpu.zeroModeOffset(operandAddr, cpu.registers[Y])
}

func (cpu Processor) ZeroIndirectXMode(operandAddr memory.Address) (memory.Address, byte) {
	locationOfAddress := cpu.memory.Read(operandAddr) + cpu.registers[X]
	return cpu.memory.ReadAddress(memory.Address(locationOfAddress)), 1
}

func (cpu Processor) ZeroIndirectYMode(operandAddr memory.Address) (memory.Address, byte) {
	locationOfAddress := cpu.memory.Read(operandAddr)
	return cpu.memory.ReadAddress(memory.Address(locationOfAddress)) + memory.Address(cpu.registers[Y]), 1
}
