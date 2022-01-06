package helpers

import "bufio"

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

func SplitWordlist(wordlist bufio.Scanner, threads int) [][]string {
	var len int = 23
	return chunkSlice(getWordlistArray(wordlist), len/threads)
}

func getWordlistArray(wordlist bufio.Scanner) []string {
	var words = make([]string, 0)
	for i := 0; wordlist.Scan(); i++ {
		words = append(words, wordlist.Text())
	}
	return words
}
