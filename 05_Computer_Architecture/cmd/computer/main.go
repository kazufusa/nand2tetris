package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/gdamore/tcell/v2"
	logic "github.com/kazufusa/nand2tetris/01_Boolean_Logic"
	memory "github.com/kazufusa/nand2tetris/03_Memory"
	computer "github.com/kazufusa/nand2tetris/05_Computer_Architecture"
	"github.com/rivo/tview"
)

type Word = computer.Word

func main() {
	sc := computer.TuiScreen{}
	clock := memory.Clock(0)

	kb := computer.TuiKeyboard{}
	ram := computer.NewMemory(&clock, &sc, &kb)
	ram.Fetch(Word{0, 0, 0, 1}, logic.I, [15]logic.Bit{})
	clock.Progress()

	rom := computer.NewROM32K()
	if err := rom.LoadHackFile(os.Args[1]); err != nil {
		log.Fatal(err)
	}

	cpu := computer.NewCPU()

	com := computer.NewComputer(&cpu, &ram, &rom, &clock)

	app := tview.NewApplication()
	textView := tview.NewTextView().
		SetChangedFunc(func() { app.Draw() })
	textView.SetInputCapture(func(key *tcell.EventKey) *tcell.EventKey {
		kb.Set(key)
		return key
	})

	go func() {
		for {
			com.FetchAndExecute(logic.O)
		}
	}()

	go func() {
		for {
			textView.Clear()
			fmt.Fprintf(textView, sc.Str())
			time.Sleep(50 * time.Millisecond)
		}
	}()

	if err := app.SetRoot(textView, true).EnableMouse(true).Run(); err != nil {
		log.Fatal(err)
	}
}
