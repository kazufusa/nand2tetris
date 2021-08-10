package assembler

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSymbolTable(t *testing.T) {
	st := NewSymbolTable()

	st.AddEntry("LOOP", 1, false)
	assert.True(t, st.Contains("LOOP"))
	addr, err := st.GetAddress("LOOP")
	assert.Equal(t, 16, addr)
	assert.NoError(t, err)

	st.AddEntry("LOOP", 1, true)
	assert.True(t, st.Contains("LOOP"))
	addr, err = st.GetAddress("LOOP")
	assert.Equal(t, 1, addr)
	assert.NoError(t, err)

	st.AddEntry("LOOP2", 2, true)
	assert.True(t, st.Contains("LOOP2"))
	addr, err = st.GetAddress("LOOP2")
	assert.Equal(t, 2, addr)
	assert.NoError(t, err)

	st.AddEntry("a", 1, false)
	addr, err = st.GetAddress("a")
	assert.True(t, st.Contains("a"))
	assert.Equal(t, 17, addr)
	assert.NoError(t, err)

	st.AddEntry("a", 2, false)
	addr, err = st.GetAddress("a")
	assert.True(t, st.Contains("a"))
	assert.Equal(t, 17, addr)
	assert.NoError(t, err)

	st.AddEntry("b", 1, false)
	addr, err = st.GetAddress("b")
	assert.True(t, st.Contains("b"))
	assert.Equal(t, 18, addr)
	assert.NoError(t, err)

	st.AddEntry("b", 2, false)
	addr, err = st.GetAddress("b")
	assert.True(t, st.Contains("b"))
	assert.Equal(t, 18, addr)
	assert.NoError(t, err)

	st.AddEntry("R0", 1, false)
	addr, err = st.GetAddress("R0")
	assert.True(t, st.Contains("R0"))
	assert.Equal(t, 0, addr)
	assert.NoError(t, err)
}
