package compiler

type TokenType string

type KeyWordType string

type Structure int

type Statement int

type Expression int

const (
	TkKeyWord     TokenType = "keyword"
	TkSymbol      TokenType = "symbol"
	TkIdentifier  TokenType = "identifier"
	TkIntConst    TokenType = "integerConstant"
	TkStringConst TokenType = "stringConstant"

	KwClass       KeyWordType = "class"
	KwMethod      KeyWordType = "method"
	KwFunction    KeyWordType = "function"
	KwConstructor KeyWordType = "constructor"
	KwInt         KeyWordType = "int"
	KwBoolean     KeyWordType = "boolean"
	KwChar        KeyWordType = "char"
	KwVoid        KeyWordType = "void"
	KwVar         KeyWordType = "var"
	KwStatic      KeyWordType = "static"
	KwField       KeyWordType = "field"
	KwLet         KeyWordType = "let"
	KwDo          KeyWordType = "do"
	KwIf          KeyWordType = "if"
	KwElse        KeyWordType = "else"
	KwWhile       KeyWordType = "while"
	KwReturn      KeyWordType = "return"
	KwTrue        KeyWordType = "true"
	KwFalse       KeyWordType = "false"
	KwNull        KeyWordType = "null"
	KwThis        KeyWordType = "this"

	StrClass Structure = iota
	StrClassVarDec
	StrType
	StrSubroutineDec
	StrParameterList
	StrSubroutineBody
	StrVarDec
	StrClassName
	StrSubroutineName
	StrVarName

	StStatements Statement = iota
	StStatement
	StLetStatement
	StIfStatement
	StWhileStatement
	StDoStatement
	StReturnStatement

	ExpExpression Expression = iota
	ExpTerm
	ExpSubroutineCall
	ExpExpressionList
	ExpOp
	ExpUnaryOp
	ExpKeywordConstant
)
