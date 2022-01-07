package src

import (
	"fmt"
	"net/http"
	"strings"
	"time"
)

func Request(total *int, threads int, wordlist_total int, url string, id int, wordlist []string, success_string string) [4]int {
	var startTime = time.Now().Unix()
	var errors int = 0
	var successes int = 0
	var hardErrors int = 0

	for i := 0; i < len(wordlist); i++ {
		resp, err := http.Get(url + wordlist[i])
		*total = *total + 1
		if !strings.Contains(success_string, fmt.Sprint(resp.StatusCode)) && err == nil {
			errors = errors + 1
		}
		if err != nil {
			hardErrors = hardErrors + 1
		}
		if err == nil && strings.Contains(success_string, fmt.Sprint(resp.StatusCode)) {
			successes = successes + 1
			var spaces int = 20

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
