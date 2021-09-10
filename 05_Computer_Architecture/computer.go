package computer

import (
	logic "github.com/kazufusa/nand2tetris/01_Boolean_Logic"
	memory "github.com/kazufusa/nand2tetris/03_Memory"
)

type IMemory interface {
	Fetch(in Word, load Bit, addr [15]Bit) (out Word)
}

type IROM32K interface {
	Fetch(addr [15]Bit) (out Word)
}

type Computer struct {
	cpu   *CPU
	ram   IMemory
	rom   IROM32K
	clock *memory.Clock

	pc       [15]logic.Bit
	inM      Word
	addressM [15]logic.Bit
}

func NewComputer(cpu *CPU, ram IMemory, rom IROM32K, clock *memory.Clock) Computer {
	return Computer{cpu: cpu, ram: ram, rom: rom, clock: clock}
}

func (com *Computer) FetchAndExecute(reset logic.Bit) {
	var outM Word
	var writeM logic.Bit
	inst := com.rom.Fetch(com.pc)
	outM, writeM, com.addressM, com.pc = com.cpu.Fetch(com.inM, inst, reset)

	com.ram.Fetch(outM, writeM, com.addressM)
	com.clock.Progress()
	com.inM = com.ram.Fetch(com.inM, logic.O, com.addressM)
}
