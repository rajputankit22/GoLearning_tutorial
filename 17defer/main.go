package main

import "fmt"

func main() {
	fmt.Println("Welcome to defer in golang")
	defer fmt.Println("a")
	defer fmt.Println("b")
	defer fmt.Println("c")
}
