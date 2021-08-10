package assembler

import (
	"fmt"
	"io/ioutil"
	"math"
	"strconv"
)

type Assembler struct {
	parser *Parser
	code   *Code
	st     *SymbolTable
}

func NewAssembler(f string) (*Assembler, error) {
	b, err := ioutil.ReadFile(f)
	if err != nil {
		return nil, err
	}
	parser := NewParser(string(b))
	code := &Code{}
	st := NewSymbolTable()
	return &Assembler{parser: parser, code: code, st: st}, nil
}

func (asm *Assembler) Assemble() (string, error) {
	ret := ""
	ln := 0
	for {
		inst := asm.parser.InstructionType()
		switch inst {
		case L_INSTRUCTION:
			symbol := asm.parser.Symbol()
			if !isNum(symbol) {
				asm.st.AddEntry(symbol, ln, true)
			}
		}

		if inst == A_INSTRUCTION || inst == C_INSTRUCTION {
			ln += 1
		}

		if asm.parser.HasMoreLines() {
			asm.parser.Advance()
		} else {
			break
		}
	}

	asm.parser.Reset()
	for {
		inst := asm.parser.InstructionType()
		switch inst {
		case A_INSTRUCTION:
			symbol := asm.parser.Symbol()
			if !isNum(symbol) {
				asm.st.AddEntry(symbol, ln, false)
			}
		}

		if asm.parser.HasMoreLines() {
			asm.parser.Advance()
		} else {
			break
		}
	}

	asm.parser.Reset()
	for {
		inst := asm.parser.InstructionType()
		switch inst {
		case A_INSTRUCTION:
			symbol := asm.parser.Symbol()
			if !isNum(symbol) {
				_symbol, err := asm.st.GetAddress(symbol)
				if err != nil {
					return "", err
				}
				symbol = fmt.Sprintf("%d", _symbol)
			}
			ret += fmt.Sprintf("%s\n", toBinary(symbol))
		case C_INSTRUCTION:
			ret += fmt.Sprintf("111%s%s%s\n",
				asm.code.Comp(asm.parser.Comp()),
				asm.code.Dest(asm.parser.Dest()),
				asm.code.Jump(asm.parser.Jump()),
			)
		}
		if asm.parser.HasMoreLines() {
			asm.parser.Advance()
		} else {
			break
		}
	}

	return ret, nil
}

func isNum(s string) bool {
	_, err := strconv.ParseInt(s, 10, 64)
	return err == nil
}

func toBinary(num string) string {
	n, _ := strconv.ParseInt(num, 10, 64)
	return fmt.Sprintf("0%d%d%d%d%d%d%d%d%d%d%d%d%d%d%d",
		uint8((n&int64(math.Pow(2, 14)))>>14),
		uint8((n&int64(math.Pow(2, 13)))>>13),
		uint8((n&int64(math.Pow(2, 12)))>>12),
		uint8((n&int64(math.Pow(2, 11)))>>11),
		uint8((n&int64(math.Pow(2, 10)))>>10),
		uint8((n&int64(math.Pow(2, 9)))>>9),
		uint8((n&int64(math.Pow(2, 8)))>>8),
		uint8((n&int64(math.Pow(2, 7)))>>7),
		uint8((n&int64(math.Pow(2, 6)))>>6),
		uint8((n&int64(math.Pow(2, 5)))>>5),
		uint8((n&int64(math.Pow(2, 4)))>>4),
		uint8((n&int64(math.Pow(2, 3)))>>3),
		uint8((n&int64(math.Pow(2, 2)))>>2),
		uint8((n&int64(math.Pow(2, 1)))>>1),
		uint8((n&int64(math.Pow(2, 0)))>>0),
	)
}
