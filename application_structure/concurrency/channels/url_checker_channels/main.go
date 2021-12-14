package main

import (
	"fmt"
	"net/http"
	"runtime"
)

func checkURL(url string, c chan string) {

	resp, err := http.Get(url)

	if err != nil {
		//fmt.Println(err)
		s := fmt.Sprintf("%s is DOWN!\n", url)
		s += fmt.Sprintf("Error: %v\n", err)
		c <- url //sending into the channel
	} else {
		defer resp.Body.Close()
		s := fmt.Sprintf("%s -> Status Code: %d \n", url, resp.StatusCode)

		s += fmt.Sprintf("%s is UP\n", url)
		c <- s

	}
}

func main() {
	urls := []string{"https://golang.org", "https://www.google.com", "https://www.medium.com"}

	//1. declare new channel using make fn
	c := make(chan string)

	// var wg sync.WaitGroup
	// wg.Add(len(urls))

	for _, url := range urls {
		go checkURL(url, c)

		//fmt.Println(strings.Repeat("#", 20))
	}
	fmt.Println("No. of goroutines running are : ", runtime.NumGoroutine())
	// for i := 0; i < len(urls); i++ {
	// 	fmt.Println(<-c)
	// }

	// for {
	// 	go checkURL(<-c, c)
	// 	fmt.Println(strings.Repeat("#", 30))
	// 	time.Sleep(time.Second * 2)
	// }

	// for url := range c {
	// 	time.Sleep(time.Second * 2)
	// 	go checkURL(url, c)
	// }

	for url := range c {
		go func(u string) {
			checkURL(u, c)
		}(url)
	}

}
