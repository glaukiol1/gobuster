package main

import (
	"example.com/hello/helpers"
)

func main() {
	var wordlist = helpers.GetWordlist("/users/glaukiollupo/Projects/pybuster/wordlist.txt")
	helpers.StartThreads(10, "http://httpbin.org/get", helpers.SplitWordlist(wordlist, 2))
}
