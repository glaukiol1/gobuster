package src

import (
	"log"
	"net/http"
	"time"
)

func Request(amount int, url string, id int) string {
	var startTime = time.Now().Unix()
	log.Println("Starting...")
	for i := 0; i < amount; i++ {
		resp, err := http.Get(url)
		if err != nil {
			log.Fatalln(err)
		}
		log.Println(resp.StatusCode, " THREAD: ", id)
	}
	log.Println(time.Now().Unix() - startTime)
	println("DONE!")
	return "GOOD JOB!"
}
