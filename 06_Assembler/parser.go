package assembler

import (
	"regexp"
	"strings"
)

type InstructionType int

const (
	A_INSTRUCTION InstructionType = iota
	C_INSTRUCTION
	L_INSTRUCTION
)

var (
	reComment = regexp.MustCompile(`//[^\n]*`)

	reA = regexp.MustCompile(`@(.+)`)
	reL = regexp.MustCompile(`^\((.+)\)$`)

	// dest=comp;jmp
	reDestCompJmp = regexp.MustCompile(`^(.*)=(.*);(.*)$`)
	reCompJmp     = regexp.MustCompile(`^(.*);(.*)$`)
	reDestComp    = regexp.MustCompile(`^(.*)=(.*$)`)
	reComp        = regexp.MustCompile(`^(.*$)`)
)

type IParser interface {
	HasMoreLines() bool
	Advance()
	InstructionType() InstructionType
	Symbol() string
	Dest() string
	Comp() string
	Jump() string
}

type Parser struct {
	lines []string
	count int
}

var _ IParser = (*Parser)(nil)

func NewParser(s string) *Parser {
	s = reComment.ReplaceAllString(s, "")
	s = strings.ReplaceAll(s, "\r", "")
	s = strings.ReplaceAll(s, " ", "")
	lines := strings.Split(s, "\n")
	for i := len(lines) - 1; i >= 0; i-- {
		if lines[i] == "" {
			lines = append(lines[:i], lines[i+1:]...)
		}
	}
	return &Parser{lines: lines}
}

func (p *Parser) HasMoreLines() bool {
	return p.count < len(p.lines)-1
}

func (p *Parser) Advance() {
	if p.HasMoreLines() {
		p.count = p.count + 1
	}
}

func (p *Parser) InstructionType() InstructionType {
	l := p.lines[p.count]
	switch []rune(l)[0] {
	case '@':
		return A_INSTRUCTION
	case '(':
		return L_INSTRUCTION
	default:
		return C_INSTRUCTION
	}
}

func (p *Parser) Symbol() string {
	l := p.lines[p.count]
	switch p.InstructionType() {
	case A_INSTRUCTION:
		m := (reA.FindStringSubmatch(l))
		if len(m) == 2 {
			return m[1]
		}
	case L_INSTRUCTION:
		m := reL.FindStringSubmatch(l)
		if len(m) == 2 {
			return m[1]
		}
	}
	return ""
}

func (p *Parser) Dest() string {
	l := p.lines[p.count]
	if m := reDestCompJmp.FindStringSubmatch(l); len(m) > 0 {
		return m[1]
	} else if m := reDestComp.FindStringSubmatch(l); len(m) > 0 {
		return m[1]
	}
	return ""
}

func (p *Parser) Comp() string {
	l := p.lines[p.count]
	if m := reDestCompJmp.FindStringSubmatch(l); len(m) > 0 {
		return m[2]
	} else if m := reCompJmp.FindStringSubmatch(l); len(m) > 0 {
		return m[1]
	} else if m := reDestComp.FindStringSubmatch(l); len(m) > 0 {
		return m[2]
	}
	return l
}

func (p *Parser) Jump() string {
	l := p.lines[p.count]
	if m := reDestCompJmp.FindStringSubmatch(l); len(m) > 0 {
		return m[3]
	} else if m := reCompJmp.FindStringSubmatch(l); len(m) > 0 {
		return m[2]
	}
	return ""
}

func (p *Parser) Reset() {
	p.count = 0
}
