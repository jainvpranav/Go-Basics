package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	// startTime := time.Now().Unix()
	links := []string{
		"https://google.com",
		"https://facebook.com",
		"https://stackoverflow.com",
		"https://amazon.com",
		"https://golang.org",
	}

	c := make(chan string)

	for _, link := range links {
		go makeRequest(link, c)
	}

	for l := range c {
		go func(l string) {
			time.Sleep(time.Second * 5)
			makeRequest(l, c)
		}(l)
	}

	// Got a diff of 3s when using routines
	// endTime := time.Now().Unix()
	// fmt.Println("Time Start: ", startTime)
	// fmt.Println("Time End: ", endTime)
	// fmt.Println("Time Lapsed: ", endTime-startTime)

}

func makeRequest(link string, c chan string) {
	_, err := http.Get(link)
	if err != nil {
		fmt.Println(link, "might be down!")
		c <- link
		return
	}
	fmt.Println(link, "is up!")
	c <- link
}
