package main

import (
	"fmt"
	"net/http"
)

func main() {
	links := []string{
		"https://www.google.com",
		"https://www.amazon.com",
		"https://www.apple.com",
		"https://www.flipkart.com",
		"https://www.cisco.com",
	}

	c := make(chan string)

	for _, link := range links {
		go checkLink(link, c)
	}

	for l := range c {
		go checkLink(l, c)
	}

	fmt.Println(<-c)
}

func checkLink(link string, c chan string) {
	_, err := http.Get(link)
	if err != nil {
		fmt.Println(link, "is down")
		c <- "Might be down"
		return
	}
	fmt.Println(link, "is running ")
	c <- "the link is running"
}
