package main

import (
	"example.com/hello/helpers"
)

func main() {
	var wordlist_path string
	var success_statuses string
	var url string
	var mode string
	var threads int

	helpers.CommandLineArgs(&wordlist_path, &success_statuses, &url, &mode, &threads)
	var wordlist = helpers.GetWordlist(wordlist_path)
	var lines = 1000 //helpers.LineCounter((wordlist_path))[0]
	helpers.StartThreads(mode, threads, url, helpers.SplitWordlist(wordlist, threads, lines), lines, wordlist_path, success_statuses, helpers.LineCounter(wordlist_path)[1])
}
