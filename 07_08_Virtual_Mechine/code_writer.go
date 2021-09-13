package vm

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

const (
	SegLocal    = "local"
	SegArgument = "argument"
	SegThis     = "this"
	SegThat     = "that"
	SegPointer  = "pointer"
	SegTemp     = "temp"
	SegConstant = "constant"
	SegStatic   = "static"

	AAdd = "add"
	ASub = "sub"
	ANeg = "neg"
	AEq  = "eq"
	AGt  = "gt"
	ALt  = "lt"
	AAnd = "and"
	AOr  = "or"
	ANot = "not"
)

type CodeWriteSpecification interface {
	// setFileName(f string)
	writeArithmetic(command string) error
	writePushPop(command Command, arg1 string, arg2 int) error
	// writeLabel(label string)
	// writeGoto(label string)
	// writeIf(label string)
	// writeFunction(name string, nVars int)
	// writeCall(name string, nArgs int)
	// writeReturn()
	close() error
}

var _ CodeWriteSpecification = (*CodeWriter)(nil)

type CodeWriter struct {
	out  string
	name string
	asm  string
	n    int
}

func NewCodeWriter(out string) *CodeWriter {
	name := strings.TrimSuffix(filepath.Base(out), filepath.Ext(out))
	return &CodeWriter{out: out, asm: "", name: name}
}

func (c *CodeWriter) write(s string) {
	c.asm += s
}

func (c *CodeWriter) incN() {
	c.n++
}

func (c *CodeWriter) writeArithmetic(cmd string) error {
	defer c.incN()
	var err error
	var s string

	switch cmd {
	case "add", "sub", "and", "or":
		s, err = twoValuesCommand(cmd)
	case "neg", "not":
		s, err = singleValueCommand(cmd)
	case "eq", "gt", "lt":
		s, err = twoValuesJumpCommand(cmd, c.n)
	default:
		err = errors.New("unknown command")
	}
	if err != nil {
		return err
	}
	c.asm += s
	return nil
}

// RAM addresses
// 0-15     Sixteen virtual registers
// 16-255   Static variables
// 256-2047 Stack
//
// SP          RAM[0]      Stack Pointer
// LCL         RAM[1]      Base address of the local segment
// ARG         RAM[2]      Base address of the argument segment
// THIS        RAM[3]      Base address of the this segment
// THAT        RAM[4]      Base address of the that segment
// TEMP        RAM[5-12]   holds the temp segment
// R13,R14,R15 RAM[13-15]  if the assembly code generated by the VM translator needs variables, it can use these registers
//
// segments:
//   local, argument, this, that:
//     LCL, ARG, THIS, THAT, respectively
//   pointer:
//     pointer 0 to THIS, pointer 1 to THAT
//     example: push pointer 0: push THIS to stack
//   temp
//     TEMP, addresses 5 to 12
//   constant:
//     stack
//   static:
//     addresses 16 to 255
func (c *CodeWriter) writePushPop(command Command, arg1 string, arg2 int) error {
	defer c.incN()

	asm := ""
	var err error
	switch command {
	case C_PUSH:
		asm, err = push(arg1, arg2, c.name)
	case C_POP:
		asm, err = pop(arg1, arg2, c.name)
	default:
		return errors.New("unknown command")
	}
	if err != nil {
		return err
	}
	c.asm += asm
	return nil
}

func (c *CodeWriter) close() error {
	return os.WriteFile(c.out, []byte(c.asm), 0777)
}

func singleValueCommand(arg1 string) (string, error) {
	var cmd string
	switch arg1 {
	case "neg":
		cmd = "-M"
	case "not":
		cmd = "!M"
	default:
		return "", errors.New("unknown command")
	}
	return fmt.Sprintf(`@SP
M=M-1
A=M
M=%s
@SP
M=M+1
`, cmd), nil
}

func twoValuesCommand(arg1 string) (string, error) {
	var cmd string
	switch arg1 {
	case "add":
		cmd = "D+M"
	case "sub":
		cmd = "M-D"
	case "and":
		cmd = "D&M"
	case "or":
		cmd = "D|M"
	default:
		return "", errors.New("unknown command")
	}
	return fmt.Sprintf(`@SP
M=M-1
A=M
D=M
@SP
A=M-1
D=%s
@SP
A=M-1
M=D
`, cmd), nil
}

