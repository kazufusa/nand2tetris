package compiler

import (
	"os"
	"strings"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/stretchr/testify/require"
)

func TestTokenizer(t *testing.T) {
	var tests = []struct {
		name string
	}{
		{"./test/Square/Main.jack"},
		{"./test/Square/Square.jack"},
		{"./test/Square/SquareGame.jack"},
		{"./test/ArrayTest/Main.jack"},
		{"./test/ExpressionLessSquare/Main.jack"},
		{"./test/ExpressionLessSquare/Square.jack"},
		{"./test/ExpressionLessSquare/SquareGame.jack"},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			expecedFile := strings.Replace(tt.name, ".jack", "T.xml", 1)
			expectedBuf, err := os.ReadFile(expecedFile)
			require.NoError(t, err)
			expected := strings.ReplaceAll(string(expectedBuf), "\r", "")

			tk, err := NewTokenizer(tt.name)
			require.NoError(t, err)
			if diff := cmp.Diff(expected, tk.ToXml()); diff != "" {
				t.Error(diff)
			}
		})
	}
}
