package logic

type Bit = uint8

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

func Or8Way(in [8]Bit) Bit {
	return Or(
		Or(
			Or(in[0], in[1]),
			Or(in[2], in[3]),
		),
		Or(
			Or(in[4], in[5]),
			Or(in[6], in[7]),
		),
	)
}

func Mux4Way16(a, b, c, d [16]Bit, sel [2]Bit) [16]Bit {
	return [16]Bit{
		Or(
			Or(
				And(a[0], And(Not(sel[0]), Not(sel[1]))),
				And(b[0], And(sel[0], Not(sel[1]))),
			),
			Or(
				And(c[0], And(Not(sel[0]), sel[1])),
				And(d[0], And(sel[0], sel[1])),
			),
		),
		Or(
			Or(
				And(a[1], And(Not(sel[0]), Not(sel[1]))),
				And(b[1], And(sel[0], Not(sel[1]))),
			),
			Or(
				And(c[1], And(Not(sel[0]), sel[1])),
				And(d[1], And(sel[0], sel[1])),
			),
		),
		Or(
			Or(
				And(a[2], And(Not(sel[0]), Not(sel[1]))),
				And(b[2], And(sel[0], Not(sel[1]))),
			),
			Or(
				And(c[2], And(Not(sel[0]), sel[1])),
				And(d[2], And(sel[0], sel[1])),
			),
		),
		Or(
			Or(
				And(a[3], And(Not(sel[0]), Not(sel[1]))),
				And(b[3], And(sel[0], Not(sel[1]))),
			),
			Or(
				And(c[3], And(Not(sel[0]), sel[1])),
				And(d[3], And(sel[0], sel[1])),
			),
		),
		Or(
			Or(
				And(a[4], And(Not(sel[0]), Not(sel[1]))),
				And(b[4], And(sel[0], Not(sel[1]))),
			),
			Or(
				And(c[4], And(Not(sel[0]), sel[1])),
				And(d[4], And(sel[0], sel[1])),
			),
		),
		Or(
			Or(
				And(a[5], And(Not(sel[0]), Not(sel[1]))),
				And(b[5], And(sel[0], Not(sel[1]))),
			),
			Or(
				And(c[5], And(Not(sel[0]), sel[1])),
				And(d[5], And(sel[0], sel[1])),
			),
		),
		Or(
			Or(
				And(a[6], And(Not(sel[0]), Not(sel[1]))),
				And(b[6], And(sel[0], Not(sel[1]))),
			),
			Or(
				And(c[6], And(Not(sel[0]), sel[1])),
				And(d[6], And(sel[0], sel[1])),
			),
		),
		Or(
			Or(
				And(a[7], And(Not(sel[0]), Not(sel[1]))),
				And(b[7], And(sel[0], Not(sel[1]))),
			),
			Or(
				And(c[7], And(Not(sel[0]), sel[1])),
				And(d[7], And(sel[0], sel[1])),
			),
		),
		Or(
			Or(
				And(a[8], And(Not(sel[0]), Not(sel[1]))),
				And(b[8], And(sel[0], Not(sel[1]))),
			),
			Or(
				And(c[8], And(Not(sel[0]), sel[1])),
				And(d[8], And(sel[0], sel[1])),
			),
		),
		Or(
			Or(
				And(a[9], And(Not(sel[0]), Not(sel[1]))),
				And(b[9], And(sel[0], Not(sel[1]))),
			),
			Or(
				And(c[9], And(Not(sel[0]), sel[1])),
				And(d[9], And(sel[0], sel[1])),
			),
		),
		Or(
			Or(
				And(a[10], And(Not(sel[0]), Not(sel[1]))),
				And(b[10], And(sel[0], Not(sel[1]))),
			),
			Or(
				And(c[10], And(Not(sel[0]), sel[1])),
				And(d[10], And(sel[0], sel[1])),
			),
		),
		Or(
			Or(
				And(a[11], And(Not(sel[0]), Not(sel[1]))),
				And(b[11], And(sel[0], Not(sel[1]))),
			),
			Or(
				And(c[11], And(Not(sel[0]), sel[1])),
				And(d[11], And(sel[0], sel[1])),
			),
		),
		Or(
			Or(
				And(a[12], And(Not(sel[0]), Not(sel[1]))),
				And(b[12], And(sel[0], Not(sel[1]))),
			),
			Or(
				And(c[12], And(Not(sel[0]), sel[1])),
				And(d[12], And(sel[0], sel[1])),
			),
		),
		Or(
			Or(
				And(a[13], And(Not(sel[0]), Not(sel[1]))),
				And(b[13], And(sel[0], Not(sel[1]))),
			),
			Or(
				And(c[13], And(Not(sel[0]), sel[1])),
				And(d[13], And(sel[0], sel[1])),
			),
		),
		Or(
			Or(
				And(a[14], And(Not(sel[0]), Not(sel[1]))),
				And(b[14], And(sel[0], Not(sel[1]))),
			),
			Or(
				And(c[14], And(Not(sel[0]), sel[1])),
				And(d[14], And(sel[0], sel[1])),
			),
		),
		Or(
			Or(
				And(a[15], And(Not(sel[0]), Not(sel[1]))),
				And(b[15], And(sel[0], Not(sel[1]))),
			),
			Or(
				And(c[15], And(Not(sel[0]), sel[1])),
				And(d[15], And(sel[0], sel[1])),
			),
		),
	}
}