func twoValuesJumpCommand(arg1 string, n int) (string, error) {
	jmp := ""
	switch arg1 {
	case "eq":
		jmp = "JEQ"
	case "gt":
		jmp = "JGT"
	case "lt":
		jmp = "JLT"
	default:
		return "", errors.New("unknown command")
	}

	return fmt.Sprintf(`@SP
A=M
A=A-1
A=A-1
D=M
A=A+1
D=D-M
@SP
M=M-1
A=M-1
M=-1
@TRUE%d
D;%s
@SP
A=M-1
M=0
(TRUE%d)
`, n, jmp, n), nil
	// 	return fmt.Sprintf(`@SP
	// A=M
	// A=A-1
	// A=A-1
	// D=M
	// A=A+1
	// D=D-M
	// @SP
	// M=M-1
	// M=M-1
	// @TRUE%d
	// D;%s
	// @SP
	// A=M
	// M=0
	// @END%d
	// 0;JMP
	//
	// (TRUE%d)
	// @SP
	// A=M
	// M=-1
	//
	// (END%d)
	// `, n, jmp, n, n, n), nil
}

func push(arg1 string, arg2 int, name string) (string, error) {
	switch arg1 {
	case SegLocal, SegArgument, SegThis, SegThat:
		return pushVirtualSeg(arg1, arg2, name)
	case SegPointer, SegTemp, SegStatic:
		return pushPointer(arg1, arg2, name)
	case SegConstant:
		return pushConstant(arg1, arg2, name)
	default:
		return "", errors.New("unknown segment")
	}
}

func pop(arg1 string, arg2 int, name string) (string, error) {
	switch arg1 {
	case SegLocal, SegArgument, SegThis, SegThat:
		return popVirtualSeg(arg1, arg2, name)
	case SegPointer, SegTemp, SegStatic:
		return popPointer(arg1, arg2, name)
	default:
		return "", errors.New("unknown segment")
	}
}

func pushVirtualSeg(arg1 string, arg2 int, name string) (string, error) {
	var addr string
	switch arg1 {
	case SegLocal:
		addr = "LCL"
	case SegArgument:
		addr = "ARG"
	case SegThis:
		addr = "THIS"
	case SegThat:
		addr = "THAT"
	default:
		return "", errors.New("unknown segment")
	}
	return fmt.Sprintf(`@%d
D=A
@%s
A=D+M
D=M
@SP
A=M
M=D
@SP
M=M+1
`, arg2, addr), nil
}

func popVirtualSeg(arg1 string, arg2 int, name string) (string, error) {
	var addr string
	switch arg1 {
	case SegLocal:
		addr = "LCL"
	case SegArgument:
		addr = "ARG"
	case SegThis:
		addr = "THIS"
	case SegThat:
		addr = "THAT"
	default:
		return "", errors.New("unknown segment")
	}
	return fmt.Sprintf(`@%s
D=M
@%d
D=D+A
@R13
M=D
@SP
M=M-1
A=M
D=M
@R13
A=M
M=D
`, addr, arg2), nil
}

func pushPointer(arg1 string, arg2 int, name string) (string, error) {
	var addr string
	switch arg1 {
	case SegPointer:
		if arg2 == 0 {
			addr = "THIS"
		} else if arg2 == 1 {
			addr = "THAT"
		} else {
			return "", errors.New("unknown segment")
		}
	case SegTemp:
		addr = fmt.Sprintf("R%d", 5+arg2)
	case SegStatic:
		addr = fmt.Sprintf("%d", arg2+16)
	}
	return fmt.Sprintf(`@%s
D=M
@SP
A=M
M=D
@SP
M=M+1
`, addr), nil
}

func popPointer(arg1 string, arg2 int, name string) (string, error) {
	var addr string
	switch arg1 {
	case SegPointer:
		if arg2 == 0 {
			addr = "THIS"
		} else if arg2 == 1 {
			addr = "THAT"
		} else {
			return "", errors.New("unknown segment")
		}
	case SegTemp:
		addr = fmt.Sprintf("R%d", 5+arg2)
	case SegStatic:
		addr = fmt.Sprintf("%d", arg2+16)
	}
	return fmt.Sprintf(`@SP
M=M-1
A=M
D=M
@%s
M=D
`, addr), nil
}

func pushConstant(arg1 string, arg2 int, name string) (string, error) {
	return fmt.Sprintf(`@%d
D=A
@SP
A=M
M=D
@SP
M=M+1
`, arg2), nil
}
