package gnes

import (
	"gongaware.org/gNES/cpu"
	"gongaware.org/gNES/ram"
)

type (
	Address uint16
)

type System struct {
	cpu cpu.Processor
	ram ram.Memory
}
