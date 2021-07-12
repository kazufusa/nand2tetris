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

func TestNot16(t *testing.T) {
	a := [16]Bit{O, I, O, I, O, O, O, O, O, O, O, O, O, O, O, O}
	e := [16]Bit{I, O, I, O, I, I, I, I, I, I, I, I, I, I, I, I}
	actual := Not16(a)
	if actual != e {
		t.Errorf("given(%v): expected %v, actual %v", a, e, actual)
	}
}

func TestAnd16(t *testing.T) {
	a := [16]Bit{O, O, I, I, O, O, O, O, O, O, O, O, O, O, O, O}
	b := [16]Bit{O, I, O, I, O, O, O, O, O, O, O, O, O, O, O, O}
	e := [16]Bit{O, O, O, I, O, O, O, O, O, O, O, O, O, O, O, O}
	actual := And16(a, b)
	if actual != e {
		t.Errorf("given(%v, %v): expected %v, actual %v", a, b, e, actual)
	}
}

func TestOr16(t *testing.T) {
	a := [16]Bit{O, O, I, I, O, O, O, O, O, O, O, O, O, O, O, O}
	b := [16]Bit{O, I, O, I, O, O, O, O, O, O, O, O, O, O, O, O}
	e := [16]Bit{O, I, I, I, O, O, O, O, O, O, O, O, O, O, O, O}
	actual := Or16(a, b)
	if actual != e {
		t.Errorf("given(%v, %v): expected %v, actual %v", a, b, e, actual)
	}
}

func TestMux16(t *testing.T) {
	a := [16]Bit{O, O, I, I, O, O, O, O, O, O, O, O, O, O, O, O}
	b := [16]Bit{O, I, O, I, O, O, O, O, O, O, O, O, O, O, O, O}
	s := O
	e := [16]Bit{O, O, I, I, O, O, O, O, O, O, O, O, O, O, O, O}
	actual := Mux16(a, b, s)
	if actual != e {
		t.Errorf("given(%v, %v,%v): expected %v, actual %v", a, b, s, e, actual)
	}

	s = I
	e = [16]Bit{O, I, O, I, O, O, O, O, O, O, O, O, O, O, O, O}
	actual = Mux16(a, b, s)
	if actual != e {
		t.Errorf("given(%v, %v,%v): expected %v, actual %v", a, b, s, e, actual)
	}
}

func TestOr8Way(t *testing.T) {
	var tests = []struct {
		expected Bit
		given    [8]Bit
	}{
		{expected: O, given: [8]Bit{0, 0, 0, 0, 0, 0, 0, 0}},
		{expected: I, given: [8]Bit{1, 0, 0, 0, 0, 0, 0, 0}},
		{expected: I, given: [8]Bit{0, 1, 0, 0, 0, 0, 0, 0}},
		{expected: I, given: [8]Bit{0, 0, 1, 0, 0, 0, 0, 0}},
		{expected: I, given: [8]Bit{0, 0, 0, 1, 0, 0, 0, 0}},
		{expected: I, given: [8]Bit{0, 0, 0, 0, 1, 0, 0, 0}},
		{expected: I, given: [8]Bit{0, 0, 0, 0, 0, 1, 0, 0}},
		{expected: I, given: [8]Bit{0, 0, 0, 0, 0, 0, 1, 0}},
		{expected: I, given: [8]Bit{0, 0, 0, 0, 0, 0, 0, 1}},
		{expected: I, given: [8]Bit{1, 1, 1, 1, 1, 1, 1, 1}},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(fmt.Sprintf("%v", tt.given), func(t *testing.T) {
			actual := Or8Way(tt.given)
			if actual != tt.expected {
				t.Errorf("given(%v): expected %v, actual %v", tt.given, tt.expected, actual)
			}
		})
	}
}

