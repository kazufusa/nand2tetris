package compiler

type TokenTag string

type KeyWordTag string

type StructureTag string

const (
	TkKeyWord     TokenTag = "keyword"
	TkSymbol      TokenTag = "symbol"
	TkIdentifier  TokenTag = "identifier"
	TkIntConst    TokenTag = "integerConstant"
	TkStringConst TokenTag = "stringConstant"

	KwClass       KeyWordTag = "class"
	KwMethod      KeyWordTag = "method"
	KwFunction    KeyWordTag = "function"
	KwConstructor KeyWordTag = "constructor"
	KwInt         KeyWordTag = "int"
	KwBoolean     KeyWordTag = "boolean"
	KwChar        KeyWordTag = "char"
	KwVoid        KeyWordTag = "void"
	KwVar         KeyWordTag = "var"
	KwStatic      KeyWordTag = "static"
	KwField       KeyWordTag = "field"
	KwLet         KeyWordTag = "let"
	KwDo          KeyWordTag = "do"
	KwIf          KeyWordTag = "if"
	KwElse        KeyWordTag = "else"
	KwWhile       KeyWordTag = "while"
	KwReturn      KeyWordTag = "return"
	KwTrue        KeyWordTag = "true"
	KwFalse       KeyWordTag = "false"
	KwNull        KeyWordTag = "null"
	KwThis        KeyWordTag = "this"

	StrClass           StructureTag = "class"
	StrClassVarDec     StructureTag = "classVarDec"
	StrSubroutineDec   StructureTag = "subroutineDec"
	StrParameterList   StructureTag = "parameterList"
	StrSubroutineBody  StructureTag = "subroutineBody"
	StrVarDec          StructureTag = "varDec"
	StrStatements      StructureTag = "statements"
	StrLetStatement    StructureTag = "letStatement"
	StrIfStatement     StructureTag = "ifStatement"
	StrWhileStatement  StructureTag = "whileStatement"
	StrDoStatement     StructureTag = "doStatement"
	StrReturnStatement StructureTag = "returnStatement"
	StrExpression      StructureTag = "expression"
	StrTerm            StructureTag = "term"
	StrExpressionList  StructureTag = "expressionList"
)
