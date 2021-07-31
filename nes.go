package gnes

import (
	"gongaware.org/gNES/g6502"
	"gongaware.org/gNES/memory"
)

type System struct {
	cpu       g6502.Processor
	cpuMemory memory.CPUMap
}
