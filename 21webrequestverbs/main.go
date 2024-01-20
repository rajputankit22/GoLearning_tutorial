package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

func main() {
	fmt.Println("Welcome to web request in golang")
	// PerformGetRequest()
	// PerformPostJsonRequest()
	PerformPostFormRequest()
}

func PerformGetRequest() {
	const myurl = "https://webhook.site/3041c1ac-ed8f-4c57-af57-a5c9124b86fe"
	response, err := http.Get(myurl)
	if err != nil {
		panic(err)
	}

	defer response.Body.Close()

	fmt.Println("Status code: ", response.StatusCode)
	fmt.Println("Content lenght code: ", response.ContentLength)

	var responseString strings.Builder
	content, _ := ioutil.ReadAll(response.Body)
	byteCount, _ := responseString.Write(content)
	// fmt.Println(string(content))

	fmt.Println("Byte count is = ", byteCount)
	fmt.Println("Data = ", responseString.String())

}

func PerformPostJsonRequest() {
	const myurl = "https://webhook.site/60bbe84e-2a92-494e-a967-7a69d8c0cbac"

	requestBody := strings.NewReader(`
	{
		"a" : "a"
	}
	`)

	response, err := http.Post(myurl, "application/json", requestBody)
	if err != nil {
		panic(err)
	}

	defer response.Body.Close()

	fmt.Println("Status code: ", response.StatusCode)
	fmt.Println("Content lenght code: ", response.ContentLength)

	var responseString strings.Builder
	content, _ := ioutil.ReadAll(response.Body)
	byteCount, _ := responseString.Write(content)
	// fmt.Println(string(content))

	fmt.Println("Byte count is = ", byteCount)
	fmt.Println("Data = ", responseString.String())
}

func PerformPostFormRequest() {
	const myurl = "https://webhook.site/60bbe84e-2a92-494e-a967-7a69d8c0cbac"

	data := url.Values{}
	data.Add("firstname", "ankit")

	response, err := http.PostForm(myurl, data)
	if err != nil {
		panic(err)
	}

	defer response.Body.Close()

	fmt.Println("Status code: ", response.StatusCode)
	fmt.Println("Content lenght code: ", response.ContentLength)

	var responseString strings.Builder
	content, _ := ioutil.ReadAll(response.Body)
	byteCount, _ := responseString.Write(content)
	// fmt.Println(string(content))

	fmt.Println("Byte count is = ", byteCount)
	fmt.Println("Data = ", responseString.String())
}
