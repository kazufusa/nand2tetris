package vm

import (
	"errors"
	"regexp"
	"strconv"
	"strings"
)

// Specification
// - Push/pop commands transfer data between the stack and memory segments
// - Arithmetic-logical commands perform arithmetic and logical operations
//   - add, sub, neg, eq, gt, lt, and, or, no
// - Branching commands facilitate conditional and unconditional branching operations
// - Function commands facilitate function call-and-return operations

type Command int

const (
	C_ALITHMETIC Command = iota
	C_PUSH
	C_POP
	C_LABEL
	C_GOTO
	C_IF
	C_FUNCTION
	C_RETURN
	C_CALL
)

var (
	reComment       = regexp.MustCompile(`//[^\n]*`)
	reForwardSpace  = regexp.MustCompile(`(?m)^\s*`)
	reBackwardSpace = regexp.MustCompile(`(?m)\s*$`)
	reSpace         = regexp.MustCompile(`(?m) +`)
)

type Parser struct {
	lines []string
	count int
}

func NewParser(s string) (*Parser, error) {
	s = reComment.ReplaceAllString(s, "")
	s = reForwardSpace.ReplaceAllString(s, "")
	s = reBackwardSpace.ReplaceAllString(s, "")
	s = reSpace.ReplaceAllString(s, " ")
	s = strings.ReplaceAll(s, "\r", "")
	parser := Parser{}
	parser.lines = strings.Split(s, "\n")
	for i := len(parser.lines) - 1; i >= 0; i-- {
		if parser.lines[i] == "" {
			parser.lines = append(parser.lines[:i], parser.lines[i+1:]...)
		}
	}
	return &parser, nil
}

func (p *Parser) line() string {
	return p.lines[p.count]
}

func (p *Parser) hasMoreLines() bool {
	return p.count < (len(p.lines) - 1)
}

func (p *Parser) advance() {
	p.count++
}

func (p *Parser) commandType() (cmd Command, err error) {
	args := strings.Split(p.lines[p.count], " ")
	switch args[0] {
	case "add", "sub", "neg", "eq", "gt", "lt", "and", "or", "not":
		return C_ALITHMETIC, nil
	case "push":
		return C_PUSH, nil
	case "pop":
		return C_POP, nil
	case "label":
		return C_LABEL, nil
	case "goto":
		return C_GOTO, nil
	case "if-goto":
		return C_IF, nil
	case "function":
		return C_FUNCTION, nil
	case "return":
		return C_RETURN, nil
	case "call":
		return C_CALL, nil
	}

	return C_POP, errors.New("invalid command")
}

func (p *Parser) arg1() (string, error) {
	args := strings.Split(p.lines[p.count], " ")
	if cmd, err := p.commandType(); err == nil && cmd == C_ALITHMETIC {
		return args[0], nil
	}
	if len(args) < 2 {
		return "", errors.New("arg1 is not exists")
	}
	return args[1], nil
}

func (p *Parser) arg2() (int, error) {
	args := strings.Split(p.lines[p.count], " ")
	if len(args) == 3 {
		return strconv.Atoi(args[2])
	}
	return 0, errors.New("arg2 is not exists")
}
