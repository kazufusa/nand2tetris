package arithmetic

import (
	"fmt"
	"testing"
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
		expected [16]Bit
		givenA   [16]Bit
		givenB   [16]Bit
	}{
		{
			[16]Bit{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
			[16]Bit{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
			[16]Bit{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
		},
		{
			[16]Bit{0, 1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
			[16]Bit{1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
			[16]Bit{1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
		},
		{
			[16]Bit{1, 0, 1, 1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
			[16]Bit{1, 1, 0, 1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
			[16]Bit{0, 1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
		},
		{
			[16]Bit{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 0, 1, 1},
			[16]Bit{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 1, 0, 1},
			[16]Bit{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 0, 0},
		},
		{
			[16]Bit{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 0, 0, 1},
			[16]Bit{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 1, 0, 1},
			[16]Bit{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 1, 1},
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
		expected [16]Bit
		given    [16]Bit
	}{
		{
			[16]Bit{1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
			[16]Bit{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
		},
		{
			[16]Bit{0, 1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
			[16]Bit{1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
		},
		{
			[16]Bit{1, 0, 1, 1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
			[16]Bit{0, 0, 1, 1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
		},
		{
			[16]Bit{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
			[16]Bit{1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1},
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
