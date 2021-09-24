package compiler

import (
	"fmt"
	"strings"
)

type ErrCompileFailed struct {
	token         *Token
	expected      string
	targetIsFound bool
}

func NewErrCompileFailed(token *Token, expected string) error {
	return ErrCompileFailed{token: token, expected: expected}
}

func targetNotFound(err error) error {
	switch _err := err.(type) {
	case ErrCompileFailed:
		return ErrCompileFailed{
			token:         _err.token,
			expected:      _err.expected,
			targetIsFound: false,
		}
	default:
		return nil
	}
}

func targetFound(err error) error {
	switch _err := err.(type) {
	case ErrCompileFailed:
		return ErrCompileFailed{
			token:         _err.token,
			expected:      _err.expected,
			targetIsFound: true,
		}
	default:
		return nil
	}
}

func isTargetFound(err error) bool {
	if _err, ok := err.(ErrCompileFailed); err != nil && ok {
		return _err.targetIsFound
	}
	return false
}

func (e ErrCompileFailed) Error() string {
	cmp := fmt.Sprintf(`expected: %s, actual: "%s"`, e.expected, e.token.value)
	if e.targetIsFound {
		return fmt.Sprintf(`target found but failed to compile. %s`, cmp)
	} else {
		return fmt.Sprintf(`failed to compile. %s`, cmp)
	}
}

type Node struct {
	children     []interface{}
	structureTag StructureTag
}

func (n *Node) ToString(indent string) string {
	ret := fmt.Sprintf("%s<%s>\n", indent, string(n.structureTag))
	cIndent := indent + "  "
	for _, c := range n.children {
		switch v := c.(type) {
		case *Node:
			ret = fmt.Sprintf("%s%s", ret, v.ToString(cIndent))
		case *Token:
			ret = fmt.Sprintf("%s%s%s\n", ret, cIndent, v.ToString())
		}
	}
	ret = fmt.Sprintf("%s%s</%s>\n", ret, indent, string(n.structureTag))
	return ret
}

type ICompilationEngine interface {
	compileClass() (*Node, error)
	compileClassVarDec() (*Node, error)
	compileSubroutine() (*Node, error)
	compileParameterList() (*Node, error)
	compileSubroutineBody() (*Node, error)
	compileVarDec() (*Node, error)
	compileStatements() (*Node, error)
	compileLet() (*Node, error)
	compileIf() (*Node, error)
	compileWhile() (*Node, error)
	compileDo() (*Node, error)
	compileReturn() (*Node, error)
	compileExpression() (*Node, error)
	compileTerm() (*Node, error)
	compileExpressionList() (*Node, error)
}

type CompilationEngine struct {
	tokens []Token
	iToken int
}

func NewCompilationEngine(tokens []Token) *CompilationEngine {
	return &CompilationEngine{tokens: tokens}
}

func (c *CompilationEngine) nextToken() *Token {
	if c.iToken < len(c.tokens) {
		defer func() { c.iToken++ }()
		return &c.tokens[c.iToken]
	}
	return nil
}

func (c *CompilationEngine) restoreNextToken(i int) {
	c.iToken = i
}

func (c *CompilationEngine) rollbackNextToken() {
	if c.iToken > 0 {
		c.iToken--
	}
}

// compileClass return compiled class structure and error
// class is consist of following:
// - identifiler of name
// - "{"
// - field variable declarations ( multiple )
// - static variable declarations ( multiple )
// - subroutine declarations
// - "}"
func (c *CompilationEngine) compileClass() (*Node, error) {
	token := c.nextToken()
	if token.tokenType != TkKeyWord ||
		KeyWordTag(token.value) != KwClass {
		return nil, NewErrCompileFailed(token, string(KwClass))
	}

	node := Node{structureTag: StrClass, children: []interface{}{token}}

	var child interface{}
	var err error

	// identifier
	child, err = c.processIdentifier()
	if err != nil {
		return nil, err
	}
	node.children = append(node.children, child)

	// {
	child, err = c.processSymbol("{")
	if err != nil {
		return nil, err
	}
	node.children = append(node.children, child)

	// static or field variable declaration
	for {
		child, err = c.compileClassVarDec()
		if isTargetFound(err) {
			return nil, err
		} else if err != nil {
			c.rollbackNextToken()
			break
		}
		node.children = append(node.children, child)
	}

	for {
		child, err = c.compileSubroutine()
		if isTargetFound(err) {
			return nil, err
		} else if err != nil {
			c.rollbackNextToken()
			break
		}
		node.children = append(node.children, child)
	}

	return &node, nil
}

