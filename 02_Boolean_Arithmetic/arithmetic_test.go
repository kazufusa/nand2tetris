package arithmetic

import (
	"fmt"
	"testing"

	logic "github.com/kazufusa/nand2tetris/01_Boolean_Logic"
)

func TestHalfAdder(t *testing.T) {
	var tests = []struct {
		expected [2]Bit // sum and carry
		given    [2]Bit // a, b, and c
	}{
		{[2]Bit{0, 0}, [2]Bit{0, 0}},
		{[2]Bit{1, 0}, [2]Bit{0, 1}},
		{[2]Bit{1, 0}, [2]Bit{1, 0}},
		{[2]Bit{0, 1}, [2]Bit{1, 1}},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(fmt.Sprintf("%v", tt.given), func(t *testing.T) {
			actual := [2]Bit{}
			actual[0], actual[1] = halfAdder(tt.given[0], tt.given[1])
			if actual != tt.expected {
				t.Errorf("given(%v): expected %v, actual %v", tt.given, tt.expected, actual)
			}
		})
	}
}

func TestFullAdder(t *testing.T) {
	var tests = []struct {
		expected [2]Bit // sum and carry
		given    [3]Bit // a, b, and c
	}{
		{[2]Bit{0, 0}, [3]Bit{0, 0, 0}},
		{[2]Bit{1, 0}, [3]Bit{0, 0, 1}},
		{[2]Bit{1, 0}, [3]Bit{0, 1, 0}},
		{[2]Bit{0, 1}, [3]Bit{0, 1, 1}},
		{[2]Bit{1, 0}, [3]Bit{1, 0, 0}},
		{[2]Bit{0, 1}, [3]Bit{1, 0, 1}},
		{[2]Bit{0, 1}, [3]Bit{1, 1, 0}},
		{[2]Bit{1, 1}, [3]Bit{1, 1, 1}},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(fmt.Sprintf("%v", tt.given), func(t *testing.T) {
			actual := [2]Bit{}
			actual[0], actual[1] = fullAdder(tt.given[0], tt.given[1], tt.given[2])
			if actual != tt.expected {
				t.Errorf("given(%v): expected %v, actual %v", tt.given, tt.expected, actual)
			}
		})
	}
}

