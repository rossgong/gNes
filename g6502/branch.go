package g6502

import "gongaware.org/gNES/memory"

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
