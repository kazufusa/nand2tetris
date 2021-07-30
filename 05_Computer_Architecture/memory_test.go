package computer

import (
	"testing"

	logic "github.com/kazufusa/nand2tetris/01_Boolean_Logic"
	memory "github.com/kazufusa/nand2tetris/03_Memory"
	"github.com/stretchr/testify/assert"
)

type TestScreen struct {
	rams [2]memory.RAM4096
}

func NewTestScreen(clock *memory.Clock) TestScreen {
	return TestScreen{rams: [2]memory.RAM4096{
		memory.NewRAM4096(clock),
		memory.NewRAM4096(clock),
	}}
}

func (s *TestScreen) Fetch(in Word, load Bit, addr [13]Bit) Word {
	subAddr := [12]logic.Bit{
		addr[0], addr[1], addr[2],
		addr[3], addr[4], addr[5],
		addr[6], addr[7], addr[8],
		addr[9], addr[10], addr[11],
	}
	return logic.Mux16(
		s.rams[0].Apply(in, logic.And(load, logic.Not(addr[12])), subAddr),
		s.rams[1].Apply(in, logic.And(load, addr[12]), subAddr),
		addr[12],
	)
}

type TestKeyboard struct {
}

func (k *TestKeyboard) Fetch() Word {
	return wkb
}

var (
	w0    = Word{}
	w1    = Word{1}
	w2    = Word{0, 1}
	w3    = Word{1, 1}
	w4    = Word{0, 0, 1}
	w5    = Word{1, 0, 1}
	w6    = Word{0, 1, 1}
	w7    = Word{1, 1, 1}
	w8    = Word{0, 0, 0, 1}
	wkb   = Word{1, 0, 1, 0, 1, 0, 1}
	addr0 = [15]logic.Bit{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}
	addr1 = [15]logic.Bit{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 0, 0}
	addr2 = [15]logic.Bit{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 0}
	addr3 = [15]logic.Bit{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 1, 0}
	addr4 = [15]logic.Bit{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1}
	addr5 = [15]logic.Bit{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 0, 1}
	addr6 = [15]logic.Bit{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 1}

	_ IScreen   = (*TestScreen)(nil)
	_ IKeyboard = (*TestKeyboard)(nil)
)

