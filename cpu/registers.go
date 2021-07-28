package cpu

//Register labels
const (
	A = iota
	X
	Y
	StackPointer
	Status
	//Number of normal CPU registers (excludes PC)
	numRegisters
)
