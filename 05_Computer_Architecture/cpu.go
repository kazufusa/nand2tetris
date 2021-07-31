package computer

import (
	logic "github.com/kazufusa/nand2tetris/01_Boolean_Logic"
	arithmetic "github.com/kazufusa/nand2tetris/02_Boolean_Arithmetic"
	memory "github.com/kazufusa/nand2tetris/03_Memory"
)

type Bit = logic.Bit

type Word = [16]Bit

type Address = [15]Bit

type CPU struct {
	a, d   memory.Register
	pc     memory.PC
	clock  *memory.Clock
	outM   Word
	writeM Bit
}

func NewCPU() CPU {
	c := memory.Clock(0)
	return CPU{
		a:     memory.NewRegister(&c),
		d:     memory.NewRegister(&c),
		pc:    memory.NewPC(&c),
		clock: &c,
	}
}

func (cpu *CPU) Fetch(inM, inst Word, reset Bit) (outM Word, writeM Bit, addressM, pc [15]Bit) {
	// A, D, PC
	i, a, cccccc, ddd, jjj := cpu.decode(inst)

	outM, zr, ng := arithmetic.ALU(
		cpu.d.Apply(logic.O, inst),
		logic.Mux16(cpu.a.Apply(logic.O, inst), inM, a),
		cccccc[5], cccccc[4], cccccc[3], cccccc[2], cccccc[1], cccccc[0],
	)

	regA := cpu.a.Apply(logic.Or(logic.Not(i), logic.And(i, ddd[2])), logic.Mux16(inst, outM, i))

	_ = cpu.d.Apply(logic.And(i, ddd[1]), outM)

	writeM = logic.And(i, ddd[0])

	_ = cpu.pc.Apply(
		regA,
		logic.And(i, cpu.jump(zr, ng, jjj)),
		logic.Or(logic.Not(i), logic.Not(cpu.jump(zr, ng, jjj))),
		reset,
	)

	cpu.clock.Progress()

	_addressM := cpu.a.Apply(logic.O, inM)
	copy(addressM[:], _addressM[:15])

	_pc := cpu.pc.Apply(inM, logic.O, logic.O, logic.O)
	copy(pc[:], _pc[:15])
	return
}

func (c *CPU) decode(inst Word) (i Bit, a Bit, cccccc []Bit, ddd, jjj []Bit) {
	return inst[15],
		inst[12],
		inst[6:12],
		inst[3:6],
		inst[0:3]
}

func (c *CPU) jump(zr, ng Bit, jjj []Bit) Bit {
	JGT := logic.And(
		logic.And(logic.And(logic.Not(jjj[2]), logic.Not(jjj[1])), jjj[0]),
		logic.And(logic.Not(zr), logic.Not(ng)),
	)
	JEQ := logic.And(
		logic.And(logic.And(logic.Not(jjj[2]), jjj[1]), logic.Not(jjj[0])),
		zr,
	)
	JGE := logic.And(
		logic.And(logic.And(logic.Not(jjj[2]), jjj[1]), jjj[0]),
		logic.Or(zr, logic.Not(ng)),
	)
	JLT := logic.And(
		logic.And(logic.And(jjj[2], logic.Not(jjj[1])), logic.Not(jjj[0])),
		ng,
	)
	JNE := logic.And(
		logic.And(logic.And(jjj[2], logic.Not(jjj[1])), jjj[0]),
		logic.Not(zr),
	)
	JLE := logic.And(
		logic.And(logic.And(jjj[2], jjj[1]), logic.Not(jjj[0])),
		logic.Or(zr, ng),
	)
	JMP := logic.And(logic.And(jjj[2], jjj[1]), jjj[0])
	return logic.Or(
		logic.Or(
			logic.Or(JGT, JEQ),
			logic.Or(JGE, JLT),
		),
		logic.Or(
			logic.Or(JNE, JLE),
			JMP,
		),
	)
}
