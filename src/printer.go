package src

import (
	"fmt"
	"strings"
)

var width int = 60

var rstring = strings.Repeat("=", width)

func PrintSummary(successes, errors, hardErrors, avarageDoneTime int) {
	println(rstring + "\n" + "Successful Requests: " + fmt.Sprint(successes) + "\n" + "Errors (not matching status code): " + fmt.Sprint(errors) + "\n" + "Hard errors (failed requests): " + fmt.Sprint(hardErrors) + "\n" + "Avarage thread done time: " + fmt.Sprint(avarageDoneTime) + "seconds" + "\n" + rstring)
}

func PrintIntro() {
	println(rstring + "\ngobuster\nby Glaukio L (@glaukiol1) | Inspired by gobuster\n" + rstring + "\n[+] Mode         : {mode}\n[+] Url/Domain   : {url}\n[+] Wordlist     : {w}\n[+] Status codes : {s}\n[+] Threads:     : {threads}\n" + rstring + "\n{out_time} Starting pybuster" + "\n" + rstring)
}
