package compiler

import (
	"fmt"
	"os"
	"strings"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestCompilationEngine(t *testing.T) {
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
			tk, err := NewTokenizer(tt.name)
			require.NoError(t, err)

			ce := NewCompilationEngine(tk.tokens)

			tree, err := ce.compileClass()
			require.NoError(t, err)
			actual := tree.ToString("")

			expectedXml := strings.TrimSuffix(tt.name, ".jack") + ".xml"
			expectedBytes, err := os.ReadFile(expectedXml)
			expected := strings.ReplaceAll(string(expectedBytes), "\r", "")
			require.NoError(t, err)
			if diff := cmp.Diff(expected, actual); diff != "" {
				t.Error(diff)
			}
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
		{
			"<term>\n" +
				"  <identifier> b </identifier>\n" +
				"</term>\n",
			false, []Token{{TkIdentifier, "b"}}},
		{
			"<term>\n" +
				"  <symbol> ( </symbol>\n" +
				"  <expression>\n" +
				"    <term>\n" +
				"      <identifier> a </identifier>\n" +
				"    </term>\n" +
				"  </expression>\n" +
				"  <symbol> ) </symbol>\n" +
				"</term>\n",
			false, []Token{{TkSymbol, "("}, {TkIdentifier, "a"}, {TkSymbol, ")"}}},
		{
			"<term>\n" +
				"  <identifier> b </identifier>\n" +
				"  <symbol> [ </symbol>\n" +
				"  <expression>\n" +
				"    <term>\n" +
				"      <identifier> c </identifier>\n" +
				"    </term>\n" +
				"  </expression>\n" +
				"  <symbol> ] </symbol>\n" +
				"</term>\n",
			false,
			[]Token{{TkIdentifier, "b"}, {TkSymbol, "["}, {TkIdentifier, "c"}, {TkSymbol, "]"}},
		},
		{"<term>\n" +
			"  <symbol> - </symbol>\n" +
			"  <term>\n" +
			"    <identifier> a </identifier>\n" +
			"  </term>\n" +
			"</term>\n",
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
		{"<term>\n" +
			"  <identifier> a </identifier>\n" +
			"  <symbol> ( </symbol>\n" +
			"  <expressionList>\n" +
			"  </expressionList>\n" +
			"  <symbol> ) </symbol>\n" +
			"</term>\n",
			false, []Token{{TkIdentifier, "a"}, {TkSymbol, "("}, {TkSymbol, ")"}}},
	}
	for i, tt := range tests {
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

func TestCompileExpressionList(t *testing.T) {
	empty := "<expressionList>\n</expressionList>\n"
	var tests = []struct {
		expected        string
		emptyIsExpected bool
		given           []Token
	}{
		{empty, true, []Token{{TkSymbol, "("}}},
		{empty, true, []Token{{TkIdentifier, "b"}, {TkSymbol, "["}, {TkSymbol, ")"}, {TkSymbol, "]"}}},
		{empty, true, []Token{{TkKeyWord, "if"}}},
		{empty, true, []Token{{TkIdentifier, "b"}, {TkSymbol, "["}, {TkIdentifier, "a"}}},
		{empty, true, []Token{}},
		{
			"<expressionList>\n" +
				"  <expression>\n" +
				"    <term>\n" +
				"      <integerConstant> 2 </integerConstant>\n" +
				"    </term>\n" +
				"  </expression>\n" +
				"</expressionList>\n",
			false, []Token{{TkIntConst, "2"}},
		},
		{
			"<expressionList>\n" +
				"  <expression>\n" +
				"    <term>\n" +
				"      <integerConstant> 2 </integerConstant>\n" +
				"    </term>\n" +
				"  </expression>\n" +
				"  <symbol> , </symbol>\n" +
				"  <expression>\n" +
				"    <term>\n" +
				"      <integerConstant> 3 </integerConstant>\n" +
				"    </term>\n" +
				"  </expression>\n" +
				"</expressionList>\n",
			false, []Token{{TkIntConst, "2"}, {TkSymbol, ","}, {TkIntConst, "3"}},
		},
		{"<expressionList>\n" +
			"  <expression>\n" +
			"    <term>\n" +
			"      <integerConstant> 2 </integerConstant>\n" +
			"    </term>\n" +
			"    <symbol> + </symbol>\n" +
			"    <term>\n" +
			"      <integerConstant> 2 </integerConstant>\n" +
			"    </term>\n" +
			"  </expression>\n" +
			"</expressionList>\n",
			false, []Token{{TkIntConst, "2"}, {TkSymbol, "+"}, {TkIntConst, "2"}}},
		{
			"<expressionList>\n" +
				"  <expression>\n" +
				"    <term>\n" +
				"      <stringConstant> hello </stringConstant>\n" +
				"    </term>\n" +
				"  </expression>\n" +
				"</expressionList>\n",
			false, []Token{{TkStringConst, "hello"}},
		},
	}
	for i, tt := range tests {
		tt := tt
		t.Run(fmt.Sprintf("%d", i), func(t *testing.T) {
			ce := NewCompilationEngine(tt.given)
			tree, err := ce.compileExpressionList()
			if tt.emptyIsExpected {
				require.NoError(t, err)
				assert.Equal(t, tt.expected, tree.ToString(""))
				assert.Equal(t, 0, ce.iToken)
			} else {
				require.NoError(t, err)
				assert.Equal(t, tt.expected, tree.ToString(""))
				assert.Nil(t, ce.nextToken())
			}
		})
	}
}

func TestCompileVarDec(t *testing.T) {
	var tests = []struct {
		expected string
		err      bool
		given    []Token
	}{
		{"", true, []Token{{TkSymbol, "("}}},
		{
			"", true, []Token{
				{TkKeyWord, "var"},
				{TkIdentifier, "testType"},
				{TkIdentifier, "testName"},
			},
		},
		{
			"<varDec>\n" +
				"  <keyword> var </keyword>\n" +
				"  <identifier> testType </identifier>\n" +
				"  <identifier> testName </identifier>\n" +
				"  <symbol> ; </symbol>\n" +
				"</varDec>\n",
			false, []Token{
				{TkKeyWord, "var"},
				{TkIdentifier, "testType"},
				{TkIdentifier, "testName"},
				{TkSymbol, ";"},
			},
		},
		{
			"<varDec>\n" +
				"  <keyword> var </keyword>\n" +
				"  <keyword> int </keyword>\n" +
				"  <identifier> testName1 </identifier>\n" +
				"  <symbol> , </symbol>\n" +
				"  <identifier> testName2 </identifier>\n" +
				"  <symbol> , </symbol>\n" +
				"  <identifier> testName3 </identifier>\n" +
				"  <symbol> ; </symbol>\n" +
				"</varDec>\n",
			false, []Token{
				{TkKeyWord, "var"},
				{TkKeyWord, "int"},
				{TkIdentifier, "testName1"},
				{TkSymbol, ","},
				{TkIdentifier, "testName2"},
				{TkSymbol, ","},
				{TkIdentifier, "testName3"},
				{TkSymbol, ";"},
			},
		},
	}
	for i, tt := range tests {
		tt := tt
		t.Run(fmt.Sprintf("%d", i), func(t *testing.T) {
			ce := NewCompilationEngine(tt.given)
			tree, err := ce.compileVarDec()
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

func TestCompileLet(t *testing.T) {
	var tests = []struct {
		expected string
		err      bool
		given    []Token
	}{
		{"", true, []Token{{TkSymbol, "("}}},
		{"", true, []Token{{TkKeyWord, "let"}}},
		{"", true, []Token{{TkKeyWord, "let"}, {TkIdentifier, "a"}}},
		{
			"<letStatement>\n" +
				"  <keyword> let </keyword>\n" +
				"  <identifier> a </identifier>\n" +
				"  <symbol> = </symbol>\n" +
				"  <expression>\n" +
				"    <term>\n" +
				"      <identifier> b </identifier>\n" +
				"    </term>\n" +
				"  </expression>\n" +
				"  <symbol> ; </symbol>\n" +
				"</letStatement>\n",
			false, []Token{
				{TkKeyWord, "let"},
				{TkIdentifier, "a"},
				{TkSymbol, "="},
				{TkIdentifier, "b"},
				{TkSymbol, ";"},
			},
		},
		{
			"<letStatement>\n" +
				"  <keyword> let </keyword>\n" +
				"  <identifier> a </identifier>\n" +
				"  <symbol> [ </symbol>\n" +
				"  <expression>\n" +
				"    <term>\n" +
				"      <identifier> aa </identifier>\n" +
				"    </term>\n" +
				"  </expression>\n" +
				"  <symbol> ] </symbol>\n" +
				"  <symbol> = </symbol>\n" +
				"  <expression>\n" +
				"    <term>\n" +
				"      <identifier> b </identifier>\n" +
				"    </term>\n" +
				"  </expression>\n" +
				"  <symbol> ; </symbol>\n" +
				"</letStatement>\n",
			false, []Token{
				{TkKeyWord, "let"},
				{TkIdentifier, "a"},
				{TkSymbol, "["},
				{TkIdentifier, "aa"},
				{TkSymbol, "]"},
				{TkSymbol, "="},
				{TkIdentifier, "b"},
				{TkSymbol, ";"},
			},
		},
	}
	for i, tt := range tests {
		tt := tt
		t.Run(fmt.Sprintf("%d", i), func(t *testing.T) {
			ce := NewCompilationEngine(tt.given)
			tree, err := ce.compileLet()
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

func TestCompileIf(t *testing.T) {
	var tests = []struct {
		expected string
		err      bool
		given    []Token
	}{
		{"", true, []Token{
			{TkSymbol, "("},
			{TkSymbol, ")"},
		}},
		{"", true, []Token{
			{TkKeyWord, "if"},
			{TkSymbol, "{"},
			{TkSymbol, "}"},
		}},
		{"", true, []Token{
			{TkKeyWord, "if"},
			{TkSymbol, "("},
			{TkSymbol, ")"},
			{TkSymbol, "{"},
			{TkSymbol, "}"},
		}},
		{
			"<ifStatement>\n" +
				"  <keyword> if </keyword>\n" +
				"  <symbol> ( </symbol>\n" +
				"  <expression>\n" +
				"    <term>\n" +
				"      <keyword> true </keyword>\n" +
				"    </term>\n" +
				"  </expression>\n" +
				"  <symbol> ) </symbol>\n" +
				"  <symbol> { </symbol>\n" +
				"  <statements>\n" +
				"    <letStatement>\n" +
				"      <keyword> let </keyword>\n" +
				"      <identifier> a </identifier>\n" +
				"      <symbol> = </symbol>\n" +
				"      <expression>\n" +
				"        <term>\n" +
				"          <identifier> b </identifier>\n" +
				"        </term>\n" +
				"      </expression>\n" +
				"      <symbol> ; </symbol>\n" +
				"    </letStatement>\n" +
				"  </statements>\n" +
				"  <symbol> } </symbol>\n" +
				"</ifStatement>\n",
			false, []Token{
				{TkKeyWord, "if"},
				{TkSymbol, "("},
				{TkKeyWord, "true"},
				{TkSymbol, ")"},
				{TkSymbol, "{"},
				{TkKeyWord, "let"},
				{TkIdentifier, "a"},
				{TkSymbol, "="},
				{TkIdentifier, "b"},
				{TkSymbol, ";"},
				{TkSymbol, "}"},
			}},
		{
			"<ifStatement>\n" +
				"  <keyword> if </keyword>\n" +
				"  <symbol> ( </symbol>\n" +
				"  <expression>\n" +
				"    <term>\n" +
				"      <keyword> true </keyword>\n" +
				"    </term>\n" +
				"  </expression>\n" +
				"  <symbol> ) </symbol>\n" +
				"  <symbol> { </symbol>\n" +
				"  <statements>\n" +
				"    <letStatement>\n" +
				"      <keyword> let </keyword>\n" +
				"      <identifier> a </identifier>\n" +
				"      <symbol> = </symbol>\n" +
				"      <expression>\n" +
				"        <term>\n" +
				"          <identifier> b </identifier>\n" +
				"        </term>\n" +
				"      </expression>\n" +
				"      <symbol> ; </symbol>\n" +
				"    </letStatement>\n" +
				"  </statements>\n" +
				"  <symbol> } </symbol>\n" +
				"  <keyword> else </keyword>\n" +
				"  <symbol> { </symbol>\n" +
				"  <statements>\n" +
				"    <letStatement>\n" +
				"      <keyword> let </keyword>\n" +
				"      <identifier> b </identifier>\n" +
				"      <symbol> = </symbol>\n" +
				"      <expression>\n" +
				"        <term>\n" +
				"          <identifier> a </identifier>\n" +
				"        </term>\n" +
				"      </expression>\n" +
				"      <symbol> ; </symbol>\n" +
				"    </letStatement>\n" +
				"  </statements>\n" +
				"  <symbol> } </symbol>\n" +
				"</ifStatement>\n",
			false, []Token{
				{TkKeyWord, "if"},
				{TkSymbol, "("},
				{TkKeyWord, "true"},
				{TkSymbol, ")"},
				{TkSymbol, "{"},
				{TkKeyWord, "let"},
				{TkIdentifier, "a"},
				{TkSymbol, "="},
				{TkIdentifier, "b"},
				{TkSymbol, ";"},
				{TkSymbol, "}"},
				{TkKeyWord, "else"},
				{TkSymbol, "{"},
				{TkKeyWord, "let"},
				{TkIdentifier, "b"},
				{TkSymbol, "="},
				{TkIdentifier, "a"},
				{TkSymbol, ";"},
				{TkSymbol, "}"},
			}},
	}
	for i, tt := range tests {
		tt := tt
		t.Run(fmt.Sprintf("%d", i), func(t *testing.T) {
			ce := NewCompilationEngine(tt.given)
			tree, err := ce.compileIf()
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

func TestCompileReturn(t *testing.T) {
	var tests = []struct {
		expected string
		err      bool
		given    []Token
	}{
		{"", true, []Token{{TkSymbol, "return"}}},
		{"", true, []Token{
			{TkKeyWord, "if"},
			{TkSymbol, "("},
			{TkSymbol, ")"},
			{TkSymbol, "{"},
			{TkSymbol, "}"},
		}},
		{
			"<returnStatement>\n" +
				"  <keyword> return </keyword>\n" +
				"  <symbol> ; </symbol>\n" +
				"</returnStatement>\n",
			false, []Token{
				{TkKeyWord, "return"},
				{TkSymbol, ";"},
			}},
		{
			"<returnStatement>\n" +
				"  <keyword> return </keyword>\n" +
				"  <expression>\n" +
				"    <term>\n" +
				"      <symbol> ( </symbol>\n" +
				"      <expression>\n" +
				"        <term>\n" +
				"          <identifier> a </identifier>\n" +
				"        </term>\n" +
				"        <symbol> + </symbol>\n" +
				"        <term>\n" +
				"          <identifier> b </identifier>\n" +
				"          <symbol> ( </symbol>\n" +
				"          <expressionList>\n" +
				"          </expressionList>\n" +
				"          <symbol> ) </symbol>\n" +
				"        </term>\n" +
				"      </expression>\n" +
				"      <symbol> ) </symbol>\n" +
				"    </term>\n" +
				"  </expression>\n" +
				"  <symbol> ; </symbol>\n" +
				"</returnStatement>\n",
			false, []Token{
				{TkKeyWord, "return"},
				{TkSymbol, "("},
				{TkIdentifier, "a"},
				{TkSymbol, "+"},
				{TkIdentifier, "b"},
				{TkSymbol, "("},
				{TkSymbol, ")"},
				{TkSymbol, ")"},
				{TkSymbol, ";"},
			}},
	}
	for i, tt := range tests {
		tt := tt
		t.Run(fmt.Sprintf("%d", i), func(t *testing.T) {
			ce := NewCompilationEngine(tt.given)
			tree, err := ce.compileReturn()
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

func TestCompileWhile(t *testing.T) {
	var tests = []struct {
		expected string
		err      bool
		given    []Token
	}{
		{"", true, []Token{{TkSymbol, "while"}}},
		{"", true, []Token{
			{TkKeyWord, "return"},
			{TkSymbol, ";"},
		}},
		{"", true, []Token{
			{TkKeyWord, "while"},
			{TkSymbol, "("},
			{TkSymbol, ")"},
			{TkSymbol, "{"},
			{TkSymbol, "}"},
		}},
		{
			"<whileStatement>\n" +
				"  <keyword> while </keyword>\n" +
				"  <symbol> ( </symbol>\n" +
				"  <expression>\n" +
				"    <term>\n" +
				"      <keyword> true </keyword>\n" +
				"    </term>\n" +
				"  </expression>\n" +
				"  <symbol> ) </symbol>\n" +
				"  <symbol> { </symbol>\n" +
				"  <statements>\n" +
				"  </statements>\n" +
				"  <symbol> } </symbol>\n" +
				"</whileStatement>\n",
			false, []Token{
				{TkKeyWord, "while"},
				{TkSymbol, "("},
				{TkKeyWord, "true"},
				{TkSymbol, ")"},
				{TkSymbol, "{"},
				{TkSymbol, "}"},
			}},
	}
	for i, tt := range tests {
		tt := tt
		t.Run(fmt.Sprintf("%d", i), func(t *testing.T) {
			ce := NewCompilationEngine(tt.given)
			tree, err := ce.compileWhile()
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

func TestCompileClassVarDec(t *testing.T) {
	var tests = []struct {
		expected string
		err      bool
		given    []Token
	}{
		{"", true, []Token{{TkSymbol, "("}}},
		{
			"<classVarDec>\n" +
				"  <keyword> field </keyword>\n" +
				"  <identifier> testType </identifier>\n" +
				"  <identifier> testName </identifier>\n" +
				"  <symbol> ; </symbol>\n" +
				"</classVarDec>\n",
			false, []Token{
				{TkKeyWord, "field"},
				{TkIdentifier, "testType"},
				{TkIdentifier, "testName"},
				{TkSymbol, ";"},
			},
		},
		{
			"<classVarDec>\n" +
				"  <keyword> static </keyword>\n" +
				"  <identifier> testType </identifier>\n" +
				"  <identifier> testName </identifier>\n" +
				"  <symbol> ; </symbol>\n" +
				"</classVarDec>\n",
			false, []Token{
				{TkKeyWord, "static"},
				{TkIdentifier, "testType"},
				{TkIdentifier, "testName"},
				{TkSymbol, ";"},
			},
		},
		{
			"<classVarDec>\n" +
				"  <keyword> field </keyword>\n" +
				"  <keyword> int </keyword>\n" +
				"  <identifier> testName1 </identifier>\n" +
				"  <symbol> , </symbol>\n" +
				"  <identifier> testName2 </identifier>\n" +
				"  <symbol> , </symbol>\n" +
				"  <identifier> testName3 </identifier>\n" +
				"  <symbol> ; </symbol>\n" +
				"</classVarDec>\n",
			false, []Token{
				{TkKeyWord, "field"},
				{TkKeyWord, "int"},
				{TkIdentifier, "testName1"},
				{TkSymbol, ","},
				{TkIdentifier, "testName2"},
				{TkSymbol, ","},
				{TkIdentifier, "testName3"},
				{TkSymbol, ";"},
			},
		},
	}
	for i, tt := range tests {
		if i != 3 {
			continue
		}
		tt := tt
		t.Run(fmt.Sprintf("%d", i), func(t *testing.T) {
			ce := NewCompilationEngine(tt.given)
			tree, err := ce.compileClassVarDec()
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
