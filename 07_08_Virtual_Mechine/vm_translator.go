package vm

import (
	"fmt"
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
		arg1, err1 := t.parser.arg1()
		arg2, err2 := t.parser.arg2()

		// error handling
		switch cmd {
		case C_ALITHMETIC, C_LABEL, C_GOTO:
			if err1 != nil {
				return err1
			}
		case C_RETURN:
		default:
			if err1 != nil {
				return fmt.Errorf("'%s': %s", t.parser.line(), err1)
			}
			if err2 != nil {
				return fmt.Errorf("'%s': %s", t.parser.line(), err2)
			}
		}

		switch cmd {
		case C_PUSH, C_POP:
			t.codeWriter.writePushPop(cmd, arg1, arg2)
		case C_ALITHMETIC:
			t.codeWriter.writeArithmetic(arg1)
		case C_LABEL:
			t.codeWriter.writeLabel(arg1)
		case C_GOTO:
			t.codeWriter.writeGoto(arg1)
		case C_FUNCTION:
			t.codeWriter.writeFunction(arg1, arg2)
		case C_RETURN:
			t.codeWriter.writeReturn()
		case C_CALL:
			t.codeWriter.writeCall(arg1, arg2)
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
