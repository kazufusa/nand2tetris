package computer

import (
	"testing"

	logic "github.com/kazufusa/nand2tetris/01_Boolean_Logic"
	memory "github.com/kazufusa/nand2tetris/03_Memory"
	"github.com/stretchr/testify/assert"
)

var (
	// 0    000 0 000000 000 000 // @0
	// 1    111 1 110000 010 000 // D=M
	// 2    000 0 000000 000 001 // @1
	// 3    111 1 010011 010 000 // D=D-M
	// 4    000 0 000000 001 010 // @10
	// 5    111 0 001100 000 001 // D;JGT
	// 6    000 0 000000 000 001 // @1
	// 7    111 1 110000 010 000 // D=M
	// 8    000 0 000000 001 100 // @12
	// 9    111 0 101010 000 111 // 0;JMP
	// 10   000 0 000000 000 000 // @0
	// 11   111 1 110000 010 000 // D=M
	// 12   000 0 000000 000 010 // @2
	// 13   111 0 001100 001 000 // M=D
	// 14   000 0 000000 001 110 // @14
	// 15   111 0 101010 000 111 // 0;JMP
	maxInstructions = []Word{
		{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 1, 0, 0, 0, 0, 0, 1, 1, 1, 1, 1, 1},
		{1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 1, 0, 1, 1, 0, 0, 1, 0, 1, 1, 1, 1},
		{0, 1, 0, 1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
		{1, 0, 0, 0, 0, 0, 0, 0, 1, 1, 0, 0, 0, 1, 1, 1},
		{1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 1, 0, 0, 0, 0, 0, 1, 1, 1, 1, 1, 1},
		{0, 0, 1, 1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
		{1, 1, 1, 0, 0, 0, 0, 1, 0, 1, 0, 1, 0, 1, 1, 1},
		{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 1, 0, 0, 0, 0, 0, 1, 1, 1, 1, 1, 1},
		{0, 1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 1, 0, 0, 0, 0, 1, 1, 0, 0, 0, 1, 1, 1},
		{0, 1, 1, 1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
		{1, 1, 1, 0, 0, 0, 0, 1, 0, 1, 0, 1, 0, 1, 1, 1},
	}
)

func TestComputer(t *testing.T) {
	clock := memory.Clock(0)

	sc := NewTestScreen(&clock)
	kb := TestKeyboard{}
	ram := NewMemory(&clock, &sc, &kb)

	rom := NewROM32K()
	rom.BulkLoad(maxInstructions)

	cpu := NewCPU()

	com := NewComputer(&cpu, &ram, &rom, &clock)

	addr0 := [15]logic.Bit{}
	addr1 := [15]logic.Bit{1}
	addr2 := [15]logic.Bit{0, 1}
	w0 := Word{}
	w6 := Word{0, 1, 1}
	w10 := Word{0, 1, 0, 1}

	com.ram.Fetch(w10, logic.I, addr0)
	com.clock.Progress()
	com.ram.Fetch(w6, logic.I, addr1)
	com.clock.Progress()
	for i := 0; i < 20; i++ {
		com.FetchAndExecute(logic.O)
	}
	assert.Equal(t, w10, com.ram.Fetch(w0, logic.O, addr2))

	com.FetchAndExecute(logic.I)
	com.ram.Fetch(w6, logic.I, addr0)
	com.clock.Progress()
	com.ram.Fetch(w10, logic.I, addr1)
	com.clock.Progress()
	for i := 0; i < 20; i++ {
		com.FetchAndExecute(logic.O)
	}
	assert.Equal(t, w10, com.ram.Fetch(w0, logic.O, addr2))
}
