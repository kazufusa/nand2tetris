package vm

import (
	"os"
	"path/filepath"
	"strings"
)

type VMTranslator struct {
	parser     *Parser
	codeWriter *CodeWriter
}

func NewVMTranslator(in string) (*VMTranslator, error) {
	out := strings.TrimSuffix(in, filepath.Ext(in)) + ".asm"
	vm, err := os.ReadFile(in)
	if err != nil {
		return nil, err
	}
	parser, err := NewParser(string(vm))
	if err != nil {
		return nil, err
	}
	return &VMTranslator{
		parser:     parser,
		codeWriter: NewCodeWriter(out),
	}, nil
}

func (t *VMTranslator) Conv() error {
	for {
		t.codeWriter.write("// " + t.parser.line() + "\n")
		cmd, err := t.parser.commandType()
		if err != nil {
			return err
		}
		arg1, err := t.parser.arg1()
		if err != nil {
			return err
		}
		switch cmd {
		case C_PUSH, C_POP:
			arg2, err := t.parser.arg2()
			if err != nil {
				return err
			}
			t.codeWriter.writePushPop(cmd, arg1, arg2)
		case C_ALITHMETIC:
			t.codeWriter.writeArithmetic(arg1)
		default:
		}

		if t.parser.hasMoreLines() {
			t.codeWriter.write("\n")
			t.parser.advance()
		} else {
			return t.codeWriter.close()
		}
	}
}
