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
