package memory

import (
	logic "github.com/kazufusa/nand2tetris/01_Boolean_Logic"
)

type Word [16]logic.Bit

type Clock uint8

func NewClock() *Clock {
	var c Clock = 0
	return &c
}

func (c *Clock) Now() Clock {
	return *c
}

func (c *Clock) Progress() {
	*c = *c + 1
}

type DFF struct {
	clock     *Clock
	nextTime  Clock
	cur, next logic.Bit
}

func NewDFF(c *Clock) DFF {
	return DFF{clock: c}
}

func (d *DFF) Apply(in logic.Bit) logic.Bit {
	defer func() {
		d.nextTime = d.clock.Now() + 1
	}()

	if d.nextTime == d.clock.Now() {
		d.cur = d.next
		d.next = in
	} else if d.nextTime == d.clock.Now()+1 {
		d.next = in
	} else if d.nextTime < d.clock.Now() {
		d.cur = logic.O
		d.next = in
	}
	return d.cur
}

type Bit struct {
	dff DFF
}

func NewBit(c *Clock) Bit {
	return Bit{dff: NewDFF(c)}
}

// Input gets logic.Bits of load and in.
// If load, Input inputs in to DFF
// else Input gets output from DFF and inputs the output to DFF.
func (b *Bit) Apply(load, in logic.Bit) logic.Bit {
	return b.dff.Apply(logic.Mux(b.dff.Apply(logic.O), in, load))
}

// 16-bit register
type Register struct {
	bits [16]Bit
}

func NewRegister(c *Clock) Register {
	return Register{
		bits: [16]Bit{
			NewBit(c), NewBit(c), NewBit(c), NewBit(c),
			NewBit(c), NewBit(c), NewBit(c), NewBit(c),
			NewBit(c), NewBit(c), NewBit(c), NewBit(c),
			NewBit(c), NewBit(c), NewBit(c), NewBit(c),
		},
	}
}

func (r *Register) Apply(load logic.Bit, in Word) Word {
	return Word{
		r.bits[0].Apply(load, in[0]),
		r.bits[1].Apply(load, in[1]),
		r.bits[2].Apply(load, in[2]),
		r.bits[3].Apply(load, in[3]),
		r.bits[4].Apply(load, in[4]),
		r.bits[5].Apply(load, in[5]),
		r.bits[6].Apply(load, in[6]),
		r.bits[7].Apply(load, in[7]),
		r.bits[8].Apply(load, in[8]),
		r.bits[9].Apply(load, in[9]),
		r.bits[10].Apply(load, in[10]),
		r.bits[11].Apply(load, in[11]),
		r.bits[12].Apply(load, in[12]),
		r.bits[13].Apply(load, in[13]),
		r.bits[14].Apply(load, in[14]),
		r.bits[15].Apply(load, in[15]),
	}
}

// RAM8 consists of 8 registers
type RAM8 struct {
	registers [8]Register
}

func NewRAM8(clock *Clock) RAM8 {
	return RAM8{
		registers: [8]Register{
			NewRegister(clock), NewRegister(clock), NewRegister(clock), NewRegister(clock),
			NewRegister(clock), NewRegister(clock), NewRegister(clock), NewRegister(clock),
		},
	}
}

func (r8 *RAM8) Apply(in Word, load logic.Bit, addr [3]logic.Bit) Word {
	dAddr := logic.Dmux8Way(logic.I, addr)
	return logic.Mux8Way16(
		r8.registers[0].Apply(logic.And(load, dAddr[0]), in),
		r8.registers[1].Apply(logic.And(load, dAddr[1]), in),
		r8.registers[2].Apply(logic.And(load, dAddr[2]), in),
		r8.registers[3].Apply(logic.And(load, dAddr[3]), in),
		r8.registers[4].Apply(logic.And(load, dAddr[4]), in),
		r8.registers[5].Apply(logic.And(load, dAddr[5]), in),
		r8.registers[6].Apply(logic.And(load, dAddr[6]), in),
		r8.registers[7].Apply(logic.And(load, dAddr[7]), in),
		addr,
	)
}

// RAM64 consists of 8 RAM8
type RAM64 struct {
	rams [8]RAM8
}

// NewRAM64 returns new RAM64 object
func NewRAM64(clock *Clock) RAM64 {
	return RAM64{
		rams: [8]RAM8{
			NewRAM8(clock), NewRAM8(clock), NewRAM8(clock), NewRAM8(clock),
			NewRAM8(clock), NewRAM8(clock), NewRAM8(clock), NewRAM8(clock),
		},
	}
}

func (r *RAM64) Apply(in Word, load logic.Bit, addr [6]logic.Bit) Word {
	regAddr := [3]logic.Bit{addr[0], addr[1], addr[2]}
	ram8Addr := [3]logic.Bit{addr[3], addr[4], addr[5]}
	dAddr := logic.Dmux8Way(logic.I, ram8Addr)
	return logic.Mux8Way16(
		r.rams[0].Apply(in, logic.And(load, dAddr[0]), regAddr),
		r.rams[1].Apply(in, logic.And(load, dAddr[1]), regAddr),
		r.rams[2].Apply(in, logic.And(load, dAddr[2]), regAddr),
		r.rams[3].Apply(in, logic.And(load, dAddr[3]), regAddr),
		r.rams[4].Apply(in, logic.And(load, dAddr[4]), regAddr),
		r.rams[5].Apply(in, logic.And(load, dAddr[5]), regAddr),
		r.rams[6].Apply(in, logic.And(load, dAddr[6]), regAddr),
		r.rams[7].Apply(in, logic.And(load, dAddr[7]), regAddr),
		ram8Addr,
	)
}
