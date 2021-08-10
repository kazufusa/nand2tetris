package computer

import (
	"fmt"
	"reflect"
	"testing"

	logic "github.com/kazufusa/nand2tetris/01_Boolean_Logic"
	"github.com/stretchr/testify/assert"
)

func TestCPUDest(t *testing.T) {
	var tests = []struct {
		expectedA      Word
		expectedD      Word
		expectedM      Word
		expectedWriteM logic.Bit
		given          []logic.Bit
	}{
		{Word{}, Word{}, Word{}, logic.O, []logic.Bit{0, 0, 0}},
		{Word{0}, Word{0}, Word{1}, logic.I, []logic.Bit{1, 0, 0}},
		{Word{0}, Word{1}, Word{0}, logic.O, []logic.Bit{0, 1, 0}},
		{Word{0}, Word{1}, Word{1}, logic.I, []logic.Bit{1, 1, 0}},
		{Word{1}, Word{0}, Word{0}, logic.O, []logic.Bit{0, 0, 1}},
		{Word{1}, Word{0}, Word{1}, logic.I, []logic.Bit{1, 0, 1}},
		{Word{1}, Word{1}, Word{0}, logic.O, []logic.Bit{0, 1, 1}},
		{Word{1}, Word{1}, Word{1}, logic.I, []logic.Bit{1, 1, 1}},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(fmt.Sprintf("%v", tt.given), func(t *testing.T) {
			cpu := NewCPU()
			outM, writeM, _, _ := cpu.Fetch(
				Word{},
				Word{0, 0, 0, tt.given[0], tt.given[1], tt.given[2], 1, 1, 1, 1, 1, 1, 0, 1, 1, 1},
				logic.O,
			)
			assert.Equal(t, tt.expectedA, cpu.a.Apply(logic.O, Word{}), "A")
			assert.Equal(t, tt.expectedD, cpu.d.Apply(logic.O, Word{}), "D")
			if writeM == logic.I {
				assert.Equal(t, tt.expectedM, outM, "outM")
			}
			assert.Equal(t, tt.expectedWriteM, writeM, "writeM")
		})
	}
}

func TestCPUJump_neg0(t *testing.T) {
	jumpToWord := Word{1, 1, 1}
	jumpTo := [15]logic.Bit{1, 1, 1}
	next := [15]logic.Bit{1}
	var tests = []struct {
		expectedPc [15]logic.Bit
		given      []logic.Bit
	}{
		{next, []logic.Bit{0, 0, 0}},
		{next, []logic.Bit{1, 0, 0}},
		{next, []logic.Bit{0, 1, 0}},
		{next, []logic.Bit{1, 1, 0}},
		{jumpTo, []logic.Bit{0, 0, 1}},
		{jumpTo, []logic.Bit{1, 0, 1}},
		{jumpTo, []logic.Bit{0, 1, 1}},
		{jumpTo, []logic.Bit{1, 1, 1}},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(fmt.Sprintf("%v", tt.given), func(t *testing.T) {
			cpu := NewCPU()
			cpu.a.Apply(logic.I, jumpToWord)
			cpu.clock.Progress()
			_, _, _, pc := cpu.Fetch(
				Word{},
				Word{tt.given[0], tt.given[1], tt.given[2], 0, 0, 0, 0, 1, 0, 1, 1, 1, 0, 1, 1, 1},
				logic.O,
			)
			assert.Equal(t, tt.expectedPc, pc)
		})
	}
}

func TestCPUJump_0(t *testing.T) {
	jumpToWord := Word{1, 1, 1}
	jumpTo := [15]logic.Bit{1, 1, 1}
	next := [15]logic.Bit{1}
	var tests = []struct {
		expectedPc [15]logic.Bit
		given      []logic.Bit
	}{
		{next, []logic.Bit{0, 0, 0}},
		{next, []logic.Bit{1, 0, 0}},
		{jumpTo, []logic.Bit{0, 1, 0}},
		{jumpTo, []logic.Bit{1, 1, 0}},
		{next, []logic.Bit{0, 0, 1}},
		{next, []logic.Bit{1, 0, 1}},
		{jumpTo, []logic.Bit{0, 1, 1}},
		{jumpTo, []logic.Bit{1, 1, 1}},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(fmt.Sprintf("%v", tt.given), func(t *testing.T) {
			cpu := NewCPU()
			cpu.a.Apply(logic.I, jumpToWord)
			cpu.clock.Progress()
			_, _, _, pc := cpu.Fetch(
				Word{},
				Word{tt.given[0], tt.given[1], tt.given[2], 0, 0, 0, 0, 1, 0, 1, 0, 1, 0, 1, 1, 1},
				logic.O,
			)
			assert.Equal(t, tt.expectedPc, pc)
		})
	}
}

