package memory

import (
	"fmt"

	"gongaware.org/gNES/ppu"
)

type Address uint16

const (
	internalRAMSize = 0x800
)

type CPUMap struct {
	internalRAM [internalRAMSize]byte

	ppu *ppu.Processor
}

func (memoryMap CPUMap) Read(address Address) (byte, error) {
	switch {
	case address < 0x2000:
		return memoryMap.internalRAM[address%0x0800], nil
	case address < 0x4000:
		//PPU Registers + mirros
		return memoryMap.ppu.Registers[address%0x8], nil
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
	return LittleEndianToAddress(low, high), err
}

func (memoryMap *CPUMap) Write(address Address, data byte) {
	switch {
	case address < 0x2000:
		memoryMap.internalRAM[address%0x0800] = data
	default:
		notSupported(address, "Unknown Mapping")
	}
}

func (memoryMap *CPUMap) WriteAddress(address Address, data Address) {
	bytes := AddressToLittleEndian(data)
	memoryMap.Write(address, bytes[0])
	memoryMap.Write(address+1, bytes[1])
}

//Utility Functions
func AddressToLittleEndian(address Address) [2]byte {
	return [2]byte{byte(address & 0x00FF), byte(address >> 8)}
}

func LittleEndianToAddress(low byte, high byte) Address {
	return (Address(high) << 8) | Address(low)
}

func notSupported(address Address, extra string) (byte, error) {
	return 0xFF, fmt.Errorf("memory map location 0x%.4x not yet supported (%v)", address, extra)
}
