package src

import (
	"fmt"
	"log"
	"net/http"
	"strings"
	"time"
)

func Request(total *int, threads int, wordlist_total int, url string, id int, wordlist []string) [4]int {
	var startTime = time.Now().Unix()
	log.Println("Starting...")
	var errors int = 0
	var successes int = 0
	var hardErrors int = 0

	println(len(wordlist))

	for i := 0; i < len(wordlist); i++ {
		resp, err := http.Get(url + wordlist[i])
		*total = *total + 1
		if resp.StatusCode != 200 && err == nil {
			errors = errors + 1
		}
		if err != nil {
			hardErrors = hardErrors + 1
		}
		if err == nil && resp.StatusCode == 200 {
			successes = successes + 1
			var spaces int = 50

			println("/" + wordlist[i] + strings.Repeat(" ", spaces-len("/"+wordlist[i])) + "  Status: " + fmt.Sprint(resp.StatusCode) + " (Thread: " + fmt.Sprint(id) + ")  " + fmt.Sprint(*total) + "/" + fmt.Sprint(len(wordlist)*threads))
		}
	}
	var a [4]int
	a[0] = successes
	a[1] = errors
	a[2] = hardErrors
	a[3] = int(time.Now().Unix() - startTime)
	return a
}
