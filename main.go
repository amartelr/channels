package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	links := []string{
		"http://www.google.com",
		"http://facebook.com",
		"http://stackoverflow.com",
		"http://golang.org",
		"http://amazon.com",
	}
	// Defining sring channel
	c := make(chan string)

	for _, link := range links {
		//fmt.Println(link, "checking")
		go checkLink(link, c)
	}

	for l := range c {
		go func(link string) {
			time.Sleep(3 * time.Second)
			checkLink(link, c)
		}(l)
	}
}

func checkLink(link string, c chan string) {
	//time.Sleep(3 * time.Second)
	_, err := http.Get(link)
	if err != nil {
		fmt.Println(link, "is down")
		c <- link
		return
	}

	fmt.Println(link, "is up")
	c <- link
}
