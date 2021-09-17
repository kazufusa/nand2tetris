package vm

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

type VMTranslator struct {
	parsers    []*Parser
	codeWriter *CodeWriter
}

func NewVMTranslator(in string) (*VMTranslator, error) {
	out := strings.TrimSuffix(in, filepath.Ext(in)) + ".asm"
	fs, err := findVMs(in)
	if err != nil {
		return nil, err
	}

	var parsers []*Parser
	for _, f := range fs {
		fileName := strings.TrimSuffix(filepath.Base(f), ".vm")
		_vm, err := os.ReadFile(f)
		if err != nil {
			return nil, err
		}
		parser, err := NewParser(string(_vm), fileName)
		if err != nil {
			return nil, err
		}
		parsers = append(parsers, parser)
	}

	vm := ""
	for _, f := range fs {
		_vm, err := os.ReadFile(f)
		if err != nil {
			return nil, err
		}
		vm += "\n" + string(_vm)
	}

	return &VMTranslator{
		parsers:    parsers,
		codeWriter: NewCodeWriter(out),
	}, nil
}

func (t *VMTranslator) Conv() error {
	for _, parser := range t.parsers {
		t.codeWriter.setFileName(parser.fileName)
		err := t.conv(parser)
		if err != nil {
			return err
		}
	}
	return t.codeWriter.close()
}

func (t *VMTranslator) conv(parser *Parser) error {
	for {
		t.codeWriter.write("// " + parser.line() + "\n")
		cmd, err := parser.commandType()
		if err != nil {
			return err
		}
		arg1, err1 := parser.arg1()
		arg2, err2 := parser.arg2()

		// error handling
		switch cmd {
		case C_ALITHMETIC, C_LABEL, C_GOTO, C_IF:
			if err1 != nil {
				return err1
			}
		case C_RETURN:
		default:
			if err1 != nil {
				return fmt.Errorf("'%s': %s", parser.line(), err1)
			}
			if err2 != nil {
				return fmt.Errorf("'%s': %s", parser.line(), err2)
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
		case C_IF:
			t.codeWriter.writeIf(arg1)
		case C_FUNCTION:
			t.codeWriter.writeFunction(arg1, arg2)
		case C_RETURN:
			t.codeWriter.writeReturn()
		case C_CALL:
			t.codeWriter.writeCall(arg1, arg2)
		default:
		}

		if parser.hasMoreLines() {
			t.codeWriter.write("\n")
			parser.advance()
		} else {
			return nil
		}
	}
}

func findVMs(f string) ([]string, error) {
	s, err := os.Stat(f)
	if err != nil {
		return nil, err
	}
	if !s.IsDir() {
		return []string{f}, nil
	}
	return filepath.Glob(filepath.Join(f, "*.vm"))
}
