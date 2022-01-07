package helpers

import (
	"bufio"
	"os"
)

func LineCounter(path string) int {
	file, _ := os.Open(path)
	defer file.Close()
	fileScanner := bufio.NewScanner(file)
	lineCount := 0
	for fileScanner.Scan() {
		lineCount++
	}
	return lineCount
}

func chunkSlice(slice []string, chunkSize int) [][]string {
	var chunks [][]string
	for i := 0; i < len(slice); i += chunkSize {
		end := i + chunkSize

		// necessary check to avoid slicing beyond
		// slice capacity
		if end > len(slice) {
			end = len(slice)
		}

		chunks = append(chunks, slice[i:end])
	}

	return chunks
}

func SplitWordlist(wordlist bufio.Scanner, threads int, max int) [][]string {
	return chunkSlice(getWordlistArray(wordlist)[:max], max/threads)
}

func getWordlistArray(wordlist bufio.Scanner) []string {
	var words = make([]string, 0)
	for i := 0; wordlist.Scan(); i++ {
		words = append(words, wordlist.Text())
	}
	return words
}
