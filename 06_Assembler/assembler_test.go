package assembler

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
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