func TestAdder(t *testing.T) {
	var tests = []struct {
		expected Word
		givenA   Word
		givenB   Word
	}{
		{
			Word{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
			Word{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
			Word{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
		},
		{
			Word{0, 1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
			Word{1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
			Word{1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
		},
		{
			Word{1, 0, 1, 1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
			Word{1, 1, 0, 1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
			Word{0, 1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
		},
		{
			Word{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 0, 1, 1},
			Word{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 1, 0, 1},
			Word{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 0, 0},
		},
		{
			Word{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 0, 0, 1},
			Word{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 1, 0, 1},
			Word{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 1, 1},
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(fmt.Sprintf("%v, %v", tt.givenA, tt.givenB), func(t *testing.T) {
			actual := Adder(tt.givenA, tt.givenB)
			if actual != tt.expected {
				t.Errorf("given(%v,%v): expected %v, actual %v", tt.givenA, tt.givenB, tt.expected, actual)
			}
		})
	}
}

func TestInc16(t *testing.T) {
	var tests = []struct {
		expected Word
		given    Word
	}{
		{
			Word{1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
			Word{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
		},
		{
			Word{0, 1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
			Word{1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
		},
		{
			Word{1, 0, 1, 1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
			Word{0, 0, 1, 1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
		},
		{
			Word{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
			Word{1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1},
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(fmt.Sprintf("%v", tt.given), func(t *testing.T) {
			actual := Inc16(tt.given)
			if actual != tt.expected {
				t.Errorf("given(%v): expected %v, actual %v", tt.given, tt.expected, actual)
			}
		})
	}
}

func TestALU(t *testing.T) {
	var tests = []struct {
		name       string
		expected   Word
		expectedZr logic.Bit
		expectedNg logic.Bit
		controls   [6]Bit
		x          Word
		y          Word
	}{
		{
			"0",
			Word{},
			logic.I,
			logic.O,
			[6]Bit{1, 0, 1, 0, 1, 0},
			Word{1, 1, 1, 1, 1},
			Word{1, 1, 1},
		},
		{
			"1",
			Word{1},
			logic.O,
			logic.O,
			[6]Bit{1, 1, 1, 1, 1, 1},
			Word{1, 1, 1, 1, 1},
			Word{1, 1, 1},
		},
		{
			"-1",
			Word{1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1},
			logic.O,
			logic.I,
			[6]Bit{1, 1, 1, 0, 1, 0},
			Word{1, 1, 1, 1, 1},
			Word{1, 1, 1},
		},
		{
			"x",
			Word{0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1},
			logic.O,
			logic.I,
			[6]Bit{0, 0, 1, 1, 0, 0},
			Word{0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1},
			Word{1, 1, 1, 1, 1},
		},
		{
			"x",
			Word{1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0},
			logic.O,
			logic.O,
			[6]Bit{0, 0, 1, 1, 0, 0},
			Word{1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0},
			Word{1, 1, 1, 1, 1},
		},
		{
			"y",
			Word{0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1},
			logic.O,
			logic.I,
			[6]Bit{1, 1, 0, 0, 0, 0},
			Word{1, 1, 1, 1, 1},
			Word{0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1},
		},
		{
			"y",
			Word{1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0},
			logic.O,
			logic.O,
			[6]Bit{1, 1, 0, 0, 0, 0},
			Word{1, 1, 1, 1, 1},
			Word{1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0},
		},
		{
			"!x",
			Word{0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1},
			logic.O,
			logic.I,
			[6]Bit{0, 0, 1, 1, 0, 1},
			Word{1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0},
			Word{1, 1, 1, 1, 1},
		},
		{
			"!y",
			Word{0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1},
			logic.O,
			logic.I,
			[6]Bit{1, 1, 0, 0, 0, 1},
			Word{1, 1, 1, 1, 1},
			Word{1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0},
		},
		{
			"-x",
			Word{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
			logic.I,
			logic.O,
			[6]Bit{0, 0, 1, 1, 1, 1},
			Word{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
			Word{1, 1, 1, 1, 1},
		},
		{
			"-x",
			Word{1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1},
			logic.O,
			logic.I,
			[6]Bit{0, 0, 1, 1, 1, 1},
			Word{1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
			Word{1, 1, 1, 1, 1},
		},
		{
			"-y",
			Word{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
			logic.I,
			logic.O,
			[6]Bit{1, 1, 0, 0, 1, 1},
			Word{1, 1, 1, 1, 1},
			Word{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
		},
		{
			"-y",
			Word{1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1},
			logic.O,
			logic.I,
			[6]Bit{1, 1, 0, 0, 1, 1},
			Word{1, 1, 1, 1, 1},
			Word{1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
		},
		{
			"x+1",
			Word{0, 0, 1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
			logic.O,
			logic.O,
			[6]Bit{0, 1, 1, 1, 1, 1},
			Word{1, 1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
			Word{1, 1, 1, 1, 1},
		},
		{
			"y+1",
			Word{0, 0, 1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
			logic.O,
			logic.O,
			[6]Bit{1, 1, 0, 1, 1, 1},
			Word{1, 1, 1, 1, 1},
			Word{1, 1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
		},
		{
			"x-1",
			Word{1, 1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
			logic.O,
			logic.O,
			[6]Bit{0, 0, 1, 1, 1, 0},
			Word{0, 0, 1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
			Word{1, 1, 1, 1, 1},
		},
		{
			"y-1",
			Word{1, 1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
			logic.O,
			logic.O,
			[6]Bit{1, 1, 0, 0, 1, 0},
			Word{1, 1, 1, 1, 1},
			Word{0, 0, 1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
		},
		{
			"x+y",
			Word{1, 0, 0, 1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
			logic.O,
			logic.O,
			[6]Bit{0, 0, 0, 0, 1, 0},
			Word{1, 0, 1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
			Word{0, 0, 1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
		},
		{
			"x-y",
			Word{1, 0, 1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
			logic.O,
			logic.O,
			[6]Bit{0, 1, 0, 0, 1, 1},
			Word{1, 0, 0, 1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
			Word{0, 0, 1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
		},
		{
			"y-x",
			Word{1, 0, 1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
			logic.O,
			logic.O,
			[6]Bit{0, 0, 0, 1, 1, 1},
			Word{0, 0, 1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
			Word{1, 0, 0, 1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
		},
		{
			"x&y",
			Word{0, 0, 0, 0, 1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
			logic.O,
			logic.O,
			[6]Bit{0, 0, 0, 0, 0, 0},
			Word{0, 0, 1, 0, 1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
			Word{1, 0, 0, 1, 1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
		},
		{
			"x|y",
			Word{1, 0, 1, 1, 1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
			logic.O,
			logic.O,
			[6]Bit{0, 1, 0, 1, 0, 1},
			Word{0, 0, 1, 0, 1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
			Word{1, 0, 0, 1, 1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(fmt.Sprintf("out=%s", tt.name), func(t *testing.T) {
			actual, zr, ng := ALU(
				tt.x,
				tt.y,
				tt.controls[0],
				tt.controls[1],
				tt.controls[2],
				tt.controls[3],
				tt.controls[4],
				tt.controls[5],
			)
			if actual != tt.expected {
				t.Errorf("controls(%v): expected %v, actual %v", tt.controls, tt.expected, actual)
			}
			if zr != tt.expectedZr {
				t.Errorf("controls(%v): expected %v, actual %v", tt.controls, tt.expectedZr, zr)
			}
			if ng != tt.expectedNg {
				t.Errorf("controls(%v): expected %v, actual %v", tt.controls, tt.expectedNg, ng)
			}
		})
	}
}
