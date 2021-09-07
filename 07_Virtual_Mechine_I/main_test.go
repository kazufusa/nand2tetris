package vmI

import (
	"fmt"
	"testing"
)

func Test(t *testing.T) {
	parser, _ := NewParser(`// Pushes and adds two constants
  push constant 7
  push constant 8
  add
  //`)
	fmt.Printf("%#v\n", parser)
}
