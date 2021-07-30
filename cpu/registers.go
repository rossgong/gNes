package cpu

//Register labels
const (
	A = iota
	X
	Y
	StackPointer
	/*
		Status holds various flags set by instructions.

		Bit	Meaning

		0	Carry (1 means carry)

		1	Zero (1 means zero)

		2	IRQ (1 means disable)

		3	Decimal Mode (1 means on)

		4	Break (1 means Reset)

		5	UNUSED (ALWAYS 1)

		6	Overflow (1 means overflow)

		7	Negative (1 means negative)
	*/
	Status
	//Number of normal CPU registers (excludes PC)
	numRegisters
)

/*
	Status Flags

	Gives the appropriate bit number for the status flag
*/
const (
	//0	Carry flag (1 means carry)
	carry = iota
	//1	Zero flag (1 means zero)
	zero
	//2	IRQ flag (1 means disable)
	irq
	//3	Decimal flag Mode (1 means on)
	decimal
	//4	Break flag (1 means Reset)
	reset
	//5	UNUSED flag (ALWAYS 1)
	unused
	//6	Overflow flag (1 means overflow)
	overflow
	//7	Negative flag (1 means negative)
	negative
)
