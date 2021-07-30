package computer

import (
	"fmt"
	"testing"

	"github.com/gdamore/tcell/v2"
)

func TestTuiKeyboardSet(t *testing.T) {
	kb := NewTuiKeyboard()
	tcell.NewEventKey(tcell.KeyRune, ' ', 0)

	var tests = []struct {
		expected Word
		given    *tcell.EventKey
	}{
		{Word{0, 0, 0, 0, 0, 1, 0, 0}, tcell.NewEventKey(tcell.KeyRune, ' ', 0)},
		{Word{1, 0, 0, 0, 0, 0, 1, 0}, tcell.NewEventKey(tcell.KeyRune, 'A', 0)},
		{Word{1, 0, 0, 0, 0, 1, 1, 0}, tcell.NewEventKey(tcell.KeyRune, 'a', 0)},
		{Word{0, 1, 1, 1, 1, 1, 1, 0}, tcell.NewEventKey(tcell.KeyRune, '~', 0)},
		{Word{0, 0, 0, 0, 0, 0, 0, 1}, tcell.NewEventKey(tcell.KeyEnter, rune(13), 0)},
		{Word{1, 0, 0, 0, 0, 0, 0, 1}, tcell.NewEventKey(tcell.KeyBackspace2, rune(127), 0)},
		{Word{0, 0, 1, 1, 0, 0, 0, 1}, tcell.NewEventKey(tcell.KeyEsc, rune(0), 0)},
		{Word{1, 0, 1, 1, 0, 0, 0, 1}, tcell.NewEventKey(tcell.KeyF1, rune(0), 0)},
		{Word{0, 0, 0, 1, 1, 0, 0, 1}, tcell.NewEventKey(tcell.KeyF12, rune(0), 0)},

		{Word{}, tcell.NewEventKey(tcell.Key(0), rune(0), 0)},
		{Word{}, tcell.NewEventKey(tcell.Key(255), rune(0), 0)},
		{Word{}, tcell.NewEventKey(tcell.Key(0), rune(255), 0)},
		{Word{}, tcell.NewEventKey(tcell.Key(255), rune(255), 0)},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(fmt.Sprintf("%v,%v", tt.given.Rune(), tt.given.Key()), func(t *testing.T) {
			kb.Set(tt.given)
			actual := kb.Fetch()
			if actual != tt.expected {
				t.Errorf("given(%v): expected %v, actual %v", tt.given, tt.expected, actual)
			}
		})
	}
}
