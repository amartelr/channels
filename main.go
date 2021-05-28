package main

import (
	"fmt"
	"net/http"
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

	/* 	fmt.Println(<-c)
	   	fmt.Println(<-c)
	   	fmt.Println(<-c)
	   	fmt.Println(<-c)
	   	fmt.Println(<-c)
	*/

	for i := 0; i < len(links); i++ {
		fmt.Println(<-c)
	}
}

func checkLink(link string, c chan string) {
	_, err := http.Get(link)
	if err != nil {
		fmt.Println(link, "is down")
		c <- "is down i think"
		return
	}

	fmt.Println(link, "is up")
	c <- "Yep, it's up"
}
