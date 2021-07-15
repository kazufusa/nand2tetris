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

func (d *DFF) Input(in logic.Bit) {
	defer func() {
		d.nextTime = *d.clock + 1
	}()

	if d.nextTime == d.clock.Now() {
		d.cur = d.next
		d.next = in
	} else if d.nextTime < d.clock.Now() {
		d.cur = logic.O
		d.next = in
	}
}
func (d *DFF) Output() logic.Bit {
	switch d.clock.Now() {
	case d.nextTime - 1:
		return d.cur
	case d.nextTime:
		return d.next
	default:
		return logic.O
	}
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
func (b *Bit) Input(load, in logic.Bit) {
	b.dff.Input(logic.Mux(b.dff.Output(), in, load))
}

// Output returns cu
func (b *Bit) Output() logic.Bit {
	return b.dff.Output()
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

func (r *Register) Input(load logic.Bit, in Word) {
	r.bits[0].Input(load, in[0])
	r.bits[1].Input(load, in[1])
	r.bits[2].Input(load, in[2])
	r.bits[3].Input(load, in[3])
	r.bits[4].Input(load, in[4])
	r.bits[5].Input(load, in[5])
	r.bits[6].Input(load, in[6])
	r.bits[7].Input(load, in[7])
	r.bits[8].Input(load, in[8])
	r.bits[9].Input(load, in[9])
	r.bits[10].Input(load, in[10])
	r.bits[11].Input(load, in[11])
	r.bits[12].Input(load, in[12])
	r.bits[13].Input(load, in[13])
	r.bits[14].Input(load, in[14])
	r.bits[15].Input(load, in[15])
}

func (r *Register) Output() Word {
	return Word{
		r.bits[0].Output(),
		r.bits[1].Output(),
		r.bits[2].Output(),
		r.bits[3].Output(),
		r.bits[4].Output(),
		r.bits[5].Output(),
		r.bits[6].Output(),
		r.bits[7].Output(),
		r.bits[8].Output(),
		r.bits[9].Output(),
		r.bits[10].Output(),
		r.bits[11].Output(),
		r.bits[12].Output(),
		r.bits[13].Output(),
		r.bits[14].Output(),
		r.bits[15].Output(),
	}
}
