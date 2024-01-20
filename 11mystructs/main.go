package main

import "fmt"

func main() {

	fmt.Println("Struct in go lang")
	// no inheritance in golang; No super or parent

	ankit := User{"Ankit", "ankit@gmail.com", true, 18}
	fmt.Println(ankit)
	fmt.Printf("Ankit detsils is : %+v\n", ankit)
	fmt.Printf("Ankit emsil is : %v\n", ankit.Email)
}

type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
}