func (c *CompilationEngine) compileClassVarDec() (*Node, error) {
	var child interface{}
	var err error

	node := Node{structureTag: StrClassVarDec, children: []interface{}{}}

	child, err = c.processKeyword(KwStatic, KwField)
	if err != nil {
		return nil, err
	}
	node.children = append(node.children, child)

	child, err = c.processKeyword(KwBoolean, KwChar, KwInt)
	if err != nil {
		return nil, err
	}
	node.children = append(node.children, child)

	child, err = c.processIdentifier()
	if err != nil {
		return nil, err
	}
	node.children = append(node.children, child)

	child, err = c.processSymbol(";")
	if err != nil {
		return nil, err
	}
	node.children = append(node.children, child)

	return &node, nil
}

// let, if, while, do, return
func (c *CompilationEngine) compileStatements() (*Node, error) {
	var child interface{}
	var err error
	node := Node{structureTag: StrStatements, children: []interface{}{}}

	for {
		// let statement
		child, err = c.compileLet()
		if isTargetFound(err) {
			return nil, targetFound(err)
		} else if err != nil {
			c.rollbackNextToken()
		} else {
			node.children = append(node.children, child)
			continue
		}

		// TODO if statement
		// TODO while statement

		// do statement
		child, err = c.compileDo()
		if isTargetFound(err) {
			return nil, targetFound(err)
		} else if err != nil {
			c.rollbackNextToken()
		} else {
			node.children = append(node.children, child)
			continue
		}

		// TODO return statement
		break
	}

	return &node, nil
}

// 1. let varName = expression;
// 2. let varName[expression1] = expression2;
func (c *CompilationEngine) compileLet() (*Node, error) {
	var child interface{}
	var err error
	node := Node{structureTag: StrLetStatement, children: []interface{}{}}

	child, err = c.processKeyword(KwLet)
	if err != nil {
		return nil, targetNotFound(err)
	}
	node.children = append(node.children, child)

	child, err = c.processIdentifier()
	if err != nil {
		return nil, targetFound(err)
	}
	node.children = append(node.children, child)

	child, err = c.processSymbol("[")
	if err != nil {
		c.rollbackNextToken()
	} else {
		node.children = append(node.children, child)

		child, err = c.compileExpression()
		if err != nil {
			return nil, targetFound(err)
		}
		node.children = append(node.children, child)

		child, err = c.processSymbol("]")
		if err != nil {
			return nil, targetFound(err)
		}
		node.children = append(node.children, child)
	}

	child, err = c.processSymbol("=")
	if err != nil {
		return nil, targetFound(err)
	}
	node.children = append(node.children, child)

	child, err = c.compileExpression()
	if err != nil {
		return nil, targetFound(err)
	}
	node.children = append(node.children, child)

	child, err = c.processSymbol(";")
	if err != nil {
		return nil, targetFound(err)
	}
	node.children = append(node.children, child)

	return &node, nil
}

func (c *CompilationEngine) compileDo() (*Node, error) {
	var child interface{}
	var err error
	node := Node{structureTag: StrDoStatement, children: []interface{}{}}

	child, err = c.processKeyword(KwDo)
	if err != nil {
		return nil, targetNotFound(err)
	}
	node.children = append(node.children, child)

	// subrouutine call
	// 1. subroutineName(expressionList)
	// 2. (className|varName).subroutineName(expressionList)
	child, err = c.processIdentifier()
	if err != nil {
		return nil, targetFound(err)
	}
	node.children = append(node.children, child)

	child, err = c.processSymbol(".")
	if err != nil {
		c.rollbackNextToken()
	} else {
		node.children = append(node.children, child)

		child, err = c.processIdentifier()
		if err != nil {
			return nil, targetFound(err)
		}
		node.children = append(node.children, child)
	}

	child, err = c.processSymbol("(")
	if err != nil {
		return nil, targetFound(err)
	}
	node.children = append(node.children, child)

	child, err = c.compileExpressionList()
	if isTargetFound(err) {
		return nil, targetFound(err)
	} else {
		node.children = append(node.children, child)
	}

	child, err = c.processSymbol(")")
	if err != nil {
		return nil, targetFound(err)
	}
	node.children = append(node.children, child)

	return &node, nil
}

