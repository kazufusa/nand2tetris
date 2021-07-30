package computer

import (
	"fmt"
	"strings"
	"testing"

	logic "github.com/kazufusa/nand2tetris/01_Boolean_Logic"
	"github.com/stretchr/testify/assert"
)

func BenchmarkTuiScreenStr(b *testing.B) {
	sc := TuiScreen{}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		sc.Str()
	}
}

func TestTuiScreenStr(t *testing.T) {
	sc := TuiScreen{}

	sc.words[0] = Word{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1}
	sc.words[31] = Word{1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}
	sc.words[63] = Word{1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}
	sc.words[95] = Word{1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}
	sc.words[127] = Word{0, 1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}
	sc.words[159] = Word{1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}

	out := sc.Str()
	outRune := []rune(out)

	assert.Equal(t, (NROW/4)*((NCOL/2)+1), len(outRune))
	lines := strings.Split(out, "\n")
	for _, line := range lines {
		if len(line) == 0 {
			continue
		}
		assert.Equal(t, NCOL/2, len([]rune(line)))
	}

	assert.Equal(t, '⠁', ([]rune(lines[0]))[0])
	assert.Equal(t, '⡸', ([]rune(lines[0]))[255])
	assert.Equal(t, '⠈', ([]rune(lines[1]))[255])
}

func TestTuiScreenFetch(t *testing.T) {
	sc := TuiScreen{}

	sc.Fetch(
		Word{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1},
		logic.I,
		[13]logic.Bit{},
	)
	sc.Fetch(
		Word{1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
		logic.I,
		[13]logic.Bit{1, 1, 1, 1, 1},
	)
	sc.Fetch(
		Word{1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
		logic.I,
		[13]logic.Bit{1, 1, 1, 1, 1, 1},
	)
	sc.Fetch(
		Word{1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
		logic.I,
		[13]logic.Bit{1, 1, 1, 1, 1, 0, 1},
	)
	sc.Fetch(
		Word{0, 1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
		logic.I,
		[13]logic.Bit{1, 1, 1, 1, 1, 1, 1},
	)
	sc.Fetch(
		Word{1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
		logic.I,
		[13]logic.Bit{1, 1, 1, 1, 1, 0, 0, 1},
	)

	out := sc.Str()
	lines := strings.Split(out, "\n")
	assert.Equal(t, '⠁', ([]rune(lines[0]))[0])
	assert.Equal(t, '⡸', ([]rune(lines[0]))[255])
	assert.Equal(t, '⠈', ([]rune(lines[1]))[255])
}

func TestTuiScreenCoord2Index(t *testing.T) {
	sc := TuiScreen{}
	var tests = []struct {
		expected           int
		givenRow, givenCol int
	}{
		{0, 0, 0},
		{0, 0, 1},
		{0, 0, 15},
		{1, 0, 16},
		{31, 0, 511},
		{32, 1, 0},
		{8191, 255, 511},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(fmt.Sprintf("%d", tt.expected), func(t *testing.T) {
			actual := sc.coord2wordsIndex(tt.givenRow, tt.givenCol)
			if actual != tt.expected {
				t.Errorf("given(row:%v,col:%v): expected %v, actual %v",
					tt.givenRow, tt.givenCol, tt.expected, actual)
			}
		})
	}
}

func TestTuiScreenAddr2Index(t *testing.T) {
	sc := TuiScreen{}
	var tests = []struct {
		expected int
		given    [13]logic.Bit
	}{
		{0, [13]logic.Bit{0}},
		{1, [13]logic.Bit{1}},
		{2, [13]logic.Bit{0, 1}},
		{4096, [13]logic.Bit{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1}},
		{8191, [13]logic.Bit{1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1}},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(fmt.Sprintf("%d", tt.expected), func(t *testing.T) {
			actual := sc.addr2index(tt.given)
			if actual != tt.expected {
				t.Errorf("given(%v): expected %v, actual %v", tt.given, tt.expected, actual)
			}
		})
	}
}
