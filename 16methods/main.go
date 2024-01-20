package main

import "fmt"

func main() {

	fmt.Println("Struct in go lang")
	// no inheritance in golang; No super or parent

	ankit := User{"Ankit", "ankit@gmail.com", true, 18}
	fmt.Println(ankit)
	fmt.Printf("Ankit detsils is : %+v\n", ankit)
	fmt.Printf("Ankit emsil is : %v\n", ankit.Email)
	ankit.GetStatus()
	ankit.NewEmail()
}

type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
}

func (u User) GetStatus() {
	fmt.Println("Is user is active: ", u.Status)
}

func (u User) NewEmail() {
	u.Email = "test@go.dev"
	fmt.Println("Email of this user is:", u.Email)
}
