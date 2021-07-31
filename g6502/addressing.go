package g6502

import (
	"log"

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
	operand, err := cpu.memory.ReadAddress(operandAddr + memory.Address(offset))
	if err != nil {
		log.Fatal(err)
	}

	val, err := cpu.memory.Read(operand)
	if err != nil {
		log.Fatal(err)
	}
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
	operand, err := cpu.memory.Read(operandAddr)
	if err != nil {
		log.Fatal(err)
	}
	return operand, 1
}

//Indirect Mode
func (cpu Processor) IndirectMode(operandAddr memory.Address) (memory.Address, byte) {
	operand, err := cpu.memory.ReadAddress(operandAddr)
	if err != nil {
		log.Fatal(err)
	}

	data, err := cpu.memory.ReadAddress(operand)
	if err != nil {
		log.Fatal(err)
	}
	return data, 2
}

//Relative Mode
//Very unsafe
func (cpu Processor) RelativeMode(operandAddr memory.Address) (memory.Address, byte) {
	//Convert to int32 in order to preform addition with negative values
	operand, err := cpu.memory.ReadSigned(operandAddr)
	if err != nil {
		log.Fatal(err)
	}
	offset := int32(operand)
	return memory.Address(int32(cpu.pc) + offset), 1
}

//Zero Page Mode
func (cpu Processor) zeroModeOffset(operandAddr memory.Address, offset byte) (byte, byte) {
	operand, err := cpu.memory.Read(operandAddr)
	if err != nil {
		log.Fatal(err)
	}

	zeroPageAddress := memory.Address(operand + offset)
	data, err := cpu.memory.Read(zeroPageAddress)
	if err != nil {
		log.Fatal(err)
	}
	return data, 1
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
	operand, err := cpu.memory.Read(operandAddr)
	if err != nil {
		log.Fatal(err)
	}
	locationOfAddress := operand + cpu.registers[X]
	data, err := cpu.memory.ReadAddress(memory.Address(locationOfAddress))
	if err != nil {
		log.Fatal(err)
	}
	return data, 1
}

func (cpu Processor) ZeroIndirectYMode(operandAddr memory.Address) (memory.Address, byte) {
	operand, err := cpu.memory.Read(operandAddr)
	if err != nil {
		log.Fatal(err)
	}

	data, err := cpu.memory.ReadAddress(memory.Address(operand))
	if err != nil {
		log.Fatal(err)
	}
	return data + memory.Address(cpu.registers[Y]), 1
}
