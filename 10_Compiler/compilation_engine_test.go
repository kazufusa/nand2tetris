package compiler

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestCompilationEngine(t *testing.T) {
	var tests = []struct {
		name string
	}{
		// {"./test/Square/Main.jack"},
		// {"./test/Square/Square.jack"},
		// {"./test/Square/SquareGame.jack"},
		// {"./test/ArrayTest/Main.jack"},
		{"./test/ExpressionLessSquare/Main.jack"},
		// {"./test/ExpressionLessSquare/Square.jack"},
		// {"./test/ExpressionLessSquare/SquareGame.jack"},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			tk, err := NewTokenizer(tt.name)
			require.NoError(t, err)

			ce := NewCompilationEngine(tk.tokens)

			tree, err := ce.compileClass()
			require.NoError(t, err)
			fmt.Println(tree.ToString(""))
		})
	}
}

func TestCompileTerm(t *testing.T) {
	var tests = []struct {
		expected string
		err      bool
		given    []Token
	}{
		{"", true, []Token{}},
		{"", true, []Token{{TkSymbol, "("}}},
		{"", true, []Token{{TkKeyWord, "if"}}},
		{"", true, []Token{{TkIdentifier, "b"}, {TkSymbol, "["}, {TkSymbol, ")"}, {TkSymbol, "]"}}},
		{"", true, []Token{{TkIdentifier, "b"}, {TkSymbol, "["}, {TkIdentifier, "a"}}},
		{"<term>\n  <identifier> b </identifier>\n</term>\n", false, []Token{{TkIdentifier, "b"}}},
		{`<term>
  <symbol> ( </symbol>
  <expression>
    <term>
      <identifier> a </identifier>
    </term>
  </expression>
  <symbol> ) </symbol>
</term>
`, false, []Token{{TkSymbol, "("}, {TkIdentifier, "a"}, {TkSymbol, ")"}}},
		{`<term>
  <identifier> b </identifier>
  <symbol> [ </symbol>
  <expression>
    <term>
      <identifier> c </identifier>
    </term>
  </expression>
  <symbol> ] </symbol>
</term>
`,
			false,
			[]Token{{TkIdentifier, "b"}, {TkSymbol, "["}, {TkIdentifier, "c"}, {TkSymbol, "]"}},
		},
		{"<term>\n  <symbol> - </symbol>\n  <term>\n    <identifier> a </identifier>\n  </term>\n</term>\n",
			false, []Token{{TkSymbol, "-"}, {TkIdentifier, "a"}}},
		{"<term>\n  <integerConstant> 2 </integerConstant>\n</term>\n",
			false, []Token{{TkIntConst, "2"}}},
		{"<term>\n  <stringConstant> hello </stringConstant>\n</term>\n",
			false, []Token{{TkStringConst, "hello"}}},
		{"<term>\n  <keyword> this </keyword>\n</term>\n",
			false, []Token{{TkKeyWord, "this"}}},
		{"<term>\n  <keyword> true </keyword>\n</term>\n",
			false, []Token{{TkKeyWord, "true"}}},
		{"<term>\n  <keyword> false </keyword>\n</term>\n",
			false, []Token{{TkKeyWord, "false"}}},
		{"<term>\n  <keyword> null </keyword>\n</term>\n",
			false, []Token{{TkKeyWord, "null"}}},
		{"<term>\n  <keyword> this </keyword>\n</term>\n",
			false, []Token{{TkKeyWord, "this"}}},
		{`<term>
  <identifier> a </identifier>
  <symbol> ( </symbol>
	<expressionList>
	</expressionList>
  <symbol> ) </symbol>
</term>
`,
			false, []Token{{TkIdentifier, "a"}, {TkSymbol, "("}, {TkSymbol, ")"}}},
	}
	for i, tt := range tests {
		if i != 16 {
			continue
		}
		tt := tt
		t.Run(fmt.Sprintf("%d", i), func(t *testing.T) {
			ce := NewCompilationEngine(tt.given)
			tree, err := ce.compileTerm()
			if tt.err {
				require.Error(t, err)
				assert.Equal(t, 0, ce.iToken)
				if len(tt.given) > 0 {
					assert.Equal(t, &tt.given[0], ce.nextToken())
				}
			} else {
				require.NoError(t, err)
				assert.Equal(t, tt.expected, tree.ToString(""))
				assert.Nil(t, ce.nextToken())
			}
		})
	}
}

func TestCompileExpression(t *testing.T) {
	var tests = []struct {
		expected string
		err      bool
		given    []Token
	}{
		{"", true, []Token{}},
		{"", true, []Token{{TkSymbol, "("}}},
		{"", true, []Token{{TkKeyWord, "if"}}},
		{"", true, []Token{{TkIdentifier, "b"}, {TkSymbol, "["}, {TkSymbol, ")"}, {TkSymbol, "]"}}},
		{"", true, []Token{{TkIdentifier, "b"}, {TkSymbol, "["}, {TkIdentifier, "a"}}},
		{
			"<expression>\n" +
				"  <term>\n" +
				"    <integerConstant> 2 </integerConstant>\n" +
				"  </term>\n" +
				"</expression>\n",
			false, []Token{{TkIntConst, "2"}}},
		{
			"<expression>\n" +
				"  <term>\n" +
				"    <integerConstant> 2 </integerConstant>\n" +
				"  </term>\n" +
				"  <symbol> + </symbol>\n" +
				"  <term>\n" +
				"    <integerConstant> 2 </integerConstant>\n" +
				"  </term>\n" +
				"</expression>\n",
			false, []Token{{TkIntConst, "2"}, {TkSymbol, "+"}, {TkIntConst, "2"}}},
		{
			"<expression>\n" +
				"  <term>\n" +
				"    <stringConstant> hello </stringConstant>\n" +
				"  </term>\n" +
				"</expression>\n",
			false, []Token{{TkStringConst, "hello"}}},
	}
	for i, tt := range tests {
		tt := tt
		t.Run(fmt.Sprintf("%d", i), func(t *testing.T) {
			ce := NewCompilationEngine(tt.given)
			tree, err := ce.compileExpression()
			if tt.err {
				require.Error(t, err)
				assert.Equal(t, 0, ce.iToken)
				if len(tt.given) > 0 {
					assert.Equal(t, &tt.given[0], ce.nextToken())
				}
			} else {
				require.NoError(t, err)
				assert.Equal(t, tt.expected, tree.ToString(""))
				assert.Nil(t, ce.nextToken())
			}
		})
	}

}