func (c *CompilationEngine) compileExpressionList() (*Node, error) {
	var child interface{}
	var err error
	node := Node{structureTag: StrExpressionList, children: []interface{}{}}

	child, err = c.compileExpression()
	if isTargetFound(err) {
		return nil, targetFound(err)
	} else if err != nil {
		// FIXME
		c.rollbackNextToken()

		c.rollbackNextToken()
		c.rollbackNextToken()
		return &node, nil
	} else {
		node.children = append(node.children, child)
	}

	for {
		child, err = c.processSymbol(",")
		if err != nil {
			c.rollbackNextToken()
			break
		}
		node.children = append(node.children, child)

		child, err = c.compileExpression()
		if err != nil {
			return nil, targetFound(err)
		}
		node.children = append(node.children, child)
	}

	return &node, nil
}

func (c *CompilationEngine) compileVarDec() (*Node, error) {
	var child interface{}
	var err error
	node := Node{structureTag: StrVarDec, children: []interface{}{}}

	child, err = c.processKeyword(KwVar)
	if err != nil {
		return nil, targetNotFound(err)
	}
	node.children = append(node.children, child)

	child, err = c.processIdentifier()
	if err != nil {
		return nil, targetFound(err)
	}
	node.children = append(node.children, child)

	child, err = c.processIdentifier()
	if err != nil {
		return nil, targetFound(err)
	}
	node.children = append(node.children, child)

	child, err = c.processSymbol(";")
	if err != nil {
		return nil, targetFound(err)
	}
	node.children = append(node.children, child)

	return &node, nil
}

func (c *CompilationEngine) compileSubroutine() (interface{}, error) {
	var child interface{}
	var err error
	node := Node{structureTag: StrSubroutineDec, children: []interface{}{}}

	child, err = c.processKeyword(KwFunction, KwMethod)
	if err != nil {
		return nil, targetNotFound(err)
	}
	node.children = append(node.children, child)

	child, err = c.processKeyword(KwInt, KwBoolean, KwChar, KwVoid)
	if err != nil {
		return nil, targetFound(err)
	}
	node.children = append(node.children, child)

	child, err = c.processIdentifier()
	if err != nil {
		return nil, targetFound(err)
	}
	node.children = append(node.children, child)

	child, err = c.processSymbol("(")
	if err != nil {
		return nil, targetFound(err)
	}
	node.children = append(node.children, child)

	child, err = c.compileParameterList()
	if err != nil {
		return nil, targetFound(err)
	}
	node.children = append(node.children, child)

	child, err = c.processSymbol(")")
	if err != nil {
		return nil, targetFound(err)
	}
	node.children = append(node.children, child)

	child, err = c.compileSubroutineBody()
	if err != nil {
		return nil, targetFound(err)
	}
	node.children = append(node.children, child)

	return &node, nil
}

func (c *CompilationEngine) compileParameterList() (interface{}, error) {
	var child interface{}
	var err error

	node := Node{structureTag: StrParameterList, children: []interface{}{}}

	for {
		child, err = c.processKeyword(KwInt, KwBoolean, KwChar)
		if err != nil {
			break
		}
		node.children = append(node.children, child)

		child, err = c.processIdentifier()
		if err != nil {
			return nil, targetFound(err)
		}
		node.children = append(node.children, child)

		child, err = c.processSymbol(",")
		if err != nil {
			break
		}
		node.children = append(node.children, child)
	}
	c.rollbackNextToken()

	return &node, nil
}

func (c *CompilationEngine) compileSubroutineBody() (interface{}, error) {
	var child interface{}
	var err error

	node := Node{structureTag: StrSubroutineBody, children: []interface{}{}}

	child, err = c.processSymbol("{")
	if err != nil {
		return nil, targetNotFound(err)
	}
	node.children = append(node.children, child)

	for {
		child, err = c.compileVarDec()
		if isTargetFound(err) {
			return nil, err
		} else if err != nil {
			c.rollbackNextToken()
			break
		}
		node.children = append(node.children, child)
	}

	child, err = c.compileStatements()
	if isTargetFound(err) {
		return nil, err
	} else if err != nil {
		c.rollbackNextToken()
	} else {
		node.children = append(node.children, child)
	}

	// child, err = c.processSymbol("}")
	// if err != nil {
	// 	return nil, err
	// }
	// node.children = append(node.children, child)

	return &node, nil
}

