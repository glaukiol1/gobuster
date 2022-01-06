package helpers

import (
	"bufio"
	"log"
	"os"
)

func GetWordlist(path string) *bufio.Scanner {

	f, err := os.Open(path)

	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	return bufio.NewScanner(f)

	// if err := scanner.Err(); err != nil {
	// 	log.Fatal(err)
	// }
}
