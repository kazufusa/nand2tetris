package vm

import (
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/google/go-cmp/cmp"
	logic "github.com/kazufusa/nand2tetris/01_Boolean_Logic"
	computer "github.com/kazufusa/nand2tetris/05_Computer_Architecture"
	assembler "github.com/kazufusa/nand2tetris/06_Assembler"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// func TestAsm(t *testing.T) {
// 	var tests = []string{"NestedCall"}
// 	for _, tt := range tests {
// 		tt := tt
// 		t.Run(tt, func(t *testing.T) {
// 			vm := filepath.Join("examples/asm_test", tt+".vm")
// 			asm := filepath.Join("examples/asm_test", tt+".asm")
// 			expectedAsm := filepath.Join("examples/asm_test", tt+"expected.asm")
// 			fmt.Println(vm, asm, expectedAsm)
//
// 			translator, err := NewVMTranslator(vm)
// 			require.NoError(t, err)
// 			err = translator.Conv()
// 			require.NoError(t, err)
// 		})
// 	}
// }

func TestFinalStateOfRam(t *testing.T) {
	var tests = []struct {
		name        string
		expected    map[int]int
		prepare     func(*computer.Computer)
		noBootstrap bool
	}{
		{"SimpleTest.vm", map[int]int{
			1:   500,
			2:   1000,
			3:   3000,
			4:   4000,
			0:   257,
			256: -1, 257: 17,
		}, defaultPreparation, true},
		{"SimpleAdd.vm", map[int]int{
			1: 500,
			2: 1000,
			3: 3000,
			4: 4000,
			0: 257, 256: 15, 257: 8,
		}, defaultPreparation, true},
		{"StackTest.vm", map[int]int{
			1:   500,
			2:   1000,
			3:   3000,
			4:   4000,
			0:   266,
			256: -1,
			257: 0,
			258: 0,
			259: 0,
			260: -1,
			261: 0,
			262: -1,
			263: 0,
			264: 0,
			265: -91,
			266: 82,
			267: 112,
		}, defaultPreparation, true},
		{"BasicTest.vm", map[int]int{
			1:    500,
			2:    1000,
			3:    3000,
			4:    4000,
			0:    257,
			11:   510,
			13:   4002,
			256:  472,
			257:  510,
			258:  36,
			500:  10,
			1001: 21,
			1002: 22,
			3006: 36,
			4005: 45,
			4002: 42,
		}, defaultPreparation, true},
		{"PointerTest.vm", map[int]int{
			1:    500,
			2:    1000,
			3:    3030,
			4:    3040,
			0:    257,
			13:   3046,
			256:  6084,
			257:  46,
			3032: 32,
			3046: 46,
		}, defaultPreparation, true},
		{"StaticTest.vm", map[int]int{
			1:   500,
			2:   1000,
			3:   3000,
			4:   4000,
			0:   257,
			16:  888,
			17:  333,
			18:  111,
			256: 1110,
			257: 888,
			258: 888,
		}, defaultPreparation, true},
		{"BasicLoop.vm", map[int]int{
			1:    500,
			2:    1000,
			3:    3000,
			4:    4000,
			0:    257,
			13:   1000,
			256:  15,
			257:  1,
			500:  15,
			1000: 0,
		}, defaultPreparation, true},
		{"FibonacciSeries.vm", map[int]int{
			1:    500,
			2:    1000,
			3:    3000,
			4:    4004,
			0:    256,
			13:   1000,
			256:  0,
			257:  1,
			1000: 0,
			1001: 4000,
			4000: 0,
			4001: 1,
			4002: 1,
			4003: 2,
			4004: 3,
			4005: 5,
		}, defaultPreparation, true},
		{"SimpleFunction.vm", map[int]int{
			1:   305,
			2:   300,
			3:   3010,
			4:   4010,
			0:   311,
			13:  313,
			14:  10000,
			310: 0,
			311: 37,
			312: 10000,
			313: 305,
			314: 300,
			315: 3010,
			316: 4010,
			317: 0,
			318: 0,
		},
			func(com *computer.Computer) {
				com.WriteRom(0, 317)
				com.WriteRom(1, 317)
				com.WriteRom(2, 310)
				com.WriteRom(3, 3000)
				com.WriteRom(4, 4000)
				com.WriteRom(310, 1234)
				com.WriteRom(311, 37)
				com.WriteRom(312, 10000)
				com.WriteRom(313, 305)
				com.WriteRom(314, 300)
				com.WriteRom(315, 3010)
				com.WriteRom(316, 4010)
			}, true,
		},
		{"StaticsTest", map[int]int{
			// initial SP: 256, sys.Init: 352, WHILE: 366, retAddr1: 54
			// retaddr2: 418
			0:   263,
			1:   261,
			2:   256,
			3:   0,
			4:   0,
			5:   0,
			13:  263,
			14:  628,
			16:  6,
			17:  8,
			18:  23,
			19:  15,
			256: 54,
			257: 0,
			258: 0,
			259: 0,
			260: 0,
			261: -2,
			262: 8,
			263: 261,
			264: 256,
			265: 0,
			266: 0,
			267: 8,
			268: 15,
		}, nil, false},
		{"NestedCall", map[int]int{
			0:   261,
			1:   261,
			2:   256,
			3:   4000,
			4:   5000,
			5:   135,
			6:   246,
			13:  262,
			14:  130,
			256: 54,
			257: 0,
			258: 0,
			259: 0,
			260: 0,
			261: 246,
			262: 261,
			263: 256,
			264: 4000,
			265: 5000,
			266: 0,
			267: 200,
			268: 40,
			269: 6,
			270: 0,
			271: 246,
			272: 246,
			273: 46,
			274: 6,
			275: 0,
			276: 5001,
			277: 135,
			278: 12,
		}, nil, false},
		{"FibonacciElement", map[int]int{
			0:   262,
			1:   261,
			2:   256,
			3:   0,
			4:   0,
			13:  263,
			14:  431,
			256: 54,
			257: 0,
			258: 0,
			259: 0,
			260: 0,
			261: 3,
			262: 431,
			263: 261,
			264: 256,
			265: 0,
			266: 0,
			267: 3,
			268: 2,
			269: 312,
			270: 267,
			271: 261,
			272: 0,
			273: 0,
			274: 2,
			275: 1,
			276: 312,
			277: 274,
			278: 268,
			279: 0,
			280: 0,
			281: 1,
			282: 1,
			283: 312,
			284: 281,
			285: 275,
			286: 0,
			287: 0,
			288: 1,
			289: 2,
		}, nil, false},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			asm := filepath.Join("examples", strings.TrimSuffix(tt.name, ".vm")+".asm")
			defer os.Remove(asm)

			vm := filepath.Join("examples", tt.name)
			translator, err := NewVMTranslator(vm)
			require.NoError(t, err)
			if tt.noBootstrap {
				translator.codeWriter.asm = ""
			}
			err = translator.Conv()
			require.NoError(t, err)

			actual := EmulatedResult(t, asm, tt.prepare)
			if !cmp.Equal(tt.expected, actual) {
				t.Error(cmp.Diff(tt.expected, actual))
			}
		})
	}
}

