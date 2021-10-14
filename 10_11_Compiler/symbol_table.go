package compiler

import "errors"

type Kind int

const (
	KdStatic Kind = iota
	KdField
	KdArg
	KdVar
	KdNone
)

var (
	ErrSymbolAlreadyExists = errors.New("the symbol already exists")
	ErrSymbolNotFound      = errors.New("the symbol not found")
)

type Symbol struct {
	name   string
	sType  string
	kind   Kind
	number int
}

type SymbolTable struct {
	Symbols []Symbol
}

func NewSymbolTable() *SymbolTable {
	return &SymbolTable{}
}

func (st *SymbolTable) Reset() {
	st.Symbols = []Symbol{}
}

func (st *SymbolTable) Define(name, sType string, kind Kind) error {
	_, err := st.IndexOf(name)
	if err == nil {
		return ErrSymbolAlreadyExists
	}
	st.Symbols = append(st.Symbols, Symbol{name: name, sType: sType, kind: kind, number: st.VarCount(kind) + 1})
	return nil
}

func (st *SymbolTable) VarCount(kind Kind) int {
	ret := 0
	for _, s := range st.Symbols {
		if s.kind == kind {
			ret++
		}
	}
	return ret
}

func (st *SymbolTable) KindOf(name string) (Kind, error) {
	i, err := st.IndexOf(name)
	if err != nil {
		return KdNone, err
	}
	return st.Symbols[i].kind, nil
}

func (st *SymbolTable) TypeOf(name string) (string, error) {
	i, err := st.IndexOf(name)
	if err != nil {
		return "", err
	}
	return st.Symbols[i].sType, nil
}

func (st *SymbolTable) IndexOf(name string) (int, error) {
	for i, s := range st.Symbols {
		if s.name == name {
			return i, nil
		}
	}
	return -1, ErrSymbolNotFound
}

func (st *SymbolTable) Get(name string) *Symbol {
	i, err := st.IndexOf(name)
	if err != nil {
		return nil
	} else {
		return &st.Symbols[i]
	}
}
