package assembler

import (
	"errors"
)

var (
	VARIABLE_SYMBOLS_STARTING_ADDR = 16
)

type SymbolTable struct {
	table           map[string]int
	variableCounter int
}

func NewSymbolTable() *SymbolTable {
	table := map[string]int{
		"R0":     0,
		"R1":     1,
		"R2":     2,
		"R3":     3,
		"R4":     4,
		"R5":     5,
		"R6":     6,
		"R7":     7,
		"R8":     8,
		"R9":     9,
		"R10":    10,
		"R11":    11,
		"R12":    12,
		"R13":    13,
		"R14":    14,
		"R15":    15,
		"SP":     0,
		"LCL":    1,
		"ARG":    2,
		"THIS":   3,
		"THAT":   4,
		"SCREEN": 16384,
		"KBD":    24576,
	}
	return &SymbolTable{table: table, variableCounter: VARIABLE_SYMBOLS_STARTING_ADDR}
}

func (st *SymbolTable) AddEntry(s string, lineNo int, isLabel bool) {
	if isLabel {
		st.table[s] = lineNo
	} else {
		if _, ok := st.table[s]; ok {
			return
		}
		st.table[s] = st.variableCounter
		st.variableCounter += 1
	}
	return
}

func (st *SymbolTable) Contains(s string) bool {
	_, ok := st.table[s]
	return ok
}

func (st *SymbolTable) GetAddress(s string) (int, error) {
	if ret, ok := st.table[s]; ok {
		return ret, nil
	}
	return 0, errors.New("input entry not found")
}
