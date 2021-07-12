package logic

type Bit uint8

const (
	O Bit = iota
	I
)

func Nand(a, b Bit) Bit {
	if a == I && b == I {
		return O
	}
	return I
}

func And(a, b Bit) Bit {
	return Nand(Nand(a, b), Nand(a, b))
}

func Not(a Bit) Bit {
	return Nand(a, a)
}

func Or(a, b Bit) Bit {
	return Nand(Not(a), Not(b))
}

func Xor(a, b Bit) Bit {
	return Or(And(Not(a), b), And(a, Not(b)))
}

func Mux(a, b, sel Bit) Bit {
	return Nand(Nand(a, Not(sel)), Nand(b, sel))
}

func DMux(in, sel Bit) [2]Bit {
	return [2]Bit{And(in, Not(sel)), And(in, sel)}
}
