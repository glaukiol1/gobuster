package src

import (
	"fmt"
	"log"
	"strings"
)

var width int = 60

var rstring = strings.Repeat("=", width)

func PrintSummary(successes, errors, hardErrors, avarageDoneTime int) {
	println(rstring + "\n" + "Successful Requests: " + fmt.Sprint(successes) + "\n" + "Errors (not matching status code): " + fmt.Sprint(errors) + "\n" + "Hard errors (failed requests): " + fmt.Sprint(hardErrors) + "\n" + "Avarage thread done time: " + fmt.Sprint(avarageDoneTime) + "seconds" + "\n" + rstring)
}

func PrintIntro(mode, url, wordlist, success_status, threads string) {
	println(rstring + "\ngobuster\nby Glaukio L (@glaukiol1) | Inspired by gobuster\n" + rstring + "\n[+] Mode         : " + mode + "\n[+] Url/Domain   : " + url + "\n[+] Wordlist     : " + wordlist + "\n[+] Status codes : " + success_status + "\n[+] Threads:     : " + threads + "\n" + rstring)
	log.Println("Starting gobuster\n" + rstring)
}
