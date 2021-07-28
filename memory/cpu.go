package memory

import (
	"log"
)

type Address uint16

const (
	internalRAMSize = 0x800
)

type CPUMap struct {
	internalRAM [internalRAMSize]byte
}

func (memoryMap CPUMap) Read(address Address) byte {
	switch {
	case address < 0x2000:
		return memoryMap.internalRAM[address%0x0800]
	case address < 0x4000:
		//PPU Registers + mirros
		return notSupported(address, "PPU Registers")
	case address < 0x4018:
		//APU/IO Registers
		return notSupported(address, "APU Registers")
	case address < 0x4020:
		return notSupported(address, "Disabled")
	default:
		return notSupported(address, "Cartridge")
	}
}

func (memoryMap CPUMap) ReadSigned(address Address) int8 {
	return int8(memoryMap.Read(address))
}

//Little endian so least sig byte comes first
func (memoryMap CPUMap) ReadAddress(operandAddr Address) Address {
	low := memoryMap.Read(operandAddr)
	high := memoryMap.Read(operandAddr + 1)
	return (Address(high) << 8) + Address(low)
}

func (memoryMap *CPUMap) Write(address Address, data byte) {
	switch {
	case address < 0x2000:
		memoryMap.internalRAM[address%0x0800] = data
	default:
		notSupported(address, "Unknown Mapping")
	}
}

func notSupported(address Address, extra string) byte {
	log.Printf("Memory map location 0x%.4x not yet supported (%v)", address, extra)
	return 0xFF
}
