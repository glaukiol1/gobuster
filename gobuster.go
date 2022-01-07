package main

import (
	"example.com/hello/helpers"
)

func main() {
	var wordlist_path string = "/users/glaukiollupo/Projects/pybuster/wordlist-dir.txt"
	var wordlist = helpers.GetWordlist(wordlist_path)
	var success_statuses = "200"
	var url = "https://sans.org/"
	var mode string = "dir"
	var lines = 2000 //helpers.LineCounter((wordlist_path))[0] // doing this to limit the lines, in real cases use the
	var threads int = 20
	helpers.StartThreads(mode, threads, url, helpers.SplitWordlist(wordlist, threads, lines), lines, wordlist_path, success_statuses, helpers.LineCounter(wordlist_path)[1])
}
