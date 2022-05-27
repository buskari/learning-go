package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	urls := []string{
		"http://tibia.com",
		"http://google.com",
		"http://stackoverflow.com",
		"http://golang.com",
		"http://amazon.com",
	}

	c := make(chan string)

	for _, url := range urls {
		go checkLink(url, c)
	}

	// for {
	// 	go checkLink(<-c, c)
	// }

	// its the same thing of the for commented above
	for u := range c {
		go func(u string) {
			time.Sleep(5 * time.Second)
			checkLink(u, c)
		}(u)
	}
}

func checkLink(link string, c chan string) {
	_, err := http.Get(link)

	if err != nil {
		c <- link
		return
	}

	fmt.Println(link, "	success")
	c <- link
}
