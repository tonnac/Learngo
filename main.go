package main

import (
	"errors"
	"fmt"
	"net/http"
)

type requestResult struct {
	url    string
	status string
}

func (r requestResult) String() string {
	return fmt.Sprintf("%s is %s", r.url, r.status)
}

var errRequestFailed = errors.New("Request Failed")

func main() {
	urls := []string{
		"https://www.airbnb.com/",
		"https://www.google.com/",
		"https://www.amazon.com/",
		"https://www.reddit.com/",
		"https://soundcloud.com/",
		"https://www.facebook.com/",
		"https://www.instagram.com/",
		"https://academy.nomadcoders.co/",
	}

	var c = make(chan requestResult)

	for _, url := range urls {
		go hitURL(url, c)
	}
	// for i := 0; i < len(urls); i++ {
	// 	fmt.Println(<-c)
	// }

	var results = map[string]string{}
	for i := 0; i < len(urls); i++ {
		result := <-c
		results[result.url] = result.status
	}

	for url, status := range results {
		fmt.Println(url, status)
	}
}

func hitURL(url string, c chan<- requestResult) {
	resp, err := http.Get(url)
	status := "OK"
	if err != nil || resp.StatusCode >= 400 {
		status = fmt.Sprintf("FAILED %d", resp.StatusCode)
	}
	c <- requestResult{url: url, status: status}
}