func Mux8Way16(a, b, c, d, e, f, g, h [16]Bit, sel [3]Bit) [16]Bit {
	return [16]Bit{
		Or(
			Or(
				Or(
					And(a[0], And(And(Not(sel[0]), Not(sel[1])), Not(sel[2]))),
					And(b[0], And(And(sel[0], Not(sel[1])), Not(sel[2]))),
				),
				Or(
					And(c[0], And(And(Not(sel[0]), sel[1]), Not(sel[2]))),
					And(d[0], And(And(sel[0], sel[1]), Not(sel[2]))),
				),
			),
			Or(
				Or(
					And(e[0], And(And(Not(sel[0]), Not(sel[1])), sel[2])),
					And(f[0], And(And(sel[0], Not(sel[1])), sel[2])),
				),
				Or(
					And(g[0], And(And(Not(sel[0]), sel[1]), sel[2])),
					And(h[0], And(And(sel[0], sel[1]), sel[2])),
				),
			),
		),
		Or(
			Or(
				Or(
					And(a[1], And(And(Not(sel[0]), Not(sel[1])), Not(sel[2]))),
					And(b[1], And(And(sel[0], Not(sel[1])), Not(sel[2]))),
				),
				Or(
					And(c[1], And(And(Not(sel[0]), sel[1]), Not(sel[2]))),
					And(d[1], And(And(sel[0], sel[1]), Not(sel[2]))),
				),
			),
			Or(
				Or(
					And(e[1], And(And(Not(sel[0]), Not(sel[1])), sel[2])),
					And(f[1], And(And(sel[0], Not(sel[1])), sel[2])),
				),
				Or(
					And(g[1], And(And(Not(sel[0]), sel[1]), sel[2])),
					And(h[1], And(And(sel[0], sel[1]), sel[2])),
				),
			),
		),
		Or(
			Or(
				Or(
					And(a[2], And(And(Not(sel[0]), Not(sel[1])), Not(sel[2]))),
					And(b[2], And(And(sel[0], Not(sel[1])), Not(sel[2]))),
				),
				Or(
					And(c[2], And(And(Not(sel[0]), sel[1]), Not(sel[2]))),
					And(d[2], And(And(sel[0], sel[1]), Not(sel[2]))),
				),
			),
			Or(
				Or(
					And(e[2], And(And(Not(sel[0]), Not(sel[1])), sel[2])),
					And(f[2], And(And(sel[0], Not(sel[1])), sel[2])),
				),
				Or(
					And(g[2], And(And(Not(sel[0]), sel[1]), sel[2])),
					And(h[2], And(And(sel[0], sel[1]), sel[2])),
				),
			),
		),
		Or(
			Or(
				Or(
					And(a[3], And(And(Not(sel[0]), Not(sel[1])), Not(sel[2]))),
					And(b[3], And(And(sel[0], Not(sel[1])), Not(sel[2]))),
				),
				Or(
					And(c[3], And(And(Not(sel[0]), sel[1]), Not(sel[2]))),
					And(d[3], And(And(sel[0], sel[1]), Not(sel[2]))),
				),
			),
			Or(
				Or(
					And(e[3], And(And(Not(sel[0]), Not(sel[1])), sel[2])),
					And(f[3], And(And(sel[0], Not(sel[1])), sel[2])),
				),
				Or(
					And(g[3], And(And(Not(sel[0]), sel[1]), sel[2])),
					And(h[3], And(And(sel[0], sel[1]), sel[2])),
				),
			),
		),
		Or(
			Or(
				Or(
					And(a[4], And(And(Not(sel[0]), Not(sel[1])), Not(sel[2]))),
					And(b[4], And(And(sel[0], Not(sel[1])), Not(sel[2]))),
				),
				Or(
					And(c[4], And(And(Not(sel[0]), sel[1]), Not(sel[2]))),
					And(d[4], And(And(sel[0], sel[1]), Not(sel[2]))),
				),
			),
			Or(
				Or(
					And(e[4], And(And(Not(sel[0]), Not(sel[1])), sel[2])),
					And(f[4], And(And(sel[0], Not(sel[1])), sel[2])),
				),
				Or(
					And(g[4], And(And(Not(sel[0]), sel[1]), sel[2])),
					And(h[4], And(And(sel[0], sel[1]), sel[2])),
				),
			),
		),
		Or(
			Or(
				Or(
					And(a[5], And(And(Not(sel[0]), Not(sel[1])), Not(sel[2]))),
					And(b[5], And(And(sel[0], Not(sel[1])), Not(sel[2]))),
				),
				Or(
					And(c[5], And(And(Not(sel[0]), sel[1]), Not(sel[2]))),
					And(d[5], And(And(sel[0], sel[1]), Not(sel[2]))),
				),
			),
			Or(
				Or(
					And(e[5], And(And(Not(sel[0]), Not(sel[1])), sel[2])),
					And(f[5], And(And(sel[0], Not(sel[1])), sel[2])),
				),
				Or(
					And(g[5], And(And(Not(sel[0]), sel[1]), sel[2])),
					And(h[5], And(And(sel[0], sel[1]), sel[2])),
				),
			),
		),
		Or(
			Or(
				Or(
					And(a[6], And(And(Not(sel[0]), Not(sel[1])), Not(sel[2]))),
					And(b[6], And(And(sel[0], Not(sel[1])), Not(sel[2]))),
				),
				Or(
					And(c[6], And(And(Not(sel[0]), sel[1]), Not(sel[2]))),
					And(d[6], And(And(sel[0], sel[1]), Not(sel[2]))),
				),
			),
			Or(
				Or(
					And(e[6], And(And(Not(sel[0]), Not(sel[1])), sel[2])),
					And(f[6], And(And(sel[0], Not(sel[1])), sel[2])),
				),
				Or(
					And(g[6], And(And(Not(sel[0]), sel[1]), sel[2])),
					And(h[6], And(And(sel[0], sel[1]), sel[2])),
				),
			),
		),
		Or(
			Or(
				Or(
					And(a[7], And(And(Not(sel[0]), Not(sel[1])), Not(sel[2]))),
					And(b[7], And(And(sel[0], Not(sel[1])), Not(sel[2]))),
				),
				Or(
					And(c[7], And(And(Not(sel[0]), sel[1]), Not(sel[2]))),
					And(d[7], And(And(sel[0], sel[1]), Not(sel[2]))),
				),
			),
			Or(
				Or(
					And(e[7], And(And(Not(sel[0]), Not(sel[1])), sel[2])),
					And(f[7], And(And(sel[0], Not(sel[1])), sel[2])),
				),
				Or(
					And(g[7], And(And(Not(sel[0]), sel[1]), sel[2])),
					And(h[7], And(And(sel[0], sel[1]), sel[2])),
				),
			),
		),
		Or(
			Or(
				Or(
					And(a[8], And(And(Not(sel[0]), Not(sel[1])), Not(sel[2]))),
					And(b[8], And(And(sel[0], Not(sel[1])), Not(sel[2]))),
				),
				Or(
					And(c[8], And(And(Not(sel[0]), sel[1]), Not(sel[2]))),
					And(d[8], And(And(sel[0], sel[1]), Not(sel[2]))),
				),
			),
			Or(
				Or(
					And(e[8], And(And(Not(sel[0]), Not(sel[1])), sel[2])),
					And(f[8], And(And(sel[0], Not(sel[1])), sel[2])),
				),
				Or(
					And(g[8], And(And(Not(sel[0]), sel[1]), sel[2])),
					And(h[8], And(And(sel[0], sel[1]), sel[2])),
				),
			),
		),
		Or(
			Or(
				Or(
					And(a[9], And(And(Not(sel[0]), Not(sel[1])), Not(sel[2]))),
					And(b[9], And(And(sel[0], Not(sel[1])), Not(sel[2]))),
				),
				Or(
					And(c[9], And(And(Not(sel[0]), sel[1]), Not(sel[2]))),
					And(d[9], And(And(sel[0], sel[1]), Not(sel[2]))),
				),
			),
			Or(
				Or(
					And(e[9], And(And(Not(sel[0]), Not(sel[1])), sel[2])),
					And(f[9], And(And(sel[0], Not(sel[1])), sel[2])),
				),
				Or(
					And(g[9], And(And(Not(sel[0]), sel[1]), sel[2])),
					And(h[9], And(And(sel[0], sel[1]), sel[2])),
				),
			),
		),
		Or(
			Or(
				Or(
					And(a[10], And(And(Not(sel[0]), Not(sel[1])), Not(sel[2]))),
					And(b[10], And(And(sel[0], Not(sel[1])), Not(sel[2]))),
				),
				Or(
					And(c[10], And(And(Not(sel[0]), sel[1]), Not(sel[2]))),
					And(d[10], And(And(sel[0], sel[1]), Not(sel[2]))),
				),
			),
			Or(
				Or(
					And(e[10], And(And(Not(sel[0]), Not(sel[1])), sel[2])),
					And(f[10], And(And(sel[0], Not(sel[1])), sel[2])),
				),
				Or(
					And(g[10], And(And(Not(sel[0]), sel[1]), sel[2])),
					And(h[10], And(And(sel[0], sel[1]), sel[2])),
				),
			),
		),
		Or(
			Or(
				Or(
					And(a[11], And(And(Not(sel[0]), Not(sel[1])), Not(sel[2]))),
					And(b[11], And(And(sel[0], Not(sel[1])), Not(sel[2]))),
				),
				Or(
					And(c[11], And(And(Not(sel[0]), sel[1]), Not(sel[2]))),
					And(d[11], And(And(sel[0], sel[1]), Not(sel[2]))),
				),
			),
			Or(
				Or(
					And(e[11], And(And(Not(sel[0]), Not(sel[1])), sel[2])),
					And(f[11], And(And(sel[0], Not(sel[1])), sel[2])),
				),
				Or(
					And(g[11], And(And(Not(sel[0]), sel[1]), sel[2])),
					And(h[11], And(And(sel[0], sel[1]), sel[2])),
				),
			),
		),
		Or(
			Or(
				Or(
					And(a[12], And(And(Not(sel[0]), Not(sel[1])), Not(sel[2]))),
					And(b[12], And(And(sel[0], Not(sel[1])), Not(sel[2]))),
				),
				Or(
					And(c[12], And(And(Not(sel[0]), sel[1]), Not(sel[2]))),
					And(d[12], And(And(sel[0], sel[1]), Not(sel[2]))),
				),
			),
			Or(
				Or(
					And(e[12], And(And(Not(sel[0]), Not(sel[1])), sel[2])),
					And(f[12], And(And(sel[0], Not(sel[1])), sel[2])),
				),
				Or(
					And(g[12], And(And(Not(sel[0]), sel[1]), sel[2])),
					And(h[12], And(And(sel[0], sel[1]), sel[2])),
				),
			),
		),
		Or(
			Or(
				Or(
					And(a[13], And(And(Not(sel[0]), Not(sel[1])), Not(sel[2]))),
					And(b[13], And(And(sel[0], Not(sel[1])), Not(sel[2]))),
				),
				Or(
					And(c[13], And(And(Not(sel[0]), sel[1]), Not(sel[2]))),
					And(d[13], And(And(sel[0], sel[1]), Not(sel[2]))),
				),
			),
			Or(
				Or(
					And(e[13], And(And(Not(sel[0]), Not(sel[1])), sel[2])),
					And(f[13], And(And(sel[0], Not(sel[1])), sel[2])),
				),
				Or(
					And(g[13], And(And(Not(sel[0]), sel[1]), sel[2])),
					And(h[13], And(And(sel[0], sel[1]), sel[2])),
				),
			),
		),
		Or(
			Or(
				Or(
					And(a[14], And(And(Not(sel[0]), Not(sel[1])), Not(sel[2]))),
					And(b[14], And(And(sel[0], Not(sel[1])), Not(sel[2]))),
				),
				Or(
					And(c[14], And(And(Not(sel[0]), sel[1]), Not(sel[2]))),
					And(d[14], And(And(sel[0], sel[1]), Not(sel[2]))),
				),
			),
			Or(
				Or(
					And(e[14], And(And(Not(sel[0]), Not(sel[1])), sel[2])),
					And(f[14], And(And(sel[0], Not(sel[1])), sel[2])),
				),
				Or(
					And(g[14], And(And(Not(sel[0]), sel[1]), sel[2])),
					And(h[14], And(And(sel[0], sel[1]), sel[2])),
				),
			),
		),
		Or(
			Or(
				Or(
					And(a[15], And(And(Not(sel[0]), Not(sel[1])), Not(sel[2]))),
					And(b[15], And(And(sel[0], Not(sel[1])), Not(sel[2]))),
				),
				Or(
					And(c[15], And(And(Not(sel[0]), sel[1]), Not(sel[2]))),
					And(d[15], And(And(sel[0], sel[1]), Not(sel[2]))),
				),
			),
			Or(
				Or(
					And(e[15], And(And(Not(sel[0]), Not(sel[1])), sel[2])),
					And(f[15], And(And(sel[0], Not(sel[1])), sel[2])),
				),
				Or(
					And(g[15], And(And(Not(sel[0]), sel[1]), sel[2])),
					And(h[15], And(And(sel[0], sel[1]), sel[2])),
				),
			),
		),
	}
}

func Dmux4Way(in Bit, sel [2]Bit) [4]Bit {
	return [4]Bit{
		And(in, And(Not(sel[0]), Not(sel[1]))),
		And(in, And(sel[0], Not(sel[1]))),
		And(in, And(Not(sel[0]), sel[1])),
		And(in, And(sel[0], sel[1])),
	}
}

func Dmux8Way(in Bit, sel [3]Bit) [8]Bit {
	return [8]Bit{
		And(in, And(And(Not(sel[0]), Not(sel[1])), Not(sel[2]))),
		And(in, And(And(sel[0], Not(sel[1])), Not(sel[2]))),
		And(in, And(And(Not(sel[0]), sel[1]), Not(sel[2]))),
		And(in, And(And(sel[0], sel[1]), Not(sel[2]))),
		And(in, And(And(Not(sel[0]), Not(sel[1])), sel[2])),
		And(in, And(And(sel[0], Not(sel[1])), sel[2])),
		And(in, And(And(Not(sel[0]), sel[1]), sel[2])),
		And(in, And(And(sel[0], sel[1]), sel[2])),
	}
}
