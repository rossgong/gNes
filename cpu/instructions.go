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

Branch if no carry (CarryFlag = 0)
*/
func (cpu *Processor) BranchOnCarryClear(address memory.Address) {
	cpu.branchOnFlagClear(address, CarryFlag)
}

/*
BCS

Branch if there is carry (CarryFlag = 1)
*/
func (cpu *Processor) BranchOnCarrySet(address memory.Address) {
	cpu.branchOnFlagSet(address, CarryFlag)
}

/*
BEQ

Branch if result is zero (ZeroFlag = 1)
*/
func (cpu *Processor) BranchOnEqual(address memory.Address) {
	cpu.branchOnFlagSet(address, ZeroFlag)
}

/*
BIT Test bits
Bit 6+7 (Negative and overflow) are transferred
The zero flag is then set according to A&operand
*/
func (cpu *Processor) BitTest(operand byte) {
	var mask byte = OverflowFlag | NegativeFlag
	//Clear the bits in order to mask the operand bits on
	cpu.clearStatusFlags(mask)
	//Set the bits with the operand masked
	cpu.setStatusFlags(operand & mask)

	cpu.setZeroFlag(cpu.registers[A] & operand)
}

/*
BMI
Branch on Minus (NegativeFlag = 1)
*/
func (cpu *Processor) BranchOnNegative(address memory.Address) {
	cpu.branchOnFlagSet(address, NegativeFlag)
}

/*
BNE

Branch if not Zero (ZeroFlag = 0)
*/
func (cpu *Processor) BranchOnNotEqual(address memory.Address) {
	cpu.branchOnFlagClear(address, ZeroFlag)
}

/*
BPL

Branch if plus (NegativeFlag = 0)
*/
func (cpu *Processor) BranchOnPlus(address memory.Address) {
	cpu.branchOnFlagClear(address, NegativeFlag)
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
