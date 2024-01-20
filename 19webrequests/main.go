package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

const url = "https://lco.dev"

func main() {
	fmt.Println("Welcome to web request")
	response, err := http.Get(url)

	if err != nil {
		panic(err)
	}

	fmt.Printf("Response type %T\n", response)
	defer response.Body.Close()

	databyte, err := ioutil.ReadAll(response.Body)
	if err != nil {
		panic(err)
	}
	fmt.Println("Response is - ", string(databyte))
}
