package arithmetic

import (
	logic "github.com/kazufusa/nand2tetris/01_Boolean_Logic"
)

type Bit = logic.Bit

// halfAdder adds two bits and return bits.
func halfAdder(a, b Bit) (sum, carry Bit) {
	return logic.Xor(a, b), logic.Not(logic.Nand(a, b))
}

// fullAdder adds three bits and return bits.
func fullAdder(a, b, c Bit) (sum, carry Bit) {
	return logic.Xor(logic.Xor(a, b), c), logic.Or(
		logic.Or(
			logic.And(a, b), logic.And(b, c),
		),
		logic.And(c, a),
	)
}

// Adder adds two n-bit numbers and return n-bit number. Adder ignores the
// overflow bit.
func Adder(a, b [16]Bit) [16]Bit {
	ret := [16]Bit{}
	var carry Bit
	ret[0], carry = halfAdder(a[0], b[0])
	ret[1], carry = fullAdder(a[1], b[1], carry)
	ret[2], carry = fullAdder(a[2], b[2], carry)
	ret[3], carry = fullAdder(a[3], b[3], carry)
	ret[4], carry = fullAdder(a[4], b[4], carry)
	ret[5], carry = fullAdder(a[5], b[5], carry)
	ret[6], carry = fullAdder(a[6], b[6], carry)
	ret[7], carry = fullAdder(a[7], b[7], carry)
	ret[8], carry = fullAdder(a[8], b[8], carry)
	ret[9], carry = fullAdder(a[9], b[9], carry)
	ret[10], carry = fullAdder(a[10], b[10], carry)
	ret[11], carry = fullAdder(a[11], b[11], carry)
	ret[12], carry = fullAdder(a[12], b[12], carry)
	ret[13], carry = fullAdder(a[13], b[13], carry)
	ret[14], carry = fullAdder(a[14], b[14], carry)
	ret[15], carry = fullAdder(a[15], b[15], carry)

	return ret
}

// Inc16 adds I to a input number. Inc16 ignores the overflow bit.
func Inc16(in [16]Bit) [16]Bit {
	return logic.Not16(
		Adder(
			logic.Not16(in),
			logic.Not16(
				logic.And16(
					in,
					logic.Not16(in),
				),
			),
		),
	)
}