func TestCPUJump_1(t *testing.T) {
	jumpToWord := Word{1, 1, 1}
	jumpTo := [15]logic.Bit{1, 1, 1}
	next := [15]logic.Bit{1}
	var tests = []struct {
		expectedPc [15]logic.Bit
		given      []logic.Bit
	}{
		{next, []logic.Bit{0, 0, 0}},
		{jumpTo, []logic.Bit{1, 0, 0}},
		{next, []logic.Bit{0, 1, 0}},
		{jumpTo, []logic.Bit{1, 1, 0}},
		{next, []logic.Bit{0, 0, 1}},
		{jumpTo, []logic.Bit{1, 0, 1}},
		{next, []logic.Bit{0, 1, 1}},
		{jumpTo, []logic.Bit{1, 1, 1}},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(fmt.Sprintf("%v", tt.given), func(t *testing.T) {
			cpu := NewCPU()
			cpu.a.Apply(logic.I, jumpToWord)
			cpu.clock.Progress()
			_, _, _, pc := cpu.Fetch(
				Word{},
				Word{tt.given[0], tt.given[1], tt.given[2], 0, 0, 0, 1, 1, 1, 1, 1, 1, 0, 1, 1, 1},
				logic.O,
			)
			assert.Equal(t, tt.expectedPc, pc)
		})
	}
}

func TestCPU_Add(t *testing.T) {
	// Add.asc
	// 0000000000000010 @2
	// 1110110000010000 D=A
	// 0000000000000011 @3
	// 1110000010010000 D=D+A
	// 0000000000000000 @0
	// 1110001100001000 M=D
	var cmds = []Word{
		{0, 1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 1, 0, 0, 0, 0, 0, 1, 1, 0, 1, 1, 1},
		{1, 1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 1, 0, 0, 1, 0, 0, 0, 0, 0, 1, 1, 1},
		{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 1, 0, 0, 0, 0, 1, 1, 0, 0, 0, 1, 1, 1},
	}

	cpu := NewCPU()

	cpu.Fetch(Word{}, cmds[0], logic.O)
	cpu.Fetch(Word{}, cmds[1], logic.O)
	cpu.Fetch(Word{}, cmds[2], logic.O)
	cpu.Fetch(Word{}, cmds[3], logic.O)
	cpu.Fetch(Word{}, cmds[4], logic.O)
	outM, writeM, addressM, pc := cpu.Fetch(Word{}, cmds[5], logic.O)

	assert.Equal(t, Word{1, 0, 1}, outM, "invalid outM")
	assert.Equal(t, logic.I, writeM, "invalid writeM")
	assert.Equal(t, [15]logic.Bit{}, addressM, "invalid addressM")
	assert.Equal(t, [15]logic.Bit{0, 1, 1}, pc, "invalid pc")
	assert.Equal(t, Word{0, 0, 0}, cpu.a.Apply(logic.O, Word{}), "invalid regA")
	assert.Equal(t, Word{1, 0, 1}, cpu.d.Apply(logic.O, Word{}), "invalid regD")
}

