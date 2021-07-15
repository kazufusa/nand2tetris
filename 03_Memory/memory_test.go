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
			dff.Input(tt.given)

			actual := dff.Output()
			if actual != tt.expectedBeforeProgress {
				t.Errorf("given(%v): expected %v, actual %v", tt.given, tt.expectedBeforeProgress, actual)
			}

			clock.Progress()

			actual = dff.Output()
			if actual != tt.expected {
				t.Errorf("given(%v): expected %v, actual %v", tt.given, tt.expected, actual)
			}
		})
	}
}

func TestBit(t *testing.T) {
	clock := NewClock()
	bit := NewBit(clock)

	bit.Input(logic.I, logic.I)
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
			bit.Input(tt.given[0], tt.given[1])
			actual := bit.Output()
			if actual != tt.expected {
				t.Errorf("given(%v): expected %v, actual %v", tt.given, tt.expected, actual)
			}
		})
	}
}

func TestRegister(t *testing.T) {
	clock := NewClock()
	register := NewRegister(clock)

	register.Input(
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
			register.Input(tt.givenLoad, tt.givenIn)
			actual := register.Output()
			if actual != tt.expected {
				t.Errorf("given(%v): expected %v, actual %v", tt.givenLoad, tt.expected, actual)
			}
		})
	}
}
