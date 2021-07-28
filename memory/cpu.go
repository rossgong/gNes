package memory

import (
	"fmt"
)

type Address uint16

const (
	internalRAMSize = 0x800
)

type CPUMap struct {
	internalRAM [internalRAMSize]byte
}

func (memoryMap CPUMap) Read(address Address) (byte, error) {
	switch {
	case address < 0x2000:
		return memoryMap.internalRAM[address%0x0800], nil
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

func (memoryMap CPUMap) ReadSigned(address Address) (int8, error) {
	data, err := memoryMap.Read(address)
	return int8(data), err
}

//Little endian so least sig byte comes first
func (memoryMap CPUMap) ReadAddress(operandAddr Address) (Address, error) {
	low, err := memoryMap.Read(operandAddr)
	if err != nil {
		return Address(low), err
	}

	high, err := memoryMap.Read(operandAddr + 1)
	if err != nil {
		return Address(high), err
	}
	return (Address(high) << 8) + Address(low), err
}

func (memoryMap *CPUMap) Write(address Address, data byte) {
	switch {
	case address < 0x2000:
		memoryMap.internalRAM[address%0x0800] = data
	default:
		notSupported(address, "Unknown Mapping")
	}
}

func notSupported(address Address, extra string) (byte, error) {
	return 0xFF, fmt.Errorf("memory map location 0x%.4x not yet supported (%v)", address, extra)
}
