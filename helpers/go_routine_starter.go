package helpers

import (
	"fmt"

	"example.com/hello/src"
)

func StartThreads(threads int, url string) {
	c := make(chan string)

	for i := 0; i < threads; i++ {
		go func() {
			c <- src.Request(10, "https://jsonplaceholder.typicode.com/posts", 0)
		}()
	}

	for i := 0; i < 2; i++ {
		// Await both of these values
		// simultaneously, printing each one as it arrives.
		select {
		case msg1 := <-c:
			fmt.Println(msg1)
		}
	}
}
