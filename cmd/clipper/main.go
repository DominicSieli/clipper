package main

import "clipper/internal/menu"
import "clipper/internal/terminal"

func main() {
	terminal.Clear()

	for true {
		menu.Menu()
	}

	terminal.Clear()
}
