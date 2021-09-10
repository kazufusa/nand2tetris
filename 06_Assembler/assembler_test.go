package assembler

import (
	"os"
	"testing"

	logic "github.com/kazufusa/nand2tetris/01_Boolean_Logic"
	computer "github.com/kazufusa/nand2tetris/05_Computer_Architecture"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestAssembler(t *testing.T) {
	expected, err := os.ReadFile("./Rect.hack")
	assert.NoError(t, err)
	asm, err := NewAssembler("./Rect.asm")
	assert.NoError(t, err)
	ret, err := asm.Assemble()
	assert.NoError(t, err)
	assert.Equal(t, string(expected), ret)
}

func TestAssemblerWithEmulator(t *testing.T) {
	asm, err := NewAssembler("./Max.asm")
	require.NoError(t, err)
	ret, err := asm.Assemble()
	require.NoError(t, err)
	com := computer.NewEmulator(ret)
	com.WriteRom(0, 14)
	com.WriteRom(1, 12)
	for i := 0; i < 20; i++ {
		com.FetchAndExecute(logic.O)
	}
	assert.Equal(t, 14, com.ROMMap()[2])
}
