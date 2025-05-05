package fileio

import "os"
import "log"

func ReadFile(file string) string {
	text, err := os.ReadFile(file)

	if err != nil {
		log.Fatal(err)
	}

	return string(text)
}
