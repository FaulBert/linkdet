package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {
	shortened_url := []string{}

	if len(os.Args) >= 2 {
		shortened_url = os.Args[1:]
	} else {
		fmt.Println("Where any url(s)?")
	}

	for _, short_url := range shortened_url {
		url, err := detective(short_url)
		if err != nil {
			fmt.Println("Error: ", err)
			continue
		}

		fmt.Printf("%s goes to %s  \n", short_url, url)
	}
}

func detective(short_url string) (string, error) {
	workers := 10

	urls := make(chan string)
	results := make(chan string)

	for i := 0; i < workers; i++ {
		go worker(urls, results)
	}

	go func() {
		urls <- short_url
	}()

	return <-results, nil
}

func worker(urls, results chan string) {
	for {
		url := <-urls

		resp, err := http.Get(url)
		if err != nil {
			results <- err.Error()
			continue
		}

		redirected := resp.Request.URL.String()

		if redirected == url {
			results <- redirected
		} else {
			urls <- redirected
		}

		resp.Body.Close()
		return
	}
}
