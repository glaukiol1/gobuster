package src

import (
	"log"
	"net/http"
	"time"
)

func Request(amount int, url string, id int) [4]int {
	var startTime = time.Now().Unix()
	log.Println("Starting...")
	var errors int = 0
	var successes int = 0
	var hardErrors int = 0
	for i := 0; i < amount; i++ {
		resp, err := http.Get(url)
		if resp.StatusCode != 200 {
			errors = errors + 1
		}

		if err != nil {
			hardErrors = hardErrors + 1
		}
		successes = successes + 1
	}
	var a [4]int
	a[0] = successes
	a[1] = errors
	a[2] = hardErrors
	a[3] = int(time.Now().Unix() - startTime)
	return a
}
