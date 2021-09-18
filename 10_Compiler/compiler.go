package compiler

type TokenType string

type KeyWordType string

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
)
