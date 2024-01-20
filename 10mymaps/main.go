package main

import "fmt"

func main() {
	fmt.Println("Welcome in golang maps")

	languages := make(map[string]string)
	languages["js"] = "JavaScript"
	languages["rb"] = "Ruby"
	languages["py"] = "Python"

	fmt.Println("List of the all languages", languages)
	fmt.Println("Js shorts form :", languages["js"])

	delete(languages, "rb")
	fmt.Println("List of the all languages", languages)

	// loops are interesting in golang

	for _, value := range languages {
		fmt.Printf("For key v, value is %v\n", value)
	}
}