func TestMux4Way16(t *testing.T) {
	a := [16]Bit{I, I, I, I, I, I, I, I, I, I, I, I, I, I, I, I}
	b := [16]Bit{O, O, O, O, O, O, O, O, I, I, I, I, I, I, I, I}
	c := [16]Bit{I, I, I, I, I, I, I, I, O, O, O, O, O, O, O, O}
	d := [16]Bit{O, O, O, O, I, I, I, I, O, O, O, O, I, I, I, I}
	var tests = []struct {
		expected [16]Bit
		givenSel [2]Bit
	}{
		{a, [2]Bit{O, O}},
		{b, [2]Bit{I, O}},
		{c, [2]Bit{O, I}},
		{d, [2]Bit{I, I}},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(fmt.Sprintf("%v", tt.givenSel), func(t *testing.T) {
			actual := Mux4Way16(a, b, c, d, tt.givenSel)
			if actual != tt.expected {
				t.Errorf("given(%v): expected %v, actual %v", tt.givenSel, tt.expected, actual)
			}
		})
	}
}
func TestMux8Way16(t *testing.T) {
	a := [16]Bit{I, I, I, I, I, I, I, I, I, I, I, I, I, I, I, I}
	b := [16]Bit{O, O, O, O, O, O, O, O, I, I, I, I, I, I, I, I}
	c := [16]Bit{I, I, I, I, I, I, I, I, O, O, O, O, O, O, O, O}
	d := [16]Bit{O, O, O, O, I, I, I, I, O, O, O, O, I, I, I, I}
	e := [16]Bit{I, I, I, I, O, O, O, O, I, I, I, I, O, O, O, O}
	f := [16]Bit{O, O, I, I, O, O, I, I, O, O, I, I, O, O, I, I}
	g := [16]Bit{I, I, O, O, I, I, O, O, I, I, O, O, I, I, O, O}
	h := [16]Bit{O, I, O, I, O, I, O, I, O, I, O, I, O, I, O, I}
	var tests = []struct {
		expected [16]Bit
		givenSel [3]Bit
	}{
		{a, [3]Bit{O, O, O}},
		{b, [3]Bit{I, O, O}},
		{c, [3]Bit{O, I, O}},
		{d, [3]Bit{I, I, O}},
		{e, [3]Bit{O, O, I}},
		{f, [3]Bit{I, O, I}},
		{g, [3]Bit{O, I, I}},
		{h, [3]Bit{I, I, I}},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(fmt.Sprintf("%v", tt.givenSel), func(t *testing.T) {
			actual := Mux8Way16(a, b, c, d, e, f, g, h, tt.givenSel)
			if actual != tt.expected {
				t.Errorf("given(%v): expected %v, actual %v", tt.givenSel, tt.expected, actual)
			}
		})
	}
}

func TestDmux4Way(t *testing.T) {
	var tests = []struct {
		expected [4]Bit
		given    [2]Bit
	}{
		{[4]Bit{1, 0, 0, 0}, [2]Bit{O, O}},
		{[4]Bit{0, 1, 0, 0}, [2]Bit{I, O}},
		{[4]Bit{0, 0, 1, 0}, [2]Bit{O, I}},
		{[4]Bit{0, 0, 0, 1}, [2]Bit{I, I}},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(fmt.Sprintf("%v", tt.given), func(t *testing.T) {
			actual := Dmux4Way(I, tt.given)
			if actual != tt.expected {
				t.Errorf("given(%v): expected %v, actual %v", tt.given, tt.expected, actual)
			}
		})
	}
}
func TestDmux8Way(t *testing.T) {
	var tests = []struct {
		expected [8]Bit
		given    [3]Bit
	}{
		{[8]Bit{1, 0, 0, 0, 0, 0, 0, 0}, [3]Bit{O, O, O}},
		{[8]Bit{0, 1, 0, 0, 0, 0, 0, 0}, [3]Bit{I, O, O}},
		{[8]Bit{0, 0, 1, 0, 0, 0, 0, 0}, [3]Bit{O, I, O}},
		{[8]Bit{0, 0, 0, 1, 0, 0, 0, 0}, [3]Bit{I, I, O}},
		{[8]Bit{0, 0, 0, 0, 1, 0, 0, 0}, [3]Bit{O, O, I}},
		{[8]Bit{0, 0, 0, 0, 0, 1, 0, 0}, [3]Bit{I, O, I}},
		{[8]Bit{0, 0, 0, 0, 0, 0, 1, 0}, [3]Bit{O, I, I}},
		{[8]Bit{0, 0, 0, 0, 0, 0, 0, 1}, [3]Bit{I, I, I}},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(fmt.Sprintf("%v", tt.given), func(t *testing.T) {
			actual := Dmux8Way(I, tt.given)
			if actual != tt.expected {
				t.Errorf("given(%v): expected %v, actual %v", tt.given, tt.expected, actual)
			}
		})
	}
}
