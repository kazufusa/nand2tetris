package main

import (
	"io/ioutil"
	"log"
	"os"
	"path"

	assembler "github.com/kazufusa/nand2tetris/06_Assembler"
)

func main() {
	in := os.Args[1]
	asm, err := assembler.NewAssembler(in)
	if err != nil {
		log.Fatal(err)
	}
	ret, err := asm.Assemble()
	if err != nil {
		log.Fatal(err)
	}
	ext := path.Ext(in)
	out := os.Args[1][0:len(in)-len(ext)] + ".hack"
	err = ioutil.WriteFile(out, []byte(ret), 0644)
	if err != nil {
		log.Fatal(err)
	}
}
