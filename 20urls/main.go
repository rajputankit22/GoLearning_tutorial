package main

import (
	"fmt"
	"net/url"
)

const myurl string = "https://lco.dev:3000/learn?coursename=reactjs&paymentid=10"

func main() {
	fmt.Println("Welcome to handle url in golang")
	fmt.Println(myurl)

	// parsing
	result, _ := url.Parse(myurl)
	fmt.Println(result.Scheme)
	fmt.Println(result.Host)
	fmt.Println(result.Path)
	fmt.Println(result.Port())
	fmt.Println(result.RawQuery)

	qparams := result.Query()
	fmt.Printf("The type of query params are: %T\n", qparams)

	fmt.Println(qparams["coursename"])

	partsOfUrl := &url.URL{
		Scheme:  "https",
		Host:    "lco.dev:3000",
		Path:    "/learn",
		RawPath: "coursename=reactjs",
	}

	anotherUrl := partsOfUrl.String()
	fmt.Println(anotherUrl)
}