func TestMemory(t *testing.T) {
	clock := memory.Clock(0)
	sc := NewTestScreen(&clock)
	kb := TestKeyboard{}
	mem := NewMemory(&clock, &sc, &kb)

	mem.Fetch(w1, logic.I, addr0)
	clock.Progress()
	mem.Fetch(w0, logic.O, addr0)
	assert.Equal(t, w1, mem.Fetch(w0, logic.O, addr0), "invalid ram4k[0]")
	assert.Equal(t, w0, mem.Fetch(w0, logic.O, addr1), "invalid ram4k[1]")
	assert.Equal(t, w0, mem.Fetch(w0, logic.O, addr2), "invalid ram4k[2]")
	assert.Equal(t, w0, mem.Fetch(w0, logic.O, addr3), "invalid ram4k[3]")
	assert.Equal(t, w0, mem.Fetch(w0, logic.O, addr4), "invalid screen.ram4k[0]")
	assert.Equal(t, w0, mem.Fetch(w0, logic.O, addr5), "invalid screen.ram4k[1]")
	assert.Equal(t, wkb, mem.Fetch(w0, logic.O, addr6), "inbalid keyboard")

	mem.Fetch(w2, logic.I, addr1)
	clock.Progress()
	assert.Equal(t, w1, mem.Fetch(w0, logic.O, addr0), "invalid ram4k[0]")
	assert.Equal(t, w2, mem.Fetch(w0, logic.O, addr1), "invalid ram4k[1]")
	assert.Equal(t, w0, mem.Fetch(w0, logic.O, addr2), "invalid ram4k[2]")
	assert.Equal(t, w0, mem.Fetch(w0, logic.O, addr3), "invalid ram4k[3]")
	assert.Equal(t, w0, mem.Fetch(w0, logic.O, addr4), "invalid screen.ram4k[0]")
	assert.Equal(t, w0, mem.Fetch(w0, logic.O, addr5), "invalid screen.ram4k[1]")
	assert.Equal(t, wkb, mem.Fetch(w0, logic.O, addr6), "inbalid keyboard")

	mem.Fetch(w3, logic.I, addr2)
	clock.Progress()
	assert.Equal(t, w1, mem.Fetch(w0, logic.O, addr0), "invalid ram4k[0]")
	assert.Equal(t, w2, mem.Fetch(w0, logic.O, addr1), "invalid ram4k[1]")
	assert.Equal(t, w3, mem.Fetch(w0, logic.O, addr2), "invalid ram4k[2]")
	assert.Equal(t, w0, mem.Fetch(w0, logic.O, addr3), "invalid ram4k[3]")
	assert.Equal(t, w0, mem.Fetch(w0, logic.O, addr4), "invalid screen.ram4k[0]")
	assert.Equal(t, w0, mem.Fetch(w0, logic.O, addr5), "invalid screen.ram4k[1]")
	assert.Equal(t, wkb, mem.Fetch(w0, logic.O, addr6), "inbalid keyboard")

	mem.Fetch(w4, logic.I, addr3)
	clock.Progress()
	assert.Equal(t, w1, mem.Fetch(w0, logic.O, addr0), "invalid ram4k[0]")
	assert.Equal(t, w2, mem.Fetch(w0, logic.O, addr1), "invalid ram4k[1]")
	assert.Equal(t, w3, mem.Fetch(w0, logic.O, addr2), "invalid ram4k[2]")
	assert.Equal(t, w4, mem.Fetch(w0, logic.O, addr3), "invalid ram4k[3]")
	assert.Equal(t, w0, mem.Fetch(w0, logic.O, addr4), "invalid screen.ram4k[0]")
	assert.Equal(t, w0, mem.Fetch(w0, logic.O, addr5), "invalid screen.ram4k[1]")
	assert.Equal(t, wkb, mem.Fetch(w0, logic.O, addr6), "inbalid keyboard")

	mem.Fetch(w5, logic.I, addr4)
	clock.Progress()
	assert.Equal(t, w1, mem.Fetch(w0, logic.O, addr0), "invalid ram4k[0]")
	assert.Equal(t, w2, mem.Fetch(w0, logic.O, addr1), "invalid ram4k[1]")
	assert.Equal(t, w3, mem.Fetch(w0, logic.O, addr2), "invalid ram4k[2]")
	assert.Equal(t, w4, mem.Fetch(w0, logic.O, addr3), "invalid ram4k[3]")
	assert.Equal(t, w5, mem.Fetch(w0, logic.O, addr4), "invalid screen.ram4k[0]")
	assert.Equal(t, w0, mem.Fetch(w0, logic.O, addr5), "invalid screen.ram4k[1]")
	assert.Equal(t, wkb, mem.Fetch(w0, logic.O, addr6), "inbalid keyboard")

	mem.Fetch(w6, logic.I, addr5)
	clock.Progress()
	assert.Equal(t, w1, mem.Fetch(w0, logic.O, addr0), "invalid ram4k[0]")
	assert.Equal(t, w2, mem.Fetch(w0, logic.O, addr1), "invalid ram4k[1]")
	assert.Equal(t, w3, mem.Fetch(w0, logic.O, addr2), "invalid ram4k[2]")
	assert.Equal(t, w4, mem.Fetch(w0, logic.O, addr3), "invalid ram4k[3]")
	assert.Equal(t, w5, mem.Fetch(w0, logic.O, addr4), "invalid screen.ram4k[0]")
	assert.Equal(t, w6, mem.Fetch(w0, logic.O, addr5), "invalid screen.ram4k[1]")
	assert.Equal(t, wkb, mem.Fetch(w0, logic.O, addr6), "inbalid keyboard")
}
