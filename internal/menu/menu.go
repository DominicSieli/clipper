package menu

import "os"
import "fmt"
import "clipper/internal/input"
import "clipper/internal/fileio"
import "clipper/internal/actions"
import "clipper/internal/terminal"
import "clipper/internal/utilities"

func Menu() string {
	text := ""
	index := 0
	file := fileio.ReadFile("/home/dominic/.cache/xfce4/clipman/textsrc")
	entries := utilities.GetEntries(file)

	if len(entries) == 0 {
		terminal.Clear()
		terminal.ColorPrintLine("red", "[NO ENTRIES FOUND]")
		fmt.Print(terminal.UNHIDE_CURSOR)
		os.Exit(0)
	}

	for true {
		render(index, entries)
		key := input.Key()
		index = scroll(key, index, len(entries))

		if actions.Escape(key) {
			terminal.Clear()
			fmt.Print(terminal.UNHIDE_CURSOR)
			os.Exit(0)
		}

		if actions.Enter(key) {
			entry := entries[index]
			utilities.SelectEntry(entry)
			terminal.Clear()
			os.Exit(0)
		}
	}

	return text
}

func scroll(key byte, index int, size int) int {
	if actions.Up(key) && index > 0 {
		index--
	}

	if actions.Down(key) && index + 1 < size {
		index++
	}

	return index
}

func render(index int, entries []string) {
	end := 0
	size := len(entries)

	terminal.Clear()

	_, height, err := terminal.TerminalSize()

	if err != nil {
		panic(err)
	}

	if index + height <= size {
		end = index + height - 2
	}

	if index + height > size {
		end = size
	}

	for i := index; i < end; i++ {
		if i != index {
			terminal.ColorPrintLine("grey", entries[i])
		}

		if i == index {
			terminal.ColorPrintLine("green", entries[i])
		}
	}

	fmt.Print(terminal.HIDE_CURSOR)
}
