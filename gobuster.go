package main

import (
	"example.com/hello/helpers"
)

func main() {
	var wordlist_path string = "/users/glaukiollupo/Projects/pybuster/wordlist-dir.txt"
	var wordlist = helpers.GetWordlist(wordlist_path)
	var success_statuses = "200,301,302,303,304"
	var mode string = "dir"
	var lines = 2000 //helpers.LineCounter((wordlist_path))
	var threads int = 20
	helpers.StartThreads(mode, threads, "http://sans.org/", helpers.SplitWordlist(wordlist, threads+1, lines), lines, wordlist_path, success_statuses)
}
