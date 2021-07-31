package cpu

import "gongaware.org/gNES/memory"

/*
ADC

A = A + operand + carry bit

Sets flags Carry, Zero, Overflow, Negative
*/
func (proc *Processor) AddWithCarry(operand byte) {
	//store A in order to set flags
	oldA := proc.registers[A]

	//Carry bit is in bit 0 so we can just mask it out and add it
	proc.registers[A] += operand + getBits(proc.registers[Status], CarryFlag)

	proc.setCarryFlag(proc.registers[A], oldA)
	proc.setZeroFlag(proc.registers[A])
	proc.setOverflowFlag(proc.registers[A], oldA, operand)
	proc.setNegativeFlag(proc.registers[A])
}

/*
AND

A = A & operand

Sets flags Zero, Negative
*/
func (proc *Processor) And(operand byte) {
	proc.registers[A] &= operand

	proc.setZeroFlag(proc.registers[A])
	proc.setNegativeFlag(proc.registers[A])
}

/*
ASL

out = data << 1 (bit shifted out goes into carry)
*/
func (proc *Processor) ShiftLeft(operand byte) byte {
	result := operand << 1

	proc.setNegativeFlag(result)
	proc.setZeroFlag(result)
	if getBits(operand, NegativeFlag) != 0 {
		proc.setStatusFlags(CarryFlag)
	} else {
		proc.clearStatusFlags(CarryFlag)
	}

	return result
}

/*
BCC

Branch if no carry (CarryFlag = 0)
*/
func (proc *Processor) BranchOnCarryClear(address memory.Address) {
	proc.branchOnFlagClear(address, CarryFlag)
}

/*
BCS

Branch if there is carry (CarryFlag = 1)
*/
func (proc *Processor) BranchOnCarrySet(address memory.Address) {
	proc.branchOnFlagSet(address, CarryFlag)
}

/*
BEQ

Branch if result is zero (ZeroFlag = 1)
*/
func (proc *Processor) BranchOnEqual(address memory.Address) {
	proc.branchOnFlagSet(address, ZeroFlag)
}

/*
BIT Test bits
Bit 6+7 (Negative and overflow) are transferred
The zero flag is then set according to A&operand
*/
func (proc *Processor) BitTest(operand byte) {
	var mask byte = OverflowFlag | NegativeFlag
	//Clear the bits in order to mask the operand bits on
	proc.clearStatusFlags(mask)
	//Set the bits with the operand masked
	proc.setStatusFlags(operand & mask)

	proc.setZeroFlag(proc.registers[A] & operand)
}

/*
BMI
Branch on Minus (NegativeFlag = 1)
*/
func (proc *Processor) BranchOnNegative(address memory.Address) {
	proc.branchOnFlagSet(address, NegativeFlag)
}

/*
BNE

Branch if not Zero (ZeroFlag = 0)
*/
func (proc *Processor) BranchOnNotEqual(address memory.Address) {
	proc.branchOnFlagClear(address, ZeroFlag)
}

/*
BPL

Branch if plus (NegativeFlag = 0)
*/
func (proc *Processor) BranchOnPlus(address memory.Address) {
	proc.branchOnFlagClear(address, NegativeFlag)
}

/*
BRK

Software interrupt.

Push PC+2 to the stack (instruction + interrupt mark). Set break status to 1 and then push to stack
*/
// func (proc *Processor)

//Utility functions
func (proc *Processor) branchOnFlagSet(address memory.Address, flag byte) {
	if proc.registers[Status]&flag == flag {
		proc.pc = address
	}
}

func (proc *Processor) branchOnFlagClear(address memory.Address, flag byte) {
	if proc.registers[Status]&flag == 0 {
		proc.pc = address
	}
}

func (proc *Processor) setZeroFlag(result byte) {
	if result == 0 {
		proc.setStatusFlags(ZeroFlag)
	} else {
		proc.clearStatusFlags(ZeroFlag)
	}
}

func (proc *Processor) setCarryFlag(result byte, original byte) {
	if original > result {
		proc.setStatusFlags(CarryFlag)
	} else {
		proc.clearStatusFlags(CarryFlag)
	}
}

func (proc *Processor) setOverflowFlag(result byte, original byte, operand byte) {
	//If the 2 operands sure a sign but the result doesn't have the same, they overflowed
	if (getBits(original, NegativeFlag) == getBits(operand, NegativeFlag)) && (getBits(original, NegativeFlag) != getBits(result, NegativeFlag)) {
		proc.setStatusFlags(OverflowFlag)
	} else {
		proc.clearStatusFlags(OverflowFlag)
	}
}

func (proc *Processor) setNegativeFlag(result byte) {
	if getBits(result, NegativeFlag) == NegativeFlag {
		proc.setStatusFlags(NegativeFlag)
	} else {
		proc.clearStatusFlags(NegativeFlag)
	}
}