func TestCPU_Max(t *testing.T) {
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
	var cmds = []Word{
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

	cpu := NewCPU()

	// R0=7 and R1=1

	// 0    000 0 000000 000 000 // @0
	outM, writeM, addressM, pc := cpu.Fetch(Word{}, cmds[0], logic.O)
	assert.Equal(t, Word{}, outM, "invalid outM")
	assert.Equal(t, logic.O, writeM, "invalid writeM")
	assert.Equal(t, [15]logic.Bit{}, addressM, "invalid addressM")
	assert.Equal(t, [15]logic.Bit{1}, pc, "invalid pc")
	assert.Equal(t, Word{0, 0, 0}, cpu.a.Apply(logic.O, Word{}), "invalid regA")
	assert.Equal(t, Word{0, 0, 0}, cpu.d.Apply(logic.O, Word{}), "invalid regD")

	// 1    111 1 110000 010 000 // D=M
	outM, writeM, addressM, pc = cpu.Fetch(Word{1, 1, 1}, cmds[1], logic.O)
	// assert.Equal(t, Word{}, outM, "invalid outM")
	assert.Equal(t, logic.O, writeM, "invalid writeM")
	assert.Equal(t, [15]logic.Bit{}, addressM, "invalid addressM")
	assert.Equal(t, [15]logic.Bit{0, 1}, pc, "invalid pc")
	assert.Equal(t, Word{0, 0, 0}, cpu.a.Apply(logic.O, Word{}), "invalid regA")
	assert.Equal(t, Word{1, 1, 1}, cpu.d.Apply(logic.O, Word{}), "invalid regD")

	// 2    000 0 000000 000 001 // @1
	outM, writeM, addressM, pc = cpu.Fetch(Word{1}, cmds[2], logic.O)
	assert.Equal(t, Word{}, outM, "invalid outM")
	assert.Equal(t, logic.O, writeM, "invalid writeM")
	assert.Equal(t, [15]logic.Bit{1}, addressM, "invalid addressM")
	assert.Equal(t, [15]logic.Bit{1, 1}, pc, "invalid pc")
	assert.Equal(t, Word{1, 0, 0}, cpu.a.Apply(logic.O, Word{}), "invalid regA")
	assert.Equal(t, Word{1, 1, 1}, cpu.d.Apply(logic.O, Word{}), "invalid regD")

	// 3    111 1 010011 010 000 // D=D-M
	outM, writeM, addressM, pc = cpu.Fetch(Word{1}, cmds[3], logic.O)
	// assert.Equal(t, Word{}, outM, "invalid outM")
	assert.Equal(t, logic.O, writeM, "invalid writeM")
	assert.Equal(t, [15]logic.Bit{1}, addressM, "invalid addressM")
	assert.Equal(t, [15]logic.Bit{0, 0, 1}, pc, "invalid pc")
	assert.Equal(t, Word{1, 0, 0}, cpu.a.Apply(logic.O, Word{}), "invalid regA")
	assert.Equal(t, Word{0, 1, 1}, cpu.d.Apply(logic.O, Word{}), "invalid regD")

	// 4    000 0 000000 001 010 // @10
	outM, writeM, addressM, pc = cpu.Fetch(Word{1}, cmds[4], logic.O)
	// assert.Equal(t, Word{}, outM, "invalid outM")
	assert.Equal(t, logic.O, writeM, "invalid writeM")
	assert.Equal(t, [15]logic.Bit{0, 1, 0, 1}, addressM, "invalid addressM")
	assert.Equal(t, [15]logic.Bit{1, 0, 1}, pc, "invalid pc")
	assert.Equal(t, Word{0, 1, 0, 1}, cpu.a.Apply(logic.O, Word{}), "invalid regA")
	assert.Equal(t, Word{0, 1, 1}, cpu.d.Apply(logic.O, Word{}), "invalid regD")

	// 5    111 0 001100 000 001 // D;JGT
	outM, writeM, addressM, pc = cpu.Fetch(Word{}, cmds[5], logic.O)
	// assert.Equal(t, Word{}, outM, "invalid outM")
	assert.Equal(t, logic.O, writeM, "invalid writeM")
	assert.Equal(t, [15]logic.Bit{0, 1, 0, 1}, addressM, "invalid addressM")
	assert.Equal(t, [15]logic.Bit{0, 1, 0, 1}, pc, "invalid pc")
	assert.Equal(t, Word{0, 1, 0, 1}, cpu.a.Apply(logic.O, Word{}), "invalid regA")
	assert.Equal(t, Word{0, 1, 1}, cpu.d.Apply(logic.O, Word{}), "invalid regD")

	// 10   000 0 000000 000 000 // @0
	outM, writeM, addressM, pc = cpu.Fetch(Word{}, cmds[10], logic.O)
	// assert.Equal(t, Word{}, outM, "invalid outM")
	assert.Equal(t, logic.O, writeM, "invalid writeM")
	assert.Equal(t, [15]logic.Bit{}, addressM, "invalid addressM")
	assert.Equal(t, [15]logic.Bit{1, 1, 0, 1}, pc, "invalid pc")
	assert.Equal(t, Word{0, 0, 0, 0}, cpu.a.Apply(logic.O, Word{}), "invalid regA")
	assert.Equal(t, Word{0, 1, 1}, cpu.d.Apply(logic.O, Word{}), "invalid regD")

	// 11   111 1 110000 010 000 // D=M
	outM, writeM, addressM, pc = cpu.Fetch(Word{1, 1, 1}, cmds[11], logic.O)
	// assert.Equal(t, Word{}, outM, "invalid outM")
	assert.Equal(t, logic.O, writeM, "invalid writeM")
	assert.Equal(t, [15]logic.Bit{}, addressM, "invalid addressM")
	assert.Equal(t, [15]logic.Bit{0, 0, 1, 1}, pc, "invalid pc")
	assert.Equal(t, Word{0, 0, 0, 0}, cpu.a.Apply(logic.O, Word{}), "invalid regA")
	assert.Equal(t, Word{1, 1, 1}, cpu.d.Apply(logic.O, Word{}), "invalid regD")

	// 12   000 0 000000 000 010 // @2
	outM, writeM, addressM, pc = cpu.Fetch(Word{1, 1, 1}, cmds[12], logic.O)
	// assert.Equal(t, Word{}, outM, "invalid outM")
	assert.Equal(t, logic.O, writeM, "invalid writeM")
	assert.Equal(t, [15]logic.Bit{0, 1}, addressM, "invalid addressM")
	assert.Equal(t, [15]logic.Bit{1, 0, 1, 1}, pc, "invalid pc")
	assert.Equal(t, Word{0, 1, 0, 0}, cpu.a.Apply(logic.O, Word{}), "invalid regA")
	assert.Equal(t, Word{1, 1, 1}, cpu.d.Apply(logic.O, Word{}), "invalid regD")

	// 13   111 0 001100 001 000 // M=D
	outM, writeM, addressM, pc = cpu.Fetch(Word{}, cmds[13], logic.O)
	assert.Equal(t, Word{1, 1, 1}, outM, "invalid outM")
	assert.Equal(t, logic.I, writeM, "invalid writeM")
	assert.Equal(t, [15]logic.Bit{0, 1}, addressM, "invalid addressM")
	assert.Equal(t, [15]logic.Bit{0, 1, 1, 1}, pc, "invalid pc")
	assert.Equal(t, Word{0, 1, 0, 0}, cpu.a.Apply(logic.O, Word{}), "invalid regA")
	assert.Equal(t, Word{1, 1, 1}, cpu.d.Apply(logic.O, Word{}), "invalid regD")

	// Reset pc
	outM, writeM, addressM, pc = cpu.Fetch(Word{}, cmds[0], logic.I)
	assert.Equal(t, [15]logic.Bit{0}, pc, "invalid pc")

	// R0=7 and R1=8

	// 0    000 0 000000 000 000 // @0
	outM, writeM, addressM, pc = cpu.Fetch(Word{}, cmds[0], logic.O)
	assert.Equal(t, Word{}, outM, "invalid outM")
	assert.Equal(t, logic.O, writeM, "invalid writeM")
	assert.Equal(t, [15]logic.Bit{}, addressM, "invalid addressM")
	assert.Equal(t, [15]logic.Bit{1}, pc, "invalid pc")
	assert.Equal(t, Word{0, 0, 0}, cpu.a.Apply(logic.O, Word{}), "invalid regA")
	assert.Equal(t, Word{1, 1, 1}, cpu.d.Apply(logic.O, Word{}), "invalid regD")

	// 1    111 1 110000 010 000 // D=M
	outM, writeM, addressM, pc = cpu.Fetch(Word{1, 1, 1}, cmds[1], logic.O)
	// assert.Equal(t, Word{}, outM, "invalid outM")
	assert.Equal(t, logic.O, writeM, "invalid writeM")
	assert.Equal(t, [15]logic.Bit{}, addressM, "invalid addressM")
	assert.Equal(t, [15]logic.Bit{0, 1}, pc, "invalid pc")
	assert.Equal(t, Word{0, 0, 0}, cpu.a.Apply(logic.O, Word{}), "invalid regA")
	assert.Equal(t, Word{1, 1, 1}, cpu.d.Apply(logic.O, Word{}), "invalid regD")

	// 2    000 0 000000 000 001 // @1
	outM, writeM, addressM, pc = cpu.Fetch(Word{}, cmds[2], logic.O)
	// assert.Equal(t, Word{}, outM, "invalid outM")
	assert.Equal(t, logic.O, writeM, "invalid writeM")
	assert.Equal(t, [15]logic.Bit{1}, addressM, "invalid addressM")
	assert.Equal(t, [15]logic.Bit{1, 1}, pc, "invalid pc")
	assert.Equal(t, Word{1, 0, 0}, cpu.a.Apply(logic.O, Word{}), "invalid regA")
	assert.Equal(t, Word{1, 1, 1}, cpu.d.Apply(logic.O, Word{}), "invalid regD")

	// 3    111 1 010011 010 000 // D=D-M
	outM, writeM, addressM, pc = cpu.Fetch(Word{0, 0, 0, 1}, cmds[3], logic.O)
	// assert.Equal(t, Word{}, outM, "invalid outM")
	assert.Equal(t, logic.O, writeM, "invalid writeM")
	assert.Equal(t, [15]logic.Bit{1}, addressM, "invalid addressM")
	assert.Equal(t, [15]logic.Bit{0, 0, 1}, pc, "invalid pc")
	assert.Equal(t, Word{1, 0, 0}, cpu.a.Apply(logic.O, Word{}), "invalid regA")
	assert.Equal(t, Word{1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1}, cpu.d.Apply(logic.O, Word{}), "invalid regD")

	// 4    000 0 000000 001 010 // @10
	outM, writeM, addressM, pc = cpu.Fetch(Word{}, cmds[4], logic.O)
	// assert.Equal(t, Word{}, outM, "invalid outM")
	assert.Equal(t, logic.O, writeM, "invalid writeM")
	assert.Equal(t, [15]logic.Bit{0, 1, 0, 1}, addressM, "invalid addressM")
	assert.Equal(t, [15]logic.Bit{1, 0, 1}, pc, "invalid pc")
	assert.Equal(t, Word{0, 1, 0, 1}, cpu.a.Apply(logic.O, Word{}), "invalid regA")
	assert.Equal(t, Word{1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1}, cpu.d.Apply(logic.O, Word{}), "invalid regD")

	// 5    111 0 001100 000 001 // D;JGT
	outM, writeM, addressM, pc = cpu.Fetch(Word{}, cmds[5], logic.O)
	// assert.Equal(t, Word{}, outM, "invalid outM")
	assert.Equal(t, logic.O, writeM, "invalid writeM")
	assert.Equal(t, [15]logic.Bit{0, 1, 0, 1}, addressM, "invalid addressM")
	assert.Equal(t, [15]logic.Bit{0, 1, 1}, pc, "invalid pc")
	assert.Equal(t, Word{0, 1, 0, 1}, cpu.a.Apply(logic.O, Word{}), "invalid regA")
	assert.Equal(t, Word{1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1}, cpu.d.Apply(logic.O, Word{}), "invalid regD")

	// 6    000 0 000000 000 001 // @1
	outM, writeM, addressM, pc = cpu.Fetch(Word{}, cmds[6], logic.O)
	// assert.Equal(t, Word{}, outM, "invalid outM")
	assert.Equal(t, logic.O, writeM, "invalid writeM")
	assert.Equal(t, [15]logic.Bit{1, 0, 0}, addressM, "invalid addressM")
	assert.Equal(t, [15]logic.Bit{1, 1, 1}, pc, "invalid pc")
	assert.Equal(t, Word{1, 0, 0}, cpu.a.Apply(logic.O, Word{}), "invalid regA")
	assert.Equal(t, Word{1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1}, cpu.d.Apply(logic.O, Word{}), "invalid regD")

	// 7    111 1 110000 010 000 // D=M
	outM, writeM, addressM, pc = cpu.Fetch(Word{0, 0, 0, 1}, cmds[7], logic.O)
	// assert.Equal(t, Word{}, outM, "invalid outM")
	assert.Equal(t, logic.O, writeM, "invalid writeM")
	assert.Equal(t, [15]logic.Bit{1, 0, 0}, addressM, "invalid addressM")
	assert.Equal(t, [15]logic.Bit{0, 0, 0, 1}, pc, "invalid pc")
	assert.Equal(t, Word{1, 0, 0}, cpu.a.Apply(logic.O, Word{}), "invalid regA")
	assert.Equal(t, Word{0, 0, 0, 1}, cpu.d.Apply(logic.O, Word{}), "invalid regD")

	// 8    000 0 000000 001 100 // @12
	outM, writeM, addressM, pc = cpu.Fetch(Word{0, 0, 0, 1}, cmds[8], logic.O)
	// assert.Equal(t, Word{}, outM, "invalid outM")
	assert.Equal(t, logic.O, writeM, "invalid writeM")
	assert.Equal(t, [15]logic.Bit{0, 0, 1, 1}, addressM, "invalid addressM")
	assert.Equal(t, [15]logic.Bit{1, 0, 0, 1}, pc, "invalid pc")
	assert.Equal(t, Word{0, 0, 1, 1}, cpu.a.Apply(logic.O, Word{}), "invalid regA")
	assert.Equal(t, Word{0, 0, 0, 1}, cpu.d.Apply(logic.O, Word{}), "invalid regD")

	// 9    111 0 101010 000 111 // 0;JMP
	outM, writeM, addressM, pc = cpu.Fetch(Word{}, cmds[9], logic.O)
	// assert.Equal(t, Word{}, outM, "invalid outM")
	assert.Equal(t, logic.O, writeM, "invalid writeM")
	assert.Equal(t, [15]logic.Bit{0, 0, 1, 1}, addressM, "invalid addressM")
	assert.Equal(t, [15]logic.Bit{0, 0, 1, 1}, pc, "invalid pc")
	assert.Equal(t, Word{0, 0, 1, 1}, cpu.a.Apply(logic.O, Word{}), "invalid regA")
	assert.Equal(t, Word{0, 0, 0, 1}, cpu.d.Apply(logic.O, Word{}), "invalid regD")

	// 12   000 0 000000 000 010 // @2
	outM, writeM, addressM, pc = cpu.Fetch(Word{}, cmds[12], logic.O)
	// assert.Equal(t, Word{}, outM, "invalid outM")
	assert.Equal(t, logic.O, writeM, "invalid writeM")
	assert.Equal(t, [15]logic.Bit{0, 1, 0, 0}, addressM, "invalid addressM")
	assert.Equal(t, [15]logic.Bit{1, 0, 1, 1}, pc, "invalid pc")
	assert.Equal(t, Word{0, 1, 0, 0}, cpu.a.Apply(logic.O, Word{}), "invalid regA")
	assert.Equal(t, Word{0, 0, 0, 1}, cpu.d.Apply(logic.O, Word{}), "invalid regD")

	// 13   111 0 001100 001 000 // M=D
	outM, writeM, addressM, pc = cpu.Fetch(Word{}, cmds[13], logic.O)
	assert.Equal(t, Word{0, 0, 0, 1}, outM, "invalid outM")
	assert.Equal(t, logic.I, writeM, "invalid writeM")
	assert.Equal(t, [15]logic.Bit{0, 1, 0, 0}, addressM, "invalid addressM")
	assert.Equal(t, [15]logic.Bit{0, 1, 1, 1}, pc, "invalid pc")
	assert.Equal(t, Word{0, 1, 0, 0}, cpu.a.Apply(logic.O, Word{}), "invalid regA")
	assert.Equal(t, Word{0, 0, 0, 1}, cpu.d.Apply(logic.O, Word{}), "invalid regD")

	// 14   000 0 000000 001 110 // @14
	outM, writeM, addressM, pc = cpu.Fetch(Word{}, cmds[14], logic.O)
	// assert.Equal(t, Word{0, 0, 0, 1}, outM, "invalid outM")
	assert.Equal(t, logic.O, writeM, "invalid writeM")
	assert.Equal(t, [15]logic.Bit{0, 1, 1, 1}, addressM, "invalid addressM")
	assert.Equal(t, [15]logic.Bit{1, 1, 1, 1}, pc, "invalid pc")
	assert.Equal(t, Word{0, 1, 1, 1}, cpu.a.Apply(logic.O, Word{}), "invalid regA")
	assert.Equal(t, Word{0, 0, 0, 1}, cpu.d.Apply(logic.O, Word{}), "invalid regD")

	// 15   111 0 101010 000 111 // 0;JMP
	outM, writeM, addressM, pc = cpu.Fetch(Word{}, cmds[15], logic.O)
	// assert.Equal(t, Word{0, 0, 0, 1}, outM, "invalid outM")
	assert.Equal(t, logic.O, writeM, "invalid writeM")
	assert.Equal(t, [15]logic.Bit{0, 1, 1, 1}, addressM, "invalid addressM")
	assert.Equal(t, [15]logic.Bit{0, 1, 1, 1}, pc, "invalid pc")
	assert.Equal(t, Word{0, 1, 1, 1}, cpu.a.Apply(logic.O, Word{}), "invalid regA")
	assert.Equal(t, Word{0, 0, 0, 1}, cpu.d.Apply(logic.O, Word{}), "invalid regD")
}

func TestCPUdecode(t *testing.T) {
	cpu := NewCPU()
	i, a, cccccc, ddd, jjj := cpu.decode(Word{0, 1, 0, 1, 1, 0, 1, 1, 1, 0, 1, 1, 1, 1, 0, 1})
	eI := logic.I
	eA := logic.I
	eCCCCCC := []logic.Bit{1, 1, 1, 0, 1, 1}
	eDDD := []logic.Bit{1, 1, 0}
	eJJJ := []logic.Bit{0, 1, 0}
	if i != eI {
		t.Errorf("expected %v, actual %v", eI, i)
	}
	if a != eA {
		t.Errorf("expected %v, actual %v", eA, a)
	}
	if !reflect.DeepEqual(cccccc, eCCCCCC) {
		t.Errorf("expected %v, actual %v", eCCCCCC, cccccc)
	}
	if !reflect.DeepEqual(ddd, eDDD) {
		t.Errorf("expected %v, actual %v", eDDD, ddd)
	}
	if !reflect.DeepEqual(jjj, eJJJ) {
		t.Errorf("expected %v, actual %v", eJJJ, jjj)
	}
}

func TestCPUjump(t *testing.T) {
	cpu := NewCPU()

	var tests = []struct {
		expected Bit
		givenZr  Bit
		givenNg  Bit
		givenJJJ []Bit
	}{
		{logic.O, logic.O, logic.O, []logic.Bit{0, 0, 0}},
		{logic.I, logic.O, logic.O, []logic.Bit{1, 1, 1}},

		{logic.I, logic.O, logic.O, []logic.Bit{1, 0, 0}}, // JGT
		{logic.O, logic.O, logic.O, []logic.Bit{0, 1, 0}}, // JEQ
		{logic.I, logic.O, logic.O, []logic.Bit{1, 1, 0}}, // JGE
		{logic.O, logic.O, logic.O, []logic.Bit{0, 0, 1}}, // JLT
		{logic.I, logic.O, logic.O, []logic.Bit{1, 0, 1}}, // JNE
		{logic.O, logic.O, logic.O, []logic.Bit{0, 1, 1}}, // JLE

		{logic.O, logic.O, logic.I, []logic.Bit{1, 0, 0}}, // JGT
		{logic.O, logic.O, logic.I, []logic.Bit{0, 1, 0}}, // JEQ
		{logic.O, logic.O, logic.I, []logic.Bit{1, 1, 0}}, // JGE
		{logic.I, logic.O, logic.I, []logic.Bit{0, 0, 1}}, // JLT
		{logic.I, logic.O, logic.I, []logic.Bit{1, 0, 1}}, // JNE
		{logic.I, logic.O, logic.I, []logic.Bit{0, 1, 1}}, // JLE

		{logic.O, logic.I, logic.O, []logic.Bit{1, 0, 0}}, // JGT
		{logic.I, logic.I, logic.O, []logic.Bit{0, 1, 0}}, // JEQ
		{logic.I, logic.I, logic.O, []logic.Bit{1, 1, 0}}, // JGE
		{logic.O, logic.I, logic.O, []logic.Bit{0, 0, 1}}, // JLT
		{logic.O, logic.I, logic.O, []logic.Bit{1, 0, 1}}, // JNE
		{logic.I, logic.I, logic.O, []logic.Bit{0, 1, 1}}, // JLE

	}
	for _, tt := range tests {
		tt := tt
		given := fmt.Sprintf("%v %v %v", tt.givenZr, tt.givenNg, tt.givenJJJ)
		t.Run(given, func(t *testing.T) {
			actual := cpu.jump(tt.givenZr, tt.givenNg, tt.givenJJJ)
			if actual != tt.expected {
				t.Errorf("given(%s): expected %v, actual %v", given, tt.expected, actual)
			}
		})
	}
}
