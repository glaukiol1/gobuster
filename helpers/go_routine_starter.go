package helpers

import (
	"example.com/hello/src"
)

func StartThreads(threads int, url string, splitWordlist [][]string) {
	c := make(chan [4]int)

	for i := 0; i < threads; i++ {
		var g = splitWordlist[i]
		var id = i
		go func() {
			c <- src.Request(url, id, g)
		}()
	}
	var successes int = 0
	var errors int = 0
	var hardErrors = 0
	var avarageDoneTime = 0
	for i := 0; i < threads; i++ {
		// Await both of these values
		// simultaneously, printing each one as it arrives.
		select {
		case arry := <-c:
			for i := 0; i < 4; i++ {
				switch i {
				case 0:
					successes += arry[i]
				case 1:
					errors += arry[i]
				case 2:
					hardErrors += arry[i]
				case 3:
					avarageDoneTime += arry[i]
				}
			}
		}
	}
	avarageDoneTime = avarageDoneTime / threads
	println(successes, errors, hardErrors, avarageDoneTime)
}
