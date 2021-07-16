package memory

import (
	"testing"

	logic "github.com/kazufusa/nand2tetris/01_Boolean_Logic"
)

func TestDFF(t *testing.T) {
	clock := NewClock()
	dff := NewDFF(clock)

	var tests = []struct {
		expectedBeforeProgress logic.Bit
		expected               logic.Bit
		given                  logic.Bit
	}{
		{logic.O, logic.O, logic.O},
		{logic.O, logic.I, logic.I},
		{logic.I, logic.O, logic.O},
		{logic.O, logic.I, logic.I},
		{logic.I, logic.O, logic.O},
		{logic.O, logic.I, logic.I},
		{logic.I, logic.O, logic.O},
		{logic.O, logic.I, logic.I},
		{logic.I, logic.O, logic.O},
		{logic.O, logic.I, logic.I},
	}
	for _, tt := range tests {
		tt := tt
		t.Run("", func(t *testing.T) {

			actual := dff.Apply(tt.given)
			if actual != tt.expectedBeforeProgress {
				t.Errorf("given(%v): expected %v, actual %v", tt.given, tt.expectedBeforeProgress, actual)
			}

			clock.Progress()

			actual = dff.Apply(tt.given)
			if actual != tt.expected {
				t.Errorf("given(%v): expected %v, actual %v", tt.given, tt.expected, actual)
			}
		})
	}
}

func TestBit(t *testing.T) {
	clock := NewClock()
	bit := NewBit(clock)

	bit.Apply(logic.I, logic.I)
	var tests = []struct {
		name     string
		expected logic.Bit
		given    [2]logic.Bit // load, in
	}{
		{"cycle 2", logic.I, [2]logic.Bit{logic.O, logic.O}},
		{"cycle 3", logic.I, [2]logic.Bit{logic.O, logic.O}},
		{"cycle 4", logic.I, [2]logic.Bit{logic.I, logic.O}},
		{"cycle 5", logic.O, [2]logic.Bit{logic.O, logic.O}},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			clock.Progress()
			actual := bit.Apply(tt.given[0], tt.given[1])
			if actual != tt.expected {
				t.Errorf("given(%v): expected %v, actual %v", tt.given, tt.expected, actual)
			}
		})
	}
}

func TestRegister(t *testing.T) {
	clock := NewClock()
	register := NewRegister(clock)

	register.Apply(
		logic.I,
		Word{1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0},
	)
	var tests = []struct {
		name      string
		expected  Word
		givenLoad logic.Bit
		givenIn   Word
	}{
		{"cycle 2",
			Word{1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0},
			logic.O,
			Word{0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1},
		},
		{"cycle 3",
			Word{1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0},
			logic.O,
			Word{0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1},
		},
		{"cycle 4",
			Word{1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0},
			logic.I,
			Word{0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1},
		},
		{"cycle 5",
			Word{0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1},
			logic.O,
			Word{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			clock.Progress()
			actual := register.Apply(tt.givenLoad, tt.givenIn)
			if actual != tt.expected {
				t.Errorf("given(%v): expected %v, actual %v", tt.givenLoad, tt.expected, actual)
			}
		})
	}
}

func TestRam8_1(t *testing.T) {
	clock := NewClock()
	ram := NewRAM8(clock)
	w0 := Word{}
	w1 := logic.Not16(Word{})
	w01 := Word{1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0}
	addrs := [][3]logic.Bit{
		{0, 0, 0},
		{1, 0, 0},
		{0, 1, 0},
		{1, 1, 0},
		{0, 0, 1},
		{1, 0, 1},
		{0, 1, 1},
		{1, 1, 1},
	}

	ram.Apply(w01, logic.I, addrs[0])
	clock.Progress()
	actual := ram.Apply(w0, logic.O, addrs[0])
	if actual != w01 {
		t.Errorf("addr(%v): expected %v, actual %v", addrs[0], w01, actual)
	}

	ram.Apply(w1, logic.I, addrs[0])
	clock.Progress()
	actual = ram.Apply(w0, logic.O, addrs[0])
	if actual != w1 {
		t.Errorf("addr(%v): expected %v, actual %v", addrs[0], w1, actual)
	}

	ram.Apply(w01, logic.I, addrs[1])
	clock.Progress()
	actual = ram.Apply(w0, logic.O, addrs[0])
	if actual != w1 {
		t.Errorf("addr(%v): expected %v, actual %v", addrs[0], w1, actual)
	}
	actual = ram.Apply(w0, logic.O, addrs[1])
	if actual != w01 {
		t.Errorf("addr(%v): expected %v, actual %v", addrs[1], w01, actual)
	}

}

func TestRam8_2(t *testing.T) {
	clock := NewClock()
	ram := NewRAM8(clock)
	w01 := Word{1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0}
	w1 := logic.Not16(Word{})
	addrs := [][3]logic.Bit{
		{0, 0, 0},
		{1, 0, 0},
		{0, 1, 0},
		{1, 1, 0},
		{0, 0, 1},
		{1, 0, 1},
		{0, 1, 1},
		{1, 1, 1},
	}
	resetRam := func(w Word) {
		for i := 0; i < 8; i++ {
			ram.Apply(w, logic.I, addrs[i])
			clock.Progress()
		}
	}

	var tests = []struct {
		name     string
		expected Word
		addr     [3]logic.Bit
	}{
		{"0", w01, addrs[0]},
		{"1", w01, addrs[1]},
		{"2", w01, addrs[2]},
		{"3", w01, addrs[3]},
		{"4", w01, addrs[4]},
		{"5", w01, addrs[5]},
		{"6", w01, addrs[6]},
		{"7", w01, addrs[7]},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			resetRam(w1)
			ram.Apply(tt.expected, logic.I, tt.addr)
			clock.Progress()
			for _, addr := range addrs {
				if addr == tt.addr {
					actual := ram.Apply(Word{}, logic.O, addr)
					if actual != tt.expected {
						t.Errorf("addr(%v): expected %v, actual %v", addr, tt.expected, actual)
					}
				} else {
					actual := ram.Apply(Word{}, logic.O, addr)
					if actual != w1 {
						t.Errorf("addr(%v): expected %v, actual %v", addr, w1, actual)
					}
				}
			}
		})
	}
}
