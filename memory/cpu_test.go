package memory

import (
	"testing"
)

func initializeTestCPUMemory() CPUMap {
	memory := CPUMap{}
	for i := range memory.internalRAM {
		memory.internalRAM[i] = byte(i % 0x100)
	}
	return memory
}

func TestCPUMapRead(t *testing.T) {
	type test struct {
		addr Address
		want byte
	}

	tests := []test{
		{0x0010, 0x10},
		{0x1110, 0x10},
	}

	memory := initializeTestCPUMemory()
	for _, conditions := range tests {
		data, err := memory.Read(conditions.addr)
		if err != nil {
			t.Errorf(err.Error())
		}
		if data != conditions.want {
			t.Fatalf("expected 0x%.2x at address 0x%.4x got 0x%.2x", conditions.want, conditions.addr, data)
		}
	}
}
