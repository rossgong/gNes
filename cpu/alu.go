package cpu

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

//Utility Functions
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