func (c *CompilationEngine) compileWhile() (interface{}, error) {
	token := c.nextToken()
	if token.tokenType != TkKeyWord ||
		KeyWordTag(token.value) != KwWhile {
		return nil, NewErrCompileFailed(token, string(KwWhile))
	}

	node := Node{structureTag: StrClass, children: []interface{}{token}}

	child, err := c.processSymbol("(")
	if err != nil {
		return nil, err
	}
	node.children = append(node.children, child)

	child, err = c.processSymbol(")")
	if err != nil {
		return nil, err
	}
	node.children = append(node.children, child)

	child, err = c.processSymbol("{")
	if err != nil {
		return nil, err
	}
	node.children = append(node.children, child)

	child, err = c.processSymbol("{")
	if err != nil {
		return nil, err
	}
	node.children = append(node.children, child)

	return &node, nil
}

// compileExpression returns compiled node tree and error
// expression is one of the following:
// - A constant
// - A variable name in scope. The variable may be static, field, local, or parameter
// - The this keyword, denoting the current object (cannot be used in functions)
// - An array element using the syntax arr[expression], where arr is a variable name of type Array in scope
// - A subroutine call that returns a non-void type
// - An expression prefixed by one of the expressions of unary operators - or ~:
// - An expression of the form expression op expression where op is one of the binary oprators(+,-,*,/,&,|,>,<,=)
// - (expression) an expression in parentheses
func (c *CompilationEngine) compileExpression() (_ *Node, err error) {
	var child interface{}
	iTokenBack := c.iToken
	defer func() {
		if err != nil {
			c.restoreNextToken(iTokenBack)
		}
	}()

	node := Node{structureTag: StrExpression, children: []interface{}{}}

	child, err = c.compileTerm()
	if err != nil {
		return nil, targetNotFound(err)
	}
	node.children = append(node.children, child)

	for {
		child, err = c.processSymbol("+", "-", "*", "/", "&", "|", "<", ">", "=")
		if err != nil {
			break
		}
		node.children = append(node.children, child)

		child, err = c.compileTerm()
		if err != nil {
			return nil, targetFound(err)
		}
		node.children = append(node.children, child)
	}

	return &node, nil
}

