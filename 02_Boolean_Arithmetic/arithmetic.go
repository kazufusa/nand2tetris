package arithmetic

import (
	logic "github.com/kazufusa/nand2tetris/01_Boolean_Logic"
)

type Bit = logic.Bit

type Word = [16]Bit

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
func Adder(a, b Word) Word {
	ret := Word{}
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
func Inc16(in Word) Word {
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

// ALU is an implementation of Hack ALU. ALU ignores the overflow bit.
// Input:
//  x  Word data input
//  y  Word data input
//  zx Zero the x input
//  nx Negate the x input
//  zy Zero the y input
//  ny Negate the y input
//  f  if f == 1 out = add(x,y) else out = and(x,y)
//  no Negate the out output
// Output:
//  out Word output
//  zr  if out==0 zr=1 else zr=0
//  ng  if out<0 ng=1 else ng=0
// Function:
//  if zx z=0
//  if nx x=!x
//  if zy y=0
//  if nz z=!z
//  if f out=x+y
//  else out=x&y
//  if no out=!out
//  if out==0 zr=1 else zr=0
//  if out<0 ng=1 else ng=0
func ALU(x, y Word, zx, nx, zy, ny, f, no Bit) (out Word, zr, ng Bit) {
	// Zero x if zx
	x = logic.Mux16(x, logic.And16(x, logic.Not16(x)), zx)
	// Not x if nx
	x = logic.Mux16(x, logic.Not16(x), nx)
	// Zero y if zy
	y = logic.Mux16(y, logic.And16(y, logic.Not16(y)), zy)
	// Not y if ny
	y = logic.Mux16(y, logic.Not16(y), ny)

	// execute f
	out = logic.Mux16(logic.And16(x, y), Adder(x, y), f)
	// Not out if no
	out = logic.Mux16(out, logic.Not16(out), no)
	zr = logic.Not(logic.Or(
		logic.Or(
			logic.Or(logic.Or(out[0], out[1]), logic.Or(out[2], out[3])),
			logic.Or(logic.Or(out[4], out[5]), logic.Or(out[6], out[7])),
		),
		logic.Or(
			logic.Or(logic.Or(out[8], out[9]), logic.Or(out[10], out[11])),
			logic.Or(logic.Or(out[12], out[13]), logic.Or(out[14], out[15])),
		),
	))
	ng = out[15]
	return
}
