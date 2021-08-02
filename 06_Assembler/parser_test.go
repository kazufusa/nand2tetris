package assembler

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewParser(t *testing.T) {
	parser := NewParser(`
  // test comment
  @123 // comment
  D =  A // comment
  // test comment
  @  17 // comment
  M = D // comment
 // comment
`)
	expected := []string{"@123", "D=A", "@17", "M=D"}
	assert.Equal(t, expected, parser.lines)
}

func TestParser(t *testing.T) {
	parser := Parser{
		lines: []string{
			"@a",
			"@123",
			"(LOOP)",
			"(456)",
			"M",
			"D=M",
			"M;JMP",
			"D=M;JMP",
			"MD=M-1",
		},
	}

	assert.Equal(t, A_INSTRUCTION, parser.InstructionType())
	assert.Equal(t, "a", parser.Symbol())

	assert.True(t, parser.HasMoreLines())
	parser.Advance()

	assert.Equal(t, A_INSTRUCTION, parser.InstructionType())
	assert.Equal(t, "123", parser.Symbol())

	assert.True(t, parser.HasMoreLines())
	parser.Advance()

	assert.Equal(t, L_INSTRUCTION, parser.InstructionType())
	assert.Equal(t, "LOOP", parser.Symbol())

	assert.True(t, parser.HasMoreLines())
	parser.Advance()

	assert.Equal(t, L_INSTRUCTION, parser.InstructionType())
	assert.Equal(t, "456", parser.Symbol())

	assert.True(t, parser.HasMoreLines())
	parser.Advance()

	assert.Equal(t, C_INSTRUCTION, parser.InstructionType())
	assert.Equal(t, "M", parser.Comp())
	assert.Equal(t, "", parser.Dest())
	assert.Equal(t, "", parser.Jump())

	assert.True(t, parser.HasMoreLines())
	parser.Advance()

	assert.Equal(t, C_INSTRUCTION, parser.InstructionType())
	assert.Equal(t, "M", parser.Comp())
	assert.Equal(t, "D", parser.Dest())
	assert.Equal(t, "", parser.Jump())

	assert.True(t, parser.HasMoreLines())
	parser.Advance()

	assert.Equal(t, C_INSTRUCTION, parser.InstructionType())
	assert.Equal(t, "M", parser.Comp())
	assert.Equal(t, "", parser.Dest())
	assert.Equal(t, "JMP", parser.Jump())

	assert.True(t, parser.HasMoreLines())
	parser.Advance()

	assert.Equal(t, C_INSTRUCTION, parser.InstructionType())
	assert.Equal(t, "M", parser.Comp())
	assert.Equal(t, "D", parser.Dest())
	assert.Equal(t, "JMP", parser.Jump())

	assert.True(t, parser.HasMoreLines())
	parser.Advance()

	assert.Equal(t, C_INSTRUCTION, parser.InstructionType())
	assert.Equal(t, "M-1", parser.Comp())
	assert.Equal(t, "MD", parser.Dest())
	assert.Equal(t, "", parser.Jump())

	assert.False(t, parser.HasMoreLines())
}
