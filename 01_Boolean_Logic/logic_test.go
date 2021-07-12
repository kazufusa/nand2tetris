package logic

import (
	"fmt"
	"testing"
)

func TestNand(t *testing.T) {
	var tests = []struct {
		expected Bit
		given    [2]Bit
	}{
		{I, [2]Bit{O, O}},
		{I, [2]Bit{O, I}},
		{I, [2]Bit{I, O}},
		{O, [2]Bit{I, I}},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(fmt.Sprintf("%v", tt.given), func(t *testing.T) {
			actual := Nand(tt.given[0], tt.given[1])
			if actual != tt.expected {
				t.Errorf("given(%v): expected %v, actual %v", tt.given, tt.expected, actual)
			}
		})
	}
}

func TestAnd(t *testing.T) {
	var tests = []struct {
		expected Bit
		given    [2]Bit
	}{
		{O, [2]Bit{O, O}},
		{O, [2]Bit{O, I}},
		{O, [2]Bit{I, O}},
		{I, [2]Bit{I, I}},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(fmt.Sprintf("%v", tt.given), func(t *testing.T) {
			actual := And(tt.given[0], tt.given[1])
			if actual != tt.expected {
				t.Errorf("given(%v): expected %v, actual %v", tt.given, tt.expected, actual)
			}
		})
	}
}

func TestOr(t *testing.T) {
	var tests = []struct {
		expected Bit
		given    [2]Bit
	}{
		{O, [2]Bit{O, O}},
		{I, [2]Bit{O, I}},
		{I, [2]Bit{I, O}},
		{I, [2]Bit{I, I}},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(fmt.Sprintf("%v", tt.given), func(t *testing.T) {
			actual := Or(tt.given[0], tt.given[1])
			if actual != tt.expected {
				t.Errorf("given(%v): expected %v, actual %v", tt.given, tt.expected, actual)
			}
		})
	}
}

func TestNot(t *testing.T) {
	var tests = []struct {
		expected Bit
		given    Bit
	}{
		{O, I},
		{I, O},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(fmt.Sprintf("%v", tt.given), func(t *testing.T) {
			actual := Not(tt.given)
			if actual != tt.expected {
				t.Errorf("given(%v): expected %v, actual %v", tt.given, tt.expected, actual)
			}
		})
	}
}

func TestXor(t *testing.T) {
	var tests = []struct {
		expected Bit
		given    [2]Bit
	}{
		{O, [2]Bit{O, O}},
		{I, [2]Bit{O, I}},
		{I, [2]Bit{I, O}},
		{O, [2]Bit{I, I}},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(fmt.Sprintf("%v", tt.given), func(t *testing.T) {
			actual := Xor(tt.given[0], tt.given[1])
			if actual != tt.expected {
				t.Errorf("given(%v): expected %v, actual %v", tt.given, tt.expected, actual)
			}
		})
	}
}

func TestMultiplexer(t *testing.T) {
	var tests = []struct {
		expected Bit
		given    [3]Bit
	}{
		{O, [3]Bit{O, O, O}},
		{O, [3]Bit{O, I, O}},
		{I, [3]Bit{I, O, O}},
		{I, [3]Bit{I, I, O}},

		{O, [3]Bit{O, O, I}},
		{I, [3]Bit{O, I, I}},
		{O, [3]Bit{I, O, I}},
		{I, [3]Bit{I, I, I}},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(fmt.Sprintf("%v", tt.given), func(t *testing.T) {
			actual := Mux(tt.given[0], tt.given[1], tt.given[2])
			if actual != tt.expected {
				t.Errorf("given(%v): expected %v, actual %v", tt.given, tt.expected, actual)
			}
		})
	}
}

func TestDemultiplexer(t *testing.T) {
	var tests = []struct {
		expected [2]Bit // {a, b}
		given    [2]Bit // {in, sel}
	}{
		{[2]Bit{O, O}, [2]Bit{O, O}},
		{[2]Bit{O, O}, [2]Bit{O, I}},
		{[2]Bit{I, O}, [2]Bit{I, O}},
		{[2]Bit{O, I}, [2]Bit{I, I}},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(fmt.Sprintf("%v", tt.given), func(t *testing.T) {
			actual := DMux(tt.given[0], tt.given[1])
			if actual != tt.expected {
				t.Errorf("given(%v): expected %v, actual %v", tt.given, tt.expected, actual)
			}
		})
	}
}