// intergerConst
// stringConst
// keywordConst (true, false, null, this)
// varName
// varName[expression]
// (expression)
// unaryOp term
// subroutineCall
func (c *CompilationEngine) compileTerm() (_ *Node, err error) {
	var child interface{}
	iTokenBack := c.iToken
	defer func() {
		if err != nil {
			c.restoreNextToken(iTokenBack)
		}
	}()

	node := Node{structureTag: StrTerm, children: []interface{}{}}

	// integerConst, or stringConst
	child, err = c.processTokenTag(TkIntConst, TkStringConst)
	if err == nil {
		node.children = append(node.children, child)
		return &node, nil
	}
	c.restoreNextToken(iTokenBack)
	node = Node{structureTag: StrTerm, children: []interface{}{}}

	// keywordConst
	child, err = c.processKeyword(KwTrue, KwFalse, KwNull, KwThis)
	if err == nil {
		node.children = append(node.children, child)
		return &node, nil
	}
	c.restoreNextToken(iTokenBack)
	node = Node{structureTag: StrTerm, children: []interface{}{}}

	// varName[expression]
	child, err = c.processIdentifier()
	if err == nil {
		node.children = append(node.children, child)

		// [
		child, err = c.processSymbol("[")
		if err == nil {
			node.children = append(node.children, child)

			// expression
			child, err = c.compileExpression()
			if err != nil {
				return &node, targetFound(err)
			} else {
				node.children = append(node.children, child)
			}

			// ]
			child, err = c.processSymbol("]")
			if err != nil {
				return &node, targetFound(err)
			} else {
				node.children = append(node.children, child)
			}

			return &node, nil
		}
	}
	c.restoreNextToken(iTokenBack)
	node = Node{structureTag: StrTerm, children: []interface{}{}}

	// (expression)
	// (
	child, err = c.processSymbol("(")
	if err == nil {
		node.children = append(node.children, child)

		// expression
		child, err = c.compileExpression()
		if err != nil {
			return &node, targetFound(err)
		} else {
			node.children = append(node.children, child)
		}

		// )
		child, err = c.processSymbol(")")
		if err != nil {
			return &node, targetFound(err)
		} else {
			node.children = append(node.children, child)
		}

		return &node, nil
	}
	c.restoreNextToken(iTokenBack)
	node = Node{structureTag: StrTerm, children: []interface{}{}}

	// unaryOp term
	child, err = c.processSymbol("-", "~")
	if err == nil {
		node.children = append(node.children, child)

		child, err = c.compileTerm()
		if err != nil {
			return &node, targetFound(err)
		} else {
			node.children = append(node.children, child)
		}
		return &node, nil
	}
	c.restoreNextToken(iTokenBack)
	node = Node{structureTag: StrTerm, children: []interface{}{}}

	// subroutineCall
	// 1. subroutineName(expressionList)
	// 2. (className|varName).subroutineName(expressionList)
	child, err = c.processIdentifier()
	if err == nil {
		node.children = append(node.children, child)

		child, err = c.processSymbol(".")
		if err != nil {
			c.rollbackNextToken()
		} else {
			node.children = append(node.children, child)

			child, err = c.processIdentifier()
			if err != nil {
				return nil, targetFound(err)
			}
			node.children = append(node.children, child)
		}

		child, err = c.processSymbol("(")
		if err == nil {
			node.children = append(node.children, child)

			child, err = c.compileExpressionList()
			fmt.Println(c.iToken)
			if isTargetFound(err) {
				return nil, targetFound(err)
			} else {
				node.children = append(node.children, child)
			}

			child, err = c.processSymbol(")")
			if err != nil {
				return nil, targetFound(err)
			}
			node.children = append(node.children, child)

			return &node, nil
		}
	}
	c.restoreNextToken(iTokenBack)
	node = Node{structureTag: StrTerm, children: []interface{}{}}

	// varName
	child, err = c.processIdentifier()
	if err == nil {
		node.children = append(node.children, child)
		return &node, nil
	}
	c.restoreNextToken(iTokenBack)

	if len(node.children) == 0 {
		return nil, NewErrCompileFailed(&Token{}, "term")
	}

	return &node, nil
}

func (c *CompilationEngine) processTokenTag(tags ...interface{}) (*Token, error) {
	token := c.nextToken()
	if token == nil {
		return nil, NewErrCompileFailed(token, toString(tags...))
	}
	for _, tag := range tags {
		if token.tokenType == tag {
			return token, nil
		}
	}
	return nil, NewErrCompileFailed(token, toString(tags...))
}

func (c *CompilationEngine) processKeyword(kws ...interface{}) (*Token, error) {
	token := c.nextToken()
	if token == nil {
		return nil, NewErrCompileFailed(token, toString(kws...))
	}
	if token.tokenType != TkKeyWord {
		return nil, NewErrCompileFailed(token, toString(kws...))
	}
	for _, kw := range kws {
		switch _kw := kw.(type) {
		case KeyWordTag:
			if token.value == string(_kw) {
				return token, nil
			}
		}
	}
	return nil, NewErrCompileFailed(token, toString(kws...))
}

func (c *CompilationEngine) processSymbol(ss ...string) (*Token, error) {
	token := c.nextToken()
	if token == nil {
		return nil, NewErrCompileFailed(&Token{value: "nil"}, strings.Join(ss, ","))
	}
	if token.tokenType == TkSymbol {
		for _, s := range ss {
			if s == token.value {
				return token, nil
			}
		}
	}
	return nil, NewErrCompileFailed(token, strings.Join(ss, ","))
}

func (c *CompilationEngine) processIdentifier() (*Token, error) {
	token := c.nextToken()
	if token == nil {
		return nil, NewErrCompileFailed(token, string(TkIdentifier))
	}
	if token.tokenType != TkIdentifier {
		return nil, NewErrCompileFailed(token, string(TkIdentifier))
	}
	return token, nil
}

func toString(as ...interface{}) string {
	ret := []string{}
	for _, v := range as {
		switch _v := v.(type) {
		case TokenTag:
			ret = append(ret, string(_v))
		case KeyWordTag:
			ret = append(ret, string(_v))
		case StructureTag:
			ret = append(ret, string(_v))
		}
	}

	return strings.Join(ret, ",")
}
