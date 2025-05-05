package utilities

import "io"
import "os"
import "os/exec"
import "strings"

func GetEntries(str string) []string {
	start := 0
	entries := []string{}

	for i, v := range str {
		if v == rune('=') && i + 1 < len(str) {
			start = i + 1
		}
	}

	for i := start; i < len(str); i++ {
		for j := i; j < len(str) - 1; j++ {
			if rune(str[j]) == rune(';') && rune(str[j-1]) != rune('\\') && rune(str[j+1]) != rune('\\') {
				entries = append(entries, strings.ReplaceAll(str[i:j], "\\;", ";"))
				i = j+1
				continue
			}
		}
	}

	return entries
}

func SelectEntry(entry string) {
	echo := exec.Command("echo", entry)
	xClip := exec.Command("xclip", "-selection", "clipboard")

	pipeReader, pipeWriter := io.Pipe()
	echo.Stdout = pipeWriter
	xClip.Stdin = pipeReader

	xClip.Stdout = os.Stdout

	err1 := echo.Start()
	if err1 != nil {
		panic(err1)
	}

	err2 := xClip.Start()
	if err2 != nil {
		panic(err2)
	}

	go func() {
		echo.Wait()
		pipeWriter.Close()
	}()

	xClip.Wait()
}