func defaultPreparation(com *computer.Computer) {
	com.WriteRom(0, 256)
	com.WriteRom(1, 500)
	com.WriteRom(2, 1000)
	com.WriteRom(3, 3000)
	com.WriteRom(4, 4000)
}

func EmulatedResult(t *testing.T, fasm string, prepare func(*computer.Computer)) map[int]int {
	asm, err := assembler.NewAssembler(fasm)
	require.NoError(t, err)
	ret, err := asm.Assemble()
	require.NoError(t, err)
	com := computer.NewEmulator(ret)
	if prepare != nil {
		prepare(com)
	}
	var i int
	for i = 0; i < 10000; i++ {
		com.FetchAndExecute(logic.O)
		if com.IsFinished() {
			break
		}
	}
	return com.ROMMap()
}

func TestFindVMs(t *testing.T) {
	actual, _ := findVMs("examples/BasicLoop.vm")
	expected := []string{"examples/BasicLoop.vm"}
	assert.Equal(t, expected, actual)

	actual, _ = findVMs("examples/StaticsTest")
	expected = []string{
		"examples/StaticsTest/Class1.vm",
		"examples/StaticsTest/Class2.vm",
		"examples/StaticsTest/Sys.vm",
	}
	assert.Equal(t, expected, actual)
}
