package computer

import (
	"strings"

	logic "github.com/kazufusa/nand2tetris/01_Boolean_Logic"
	memory "github.com/kazufusa/nand2tetris/03_Memory"
)

func NewEmulator(inst interface{}) *Computer {
	clock := memory.Clock(0)
	ram := VMemory{}
	rom := VROM32K{}
	switch x := inst.(type) {
	case [][16]uint8:
		rom.BulkLoad(x)
	case string:
		instLines := strings.Split(x, "\n")
		var _inst []Word
		for _, l := range instLines {
			if len(l) > 0 {
				_inst = append(_inst, string2Word(l))
			}
		}
		rom.BulkLoad(_inst)
	}
	cpu := NewCPU()
	return &Computer{cpu: &cpu, ram: &ram, rom: &rom, clock: &clock}
}

func (com *Computer) WriteRom(addr, value int) {
	switch x := com.ram.(type) {
	case *VMemory:
		if x.values == nil {
			x.values = make(map[int]Word)
		}
		x.values[addr] = int2Word(value)
	}
}

func (com *Computer) ROMMap() (ret map[int]int) {
	ret = make(map[int]int)
	switch x := com.ram.(type) {
	case *VMemory:
		for k, v := range x.values {
			ret[k] = word2Int(v)
		}
	}
	return ret
}

type VMemory struct {
	values map[int]Word
}

var _ IMemory = (*Memory)(nil)

type VROM32K struct {
	values map[int]Word
}

var _ IROM32K = (*VROM32K)(nil)

func (m *VMemory) Fetch(in Word, load Bit, addr [15]Bit) (out Word) {
	defer func() {
		if load == logic.I {
			m.values[addr2int(addr)] = in
		}
	}()
	if m.values == nil {
		m.values = make(map[int]Word)
	}
	return m.values[addr2int(addr)]
}

func (m *VROM32K) Fetch(addr [15]Bit) (out Word) {
	if m.values == nil {
		m.values = make(map[int]Word)
	}
	return m.values[addr2int(addr)]
}

func (m *VROM32K) BulkLoad(ws []Word) {
	if m.values == nil {
		m.values = make(map[int]Word)
	}
	for i, w := range ws {
		m.values[i] = w
	}
}

func addr2int(addr [15]Bit) int {
	return (B0)*int(addr[0]) +
		(B1)*int(addr[1]) +
		(B2)*int(addr[2]) +
		(B3)*int(addr[3]) +
		(B4)*int(addr[4]) +
		(B5)*int(addr[5]) +
		(B6)*int(addr[6]) +
		(B7)*int(addr[7]) +
		(B8)*int(addr[8]) +
		(B9)*int(addr[9]) +
		(B10)*int(addr[10]) +
		(B11)*int(addr[11]) +
		(B12)*int(addr[12]) +
		(B13)*int(addr[13]) +
		(B14)*int(addr[14])
}

func word2Int(w Word) int {
	ret := (B0)*int(w[0]) +
		(B1)*int(w[1]) +
		(B2)*int(w[2]) +
		(B3)*int(w[3]) +
		(B4)*int(w[4]) +
		(B5)*int(w[5]) +
		(B6)*int(w[6]) +
		(B7)*int(w[7]) +
		(B8)*int(w[8]) +
		(B9)*int(w[9]) +
		(B10)*int(w[10]) +
		(B11)*int(w[11]) +
		(B12)*int(w[12]) +
		(B13)*int(w[13]) +
		(B14)*int(w[14])
	if w[15] == logic.O {
		return ret
	} else {
		return ret - 32768
	}
}

func int2Word(v int) Word {
	return Word{
		uint8(v & B0 >> 0),
		uint8(v & B1 >> 1),
		uint8(v & B2 >> 2),
		uint8(v & B3 >> 3),
		uint8(v & B4 >> 4),
		uint8(v & B5 >> 5),
		uint8(v & B6 >> 6),
		uint8(v & B7 >> 7),
		uint8(v & B8 >> 8),
		uint8(v & B9 >> 9),
		uint8(v & B10 >> 10),
		uint8(v & B11 >> 11),
		uint8(v & B12 >> 12),
		uint8(v & B13 >> 13),
		uint8(v & B14 >> 14),
		uint8(v & B15 >> 15),
	}
}

func string2Word(b string) Word {
	return Word{
		b[15] - 48,
		b[14] - 48,
		b[13] - 48,
		b[12] - 48,
		b[11] - 48,
		b[10] - 48,
		b[9] - 48,
		b[8] - 48,
		b[7] - 48,
		b[6] - 48,
		b[5] - 48,
		b[4] - 48,
		b[3] - 48,
		b[2] - 48,
		b[1] - 48,
		b[0] - 48,
	}
}
