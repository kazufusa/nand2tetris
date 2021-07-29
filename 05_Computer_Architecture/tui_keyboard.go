package computer

import (
	"sync"

	"github.com/gdamore/tcell/v2"
)

var (
	_      IKeyboard = (*TuiKeyboard)(nil)
	keyMap           = map[tcell.Key]uint8{
		// tcell.KeyDEL:   127,
		tcell.KeyEnter:      128,
		tcell.KeyBackspace2: 129,
		tcell.KeyLeft:       130,
		tcell.KeyUp:         131,
		tcell.KeyRight:      132,
		tcell.KeyDown:       133,
		tcell.KeyHome:       134,
		tcell.KeyEnd:        135,
		tcell.KeyPgUp:       136,
		tcell.KeyPgDn:       137,
		tcell.KeyInsert:     138,
		tcell.KeyDelete:     139,
		tcell.KeyEsc:        140,
		tcell.KeyF1:         141,
		tcell.KeyF2:         142,
		tcell.KeyF3:         143,
		tcell.KeyF4:         144,
		tcell.KeyF5:         145,
		tcell.KeyF6:         146,
		tcell.KeyF7:         147,
		tcell.KeyF8:         148,
		tcell.KeyF9:         149,
		tcell.KeyF10:        150,
		tcell.KeyF11:        151,
		tcell.KeyF12:        152,
	}
)

type TuiKeyboard struct {
	mu   sync.RWMutex
	word Word
}

func NewTuiKeyboard() TuiKeyboard {
	return TuiKeyboard{}
}

func (kb *TuiKeyboard) Fetch() Word {
	kb.mu.RLock()
	defer kb.mu.RUnlock()
	return kb.word
}

func (kb *TuiKeyboard) Set(key *tcell.EventKey) {
	kb.mu.Lock()
	defer kb.mu.Unlock()
	code := uint8(key.Rune())
	if code < 32 || 126 < code {
		code = keyMap[key.Key()]
	}

	kb.word = Word{
		code & 1,
		code & 2 >> 1,
		code & 4 >> 2,
		code & 8 >> 3,
		code & 16 >> 4,
		code & 32 >> 5,
		code & 64 >> 6,
		code & 128 >> 7,
	}
}
