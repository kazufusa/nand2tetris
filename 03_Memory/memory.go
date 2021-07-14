package memory

import (
	logic "github.com/kazufusa/nand2tetris/01_Boolean_Logic"
)

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

func NewBit(c *Clock) *Bit {
	return &Bit{dff: NewDFF(c)}
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
