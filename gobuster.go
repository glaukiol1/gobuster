package main

import (
	"example.com/hello/helpers"
)

func main() {
	helpers.StartThreads(2, "http://httpbin.org/get")
}
