package vm

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParser(t *testing.T) {
	parser, _ := NewParser(`// Pushes and adds two constants
  // push 7 to stack
  push constant 7     // push
  // push 8 to stack
  push   constant   8 // push
  // pop 7 and 8 from stack and push 15 to stack
  add                 // arithmetic
	function Main.main
  //`)

	// line 1
	c, err := parser.commandType()
	assert.NoError(t, err)
	assert.Equal(t, C_PUSH, c)
	a1, err := parser.arg1()
	assert.NoError(t, err)
	assert.Equal(t, "constant", a1)
	a2, err := parser.arg2()
	assert.NoError(t, err)
	assert.Equal(t, 7, a2)

	assert.True(t, parser.hasMoreLines())
	parser.advance()

	// line 2
	c, err = parser.commandType()
	assert.NoError(t, err)
	assert.Equal(t, C_PUSH, c)
	a1, err = parser.arg1()
	assert.NoError(t, err)
	assert.Equal(t, "constant", a1)
	a2, err = parser.arg2()
	assert.NoError(t, err)
	assert.Equal(t, 8, a2)

	assert.True(t, parser.hasMoreLines())
	parser.advance()

	// line 3
	c, err = parser.commandType()
	assert.NoError(t, err)
	assert.Equal(t, C_ALITHMETIC, c)
	a1, err = parser.arg1()
	assert.NoError(t, err)
	assert.Equal(t, "add", a1)
	_, err = parser.arg2()
	assert.Error(t, err)

	assert.True(t, parser.hasMoreLines())
	parser.advance()

	// line 4
	c, err = parser.commandType()
	assert.NoError(t, err)
	assert.Equal(t, C_FUNCTION, c)
	a1, err = parser.arg1()
	assert.NoError(t, err)
	assert.Equal(t, "Main.main", a1)
	_, err = parser.arg2()
	assert.Error(t, err)

	assert.False(t, parser.hasMoreLines())
}
