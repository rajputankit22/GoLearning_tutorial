package main

import (
	"fmt"
	"net/http"
	"sync"
)

// func main() {
// 	go greeter("Hello")
// 	greeter("World")

// }

// func greeter(s string) {
// 	for i := 0; i < 6; i++ {
// 		time.Sleep(3 * time.Millisecond)
// 		fmt.Println(s)
// 	}
// }

var wg sync.WaitGroup

func main() {

	websiteList := []string{
		"https://lco.dev",
		"https://fb.com",
		"https://google.com",
		"https://github.dev",
	}

	for _, web := range websiteList {
		go getStatusCode(web)
		wg.Add(1)
	}

	wg.Wait()

}

func getStatusCode(endpoint string) {
	defer wg.Done()
	res, err := http.Get(endpoint)

	if err != nil {
		fmt.Println("OOPS endpoint")
	} else {
		fmt.Printf("%d is the status code for %s\n", res.StatusCode, endpoint)
	}
}
