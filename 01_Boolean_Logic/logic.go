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

func Not16(in [16]Bit) [16]Bit {
	return [16]Bit{
		Not(in[0]),
		Not(in[1]),
		Not(in[2]),
		Not(in[3]),
		Not(in[4]),
		Not(in[5]),
		Not(in[6]),
		Not(in[7]),
		Not(in[8]),
		Not(in[9]),
		Not(in[10]),
		Not(in[11]),
		Not(in[12]),
		Not(in[13]),
		Not(in[14]),
		Not(in[15]),
	}
}

func And16(a, b [16]Bit) [16]Bit {
	return [16]Bit{
		And(a[0], b[0]),
		And(a[1], b[1]),
		And(a[2], b[2]),
		And(a[3], b[3]),
		And(a[4], b[4]),
		And(a[5], b[5]),
		And(a[6], b[6]),
		And(a[7], b[7]),
		And(a[8], b[8]),
		And(a[9], b[9]),
		And(a[10], b[10]),
		And(a[11], b[11]),
		And(a[12], b[12]),
		And(a[13], b[13]),
		And(a[14], b[14]),
		And(a[15], b[15]),
	}
}

func Or16(a, b [16]Bit) [16]Bit {
	return [16]Bit{
		Or(a[0], b[0]),
		Or(a[1], b[1]),
		Or(a[2], b[2]),
		Or(a[3], b[3]),
		Or(a[4], b[4]),
		Or(a[5], b[5]),
		Or(a[6], b[6]),
		Or(a[7], b[7]),
		Or(a[8], b[8]),
		Or(a[9], b[9]),
		Or(a[10], b[10]),
		Or(a[11], b[11]),
		Or(a[12], b[12]),
		Or(a[13], b[13]),
		Or(a[14], b[14]),
		Or(a[15], b[15]),
	}
}

func Mux16(a, b [16]Bit, sel Bit) [16]Bit {
	return [16]Bit{
		Mux(a[0], b[0], sel),
		Mux(a[1], b[1], sel),
		Mux(a[2], b[2], sel),
		Mux(a[3], b[3], sel),
		Mux(a[4], b[4], sel),
		Mux(a[5], b[5], sel),
		Mux(a[6], b[6], sel),
		Mux(a[7], b[7], sel),
		Mux(a[8], b[8], sel),
		Mux(a[9], b[9], sel),
		Mux(a[10], b[10], sel),
		Mux(a[11], b[11], sel),
		Mux(a[12], b[12], sel),
		Mux(a[13], b[13], sel),
		Mux(a[14], b[14], sel),
		Mux(a[15], b[15], sel),
	}
}
