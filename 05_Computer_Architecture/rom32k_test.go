package computer

import (
	"io/ioutil"
	"os"
	"testing"

	logic "github.com/kazufusa/nand2tetris/01_Boolean_Logic"
	"github.com/stretchr/testify/assert"
)

func TestROM32K(t *testing.T) {
	instructions := []Word{
		{1, 1, 1},
		{0, 0, 0, 1},
		{1, 0, 0, 1},
		{0, 1, 0, 1},
		{1, 1, 0, 1},
		{0, 0, 1, 1},
		{1, 0, 1, 1},
		{0, 1, 1, 1},
		{1, 1, 1, 1},
	}
	rom := NewROM32K()
	rom.BulkLoad(instructions)

	addr := [15]logic.Bit{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1}
	expected := Word{1, 1, 1, 1, 1, 1, 1, 1, 1, 1}
	rom.load(addr, expected)

	assert.Equal(t, instructions[0], rom.Fetch([15]logic.Bit{0}))
	assert.Equal(t, instructions[1], rom.Fetch([15]logic.Bit{1}))
	assert.Equal(t, instructions[2], rom.Fetch([15]logic.Bit{0, 1}))
	assert.Equal(t, instructions[3], rom.Fetch([15]logic.Bit{1, 1}))
	assert.Equal(t, instructions[4], rom.Fetch([15]logic.Bit{0, 0, 1}))
	assert.Equal(t, instructions[5], rom.Fetch([15]logic.Bit{1, 0, 1}))
	assert.Equal(t, instructions[6], rom.Fetch([15]logic.Bit{0, 1, 1}))
	assert.Equal(t, instructions[7], rom.Fetch([15]logic.Bit{1, 1, 1}))

	assert.Equal(t, expected, rom.Fetch(addr))
}

func TestLoadHackFile(t *testing.T) {
	ioutil.WriteFile("./test.hack", []byte(
		`1111111111111111
0000000000000000
0101010101010101
1010101010101010
`,
	), os.ModePerm)
	defer os.Remove("./test.hack")

	rom := NewROM32K()
	rom.LoadHackFile("./test.hack")
	assert.Equal(
		t,
		Word{1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1},
		rom.Fetch([15]logic.Bit{0}),
	)
	assert.Equal(
		t,
		Word{},
		rom.Fetch([15]logic.Bit{1}),
	)
	assert.Equal(
		t,
		Word{1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0},
		rom.Fetch([15]logic.Bit{0, 1}),
	)
	assert.Equal(
		t,
		Word{0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1},
		rom.Fetch([15]logic.Bit{1, 1}),
	)
}
