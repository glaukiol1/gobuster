package main

import (
	"example.com/hello/helpers"
)

func main() {
	var wordlist = helpers.GetWordlist("/users/glaukiollupo/Projects/pybuster/wordlist-dir.txt")
	helpers.StartThreads(20, "http://sans.org/", helpers.SplitWordlist(wordlist, 21, 90000))
}
