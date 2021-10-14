package compiler

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSymbolTable(t *testing.T) {
	st := NewSymbolTable()
	assert.NoError(t, st.Define("a", "Point", KdStatic))
	sType, err := st.TypeOf("a")
	assert.NoError(t, err)
	assert.Equal(t, "Point", sType)
	kind, err := st.KindOf("a")
	assert.NoError(t, err)
	assert.Equal(t, KdStatic, kind)
	assert.Equal(t, 1, st.VarCount(KdStatic))
	assert.Equal(t, 1, st.Get("a").number)

	assert.Equal(t, ErrSymbolAlreadyExists, st.Define("a", "Point", KdStatic))

	assert.NoError(t, st.Define("b", "Point", KdStatic))
	sType, err = st.TypeOf("b")
	assert.NoError(t, err)
	assert.Equal(t, "Point", sType)
	kind, err = st.KindOf("b")
	assert.NoError(t, err)
	assert.Equal(t, KdStatic, kind)
	assert.Equal(t, 2, st.VarCount(KdStatic))
	assert.Equal(t, 2, st.Get("b").number)
}
