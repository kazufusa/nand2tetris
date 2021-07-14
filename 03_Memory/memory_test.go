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
