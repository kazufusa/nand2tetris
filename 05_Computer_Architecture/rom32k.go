package computer

import (
	"bytes"
	"io/ioutil"

	logic "github.com/kazufusa/nand2tetris/01_Boolean_Logic"
	arithmetic "github.com/kazufusa/nand2tetris/02_Boolean_Arithmetic"
	memory "github.com/kazufusa/nand2tetris/03_Memory"
)

type ROM32K struct {
	clock *memory.Clock
	rams  [2]memory.RAM16384
}

func NewROM32K() ROM32K {
	var clock memory.Clock = 0
	return ROM32K{
		clock: &clock,
		rams: [2]memory.RAM16384{
			memory.NewRAM16384(&clock),
			memory.NewRAM16384(&clock),
		},
	}
}

func (rom *ROM32K) Fetch(addr [15]logic.Bit) (out Word) {
	subAddr := [14]logic.Bit{
		addr[0], addr[1], addr[2],
		addr[3], addr[4], addr[5],
		addr[6], addr[7], addr[8],
		addr[9], addr[10], addr[11],
		addr[12], addr[13],
	}
	return logic.Mux16(
		rom.rams[0].Apply(Word{}, logic.O, subAddr),
		rom.rams[1].Apply(Word{}, logic.O, subAddr),
		addr[14],
	)
}

func (rom *ROM32K) LoadHackFile(p string) error {
	b, err := ioutil.ReadFile(p)
	if err != nil {
		return err
	}
	addr := Word{}
	for _, b := range bytes.Split(b, []byte("\n")) {
		if len(b) != 16 {
			continue
		}
		var _addr [15]logic.Bit
		copy(_addr[:], addr[0:15])
		word := Word{
			b[15] - 48,
			b[14] - 48,
			b[13] - 48,
			b[12] - 48,
			b[11] - 48,
			b[10] - 48,
			b[9] - 48,
			b[8] - 48,
			b[7] - 48,
			b[6] - 48,
			b[5] - 48,
			b[4] - 48,
			b[3] - 48,
			b[2] - 48,
			b[1] - 48,
			b[0] - 48,
		}
		rom.load(_addr, word)
		addr = arithmetic.Inc16(addr)
	}

	return nil
}

func (rom *ROM32K) BulkLoad(ws []Word) {
	addr := Word{}
	for _, w := range ws {
		var _addr [15]logic.Bit
		copy(_addr[:], addr[0:15])
		rom.load(_addr, w)
		addr = arithmetic.Inc16(addr)
	}
}

func (rom *ROM32K) load(addr [15]logic.Bit, w Word) {
	subAddr := [14]logic.Bit{
		addr[0], addr[1], addr[2],
		addr[3], addr[4], addr[5],
		addr[6], addr[7], addr[8],
		addr[9], addr[10], addr[11],
		addr[12], addr[13],
	}
	rom.rams[0].Apply(w, logic.Not(addr[14]), subAddr)
	rom.rams[1].Apply(w, addr[14], subAddr)
	rom.clock.Progress()
}
