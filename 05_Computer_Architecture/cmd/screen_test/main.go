package main

import (
	"fmt"

	logic "github.com/kazufusa/nand2tetris/01_Boolean_Logic"
	computer "github.com/kazufusa/nand2tetris/05_Computer_Architecture"
	"github.com/rivo/tview"
)

type Word = [16]logic.Bit

func main() {
	app := tview.NewApplication()
	textView := tview.NewTextView().
		SetChangedFunc(func() { app.Draw() })

	sc := computer.TuiScreen{}

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

	fmt.Fprintf(textView, sc.Str())
	textView.Clear()
	fmt.Fprintf(textView, sc.Str())

	if err := app.SetRoot(textView, true).EnableMouse(true).Run(); err != nil {
		panic(err)
	}
}
