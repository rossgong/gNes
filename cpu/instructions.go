package cpu

import "gongaware.org/gNES/memory"

/*
ADC

A = A + operand + carry bit

Sets flags Carry, Zero, Overflow, Negative
*/
func (cpu *Processor) AddWithCarry(operand byte) {
	//store A in order to set flags
	oldA := cpu.registers[A]

	//Carry bit is in bit 0 so we can just mask it out and add it
	cpu.registers[A] += operand + getBits(cpu.registers[Status], CarryFlag)

	cpu.setCarryFlag(cpu.registers[A], oldA)
	cpu.setZeroFlag(cpu.registers[A])
	cpu.setOverflowFlag(cpu.registers[A], oldA, operand)
	cpu.setNegativeFlag(cpu.registers[A])
}

/*
AND

A = A & operand

Sets flags Zero, Negative
*/
func (cpu *Processor) And(operand byte) {
	cpu.registers[A] &= operand

	cpu.setZeroFlag(cpu.registers[A])
	cpu.setNegativeFlag(cpu.registers[A])
}

/*
ASL

out = data << 1 (bit shifted out goes into carry)
*/
func (cpu *Processor) ShiftLeft(operand byte) byte {
	result := operand << 1

	cpu.setNegativeFlag(result)
	cpu.setZeroFlag(result)
	if getBits(operand, NegativeFlag) != 0 {
		cpu.setStatusFlags(CarryFlag)
	} else {
		cpu.clearStatusFlags(CarryFlag)
	}

	return result
}

/*
BCC

take branch if carry is 0
*/
func (cpu *Processor) BranchOnCarryClear(address memory.Address) {
	cpu.branchOnFlagClear(address, CarryFlag)
}

/*
BCS

take branch if carry is 0
*/
func (cpu *Processor) BranchOnCarrySet(address memory.Address) {
	cpu.branchOnFlagSet(address, CarryFlag)
}

/*
BEQ

take branch if carry is 0
*/
func (cpu *Processor) BranchOnZero(address memory.Address) {
	cpu.branchOnFlagSet(address, ZeroFlag)
}

//Utility functions
func (cpu *Processor) branchOnFlagSet(address memory.Address, flag byte) {
	if cpu.registers[Status]&flag == flag {
		cpu.pc = address
	}
}

func (cpu *Processor) branchOnFlagClear(address memory.Address, flag byte) {
	if cpu.registers[Status]&flag == 0 {
		cpu.pc = address
	}
}

func (cpu *Processor) setStatusFlags(mask byte) {
	cpu.registers[A] = setBits(cpu.registers[Status], mask)
}

func (cpu *Processor) clearStatusFlags(mask byte) {
	cpu.registers[A] = clearBits(cpu.registers[Status], mask)
}

func (cpu *Processor) setZeroFlag(result byte) {
	if result == 0 {
		cpu.setStatusFlags(ZeroFlag)
	} else {
		cpu.clearStatusFlags(ZeroFlag)
	}
}

func (cpu *Processor) setCarryFlag(result byte, original byte) {
	if original > result {
		cpu.setStatusFlags(CarryFlag)
	} else {
		cpu.clearStatusFlags(CarryFlag)
	}
}

func (cpu *Processor) setOverflowFlag(result byte, original byte, operand byte) {
	//If the 2 operands sure a sign but the result doesn't have the same, they overflowed
	if (getBits(original, NegativeFlag) == getBits(operand, NegativeFlag)) && (getBits(original, NegativeFlag) != getBits(result, NegativeFlag)) {
		cpu.setStatusFlags(OverflowFlag)
	} else {
		cpu.clearStatusFlags(OverflowFlag)
	}
}

func (cpu *Processor) setNegativeFlag(result byte) {
	if getBits(result, NegativeFlag) == NegativeFlag {
		cpu.setStatusFlags(NegativeFlag)
	} else {
		cpu.clearStatusFlags(NegativeFlag)
	}
}

func getBits(b byte, mask byte) byte {
	return b & mask
}

func setBits(b byte, mask byte) byte {
	return b | mask
}

func clearBits(b byte, mask byte) byte {
	return b &^ mask
}
