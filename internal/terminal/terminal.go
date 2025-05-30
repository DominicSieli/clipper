package terminal

import "os"
import "fmt"
import "os/exec"
import "golang.org/x/term"

const (
	RESET = "\033[0m"
	HIDE_CURSOR = "\033[?25l"
	UNHIDE_CURSOR = "\033[?25h"
	RED = "\033[38;2;255;0;0m"
	GREEN = "\033[38;2;0;255;0m"
	CYAN = "\033[38;2;0;255;255m"
	GREY = "\033[38;2;128;128;128m"
)

func Clear() {
	cmd := exec.Command("clear")
	cmd.Stdout = os.Stdout
	cmd.Run()
}

func ColorPrintLine(color string, text string) {
	switch color {
	case "red":
		color = RED
	case "green":
		color = GREEN
	case "cyan":
		color = CYAN
	case "grey":
		color = GREY
	default:
		color = ""
	}

	fmt.Println(color + text + RESET)
}

func TerminalSize() (int, int, error) {
	fd := int(os.Stdout.Fd())

	width, height, err := term.GetSize(fd)

	if err != nil {
		return 0, 0, err
	}

	return width, height, nil
}
