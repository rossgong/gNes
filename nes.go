package gnes

import (
	"gongaware.org/gNES/cpu"
	"gongaware.org/gNES/memory"
)

type System struct {
	cpu       cpu.Processor
	cpuMemory memory.CPUMap
}
