package computer

import (
	logic "github.com/kazufusa/nand2tetris/01_Boolean_Logic"
	memory "github.com/kazufusa/nand2tetris/03_Memory"
)

type Memory struct {
	ram      memory.RAM16384
	rams     [4]memory.RAM4096
	screen   IScreen
	keyboard IKeyboard
}

func NewMemory(clock *memory.Clock, screen IScreen, keyboard IKeyboard) Memory {
	return Memory{
		ram:      memory.NewRAM16384(clock),
		screen:   screen,
		keyboard: keyboard,
	}
}

// 8K RAM:        0-16383
// 8K Screen:     16384-24575
// 1Bit Keyboard: 24576
// 16383: 0 011 111111111111
// 16384: 0 100 000000000000
// 24575: 0 101 111111111111
// 24576: 0 110 000000000000
func (m *Memory) Fetch(in Word, load Bit, addr [15]Bit) (out Word) {
	addr8k := [13]logic.Bit{
		addr[0], addr[1], addr[2], addr[3],
		addr[4], addr[5], addr[6], addr[7],
		addr[8], addr[9], addr[10], addr[11],
		addr[12],
	}
	addr16k := [14]logic.Bit{
		addr[0], addr[1], addr[2], addr[3],
		addr[4], addr[5], addr[6], addr[7],
		addr[8], addr[9], addr[10], addr[11],
		addr[12], addr[13],
	}
	loads := logic.Dmux8Way(load, [3]logic.Bit{addr[12], addr[13], addr[14]})
	return logic.Mux16(
		m.ram.Apply(in, logic.And(load, logic.Not(addr[14])), addr16k),
		logic.Mux16(
			m.screen.Fetch(in, logic.And(load, logic.Or(loads[4], loads[5])), addr8k),
			m.keyboard.Fetch(),
			addr[13],
		),
		addr[14],
	)
}

type IScreen interface {
	Fetch(in Word, load Bit, addr [13]Bit) Word
}

type IKeyboard interface {
	Fetch() Word
}
