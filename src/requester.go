package src

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

func Request(url string, id int, wordlist []string) [4]int {
	var startTime = time.Now().Unix()
	log.Println("Starting...")
	var errors int = 0
	var successes int = 0
	var hardErrors int = 0

	println(len(wordlist))

	for i := 0; i < len(wordlist); i++ {
		resp, err := http.Get(url + wordlist[i])
		if resp.StatusCode != 200 {
			errors = errors + 1
		}
		if err != nil {
			hardErrors = hardErrors + 1
		}
		if err == nil && resp.StatusCode == 200 {
			successes = successes + 1
			println("/"+wordlist[i], resp.StatusCode, "ID: "+fmt.Sprint(id))
		}
	}
	var a [4]int
	a[0] = successes
	a[1] = errors
	a[2] = hardErrors
	a[3] = int(time.Now().Unix() - startTime)
	return a
}
