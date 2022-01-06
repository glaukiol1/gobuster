package main

import (
	"example.com/hello/helpers"
)

func main() {
	helpers.StartThreads(10, "http://httpbin.org/get")
	helpers.GetWordlist("go.mod").Text()
}
