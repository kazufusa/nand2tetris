package vm

import (
	"os"
	"path/filepath"
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
		name     string
		expected map[int]int
	}{
		{"SimpleTest", map[int]int{
			1:   500,
			2:   1000,
			3:   3000,
			4:   4000,
			0:   257,
			256: -1, 257: 17}},
		{"SimpleAdd", map[int]int{
			1: 500,
			2: 1000,
			3: 3000,
			4: 4000,
			0: 257, 256: 15, 257: 8}},
		{"StackTest", map[int]int{
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
		}},
		{"BasicTest", map[int]int{
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
		}},
		{"PointerTest", map[int]int{
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
		}},
		{"StaticTest", map[int]int{
			1:   500,
			2:   1000,
			3:   3000,
			4:   4000,
			0:   257,
			17:  111,
			19:  333,
			24:  888,
			256: 1110,
			257: 888,
			258: 888,
		}},
		{"BasicLoop", map[int]int{
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
		}},
		{"FibonacciSeries", map[int]int{
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
		}},
		// {"NestedCall", map[int]int{
		// 	1:   500,
		// 	2:   1000,
		// 	3:   3000,
		// 	4:   4000,
		// 	0:   257,
		// 	17:  111,
		// 	19:  333,
		// 	24:  888,
		// 	256: 1110,
		// 	257: 888,
		// 	258: 888,
		// }},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			vm := filepath.Join("examples/final_ram_state", tt.name+".vm")
			asm := filepath.Join("examples/final_ram_state", tt.name+".asm")
			defer os.Remove(asm)

			translator, err := NewVMTranslator(vm)
			assert.NoError(t, err)

			err = translator.Conv()
			require.NoError(t, err)

			actual := EmulatedResult(t, asm)
			// assert.Equal(t, tt.expected, actual)
			if !cmp.Equal(tt.expected, actual) {
				t.Error(cmp.Diff(tt.expected, actual))
			}
		})
	}
}

func EmulatedResult(t *testing.T, fasm string) map[int]int {
	asm, err := assembler.NewAssembler(fasm)
	require.NoError(t, err)
	ret, err := asm.Assemble()
	require.NoError(t, err)
	com := computer.NewEmulator(ret)
	com.WriteRom(0, 256)
	com.WriteRom(1, 500)
	com.WriteRom(2, 1000)
	com.WriteRom(3, 3000)
	com.WriteRom(4, 4000)
	for i := 0; i < 10000; i++ {
		com.FetchAndExecute(logic.O)
		if com.IsFinished() {
			break
		}
	}
	return com.ROMMap()
}
