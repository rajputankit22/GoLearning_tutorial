package main

import "fmt"

func main() {
	fmt.Println("Welcome in if else")

	loginCount := 10
	var result string

	if loginCount < 10 {
		result = "Regular user"
	} else if loginCount > 10 {
		result = "Watch out"
	} else {
		result = "Exactly 10 login count"
	}

	fmt.Println(result)

	if 9%2 == 0 {
		fmt.Println("Number is even")
	} else {
		fmt.Println("Number is odd")
	}

	if num := 3; num < 10 {
		fmt.Println("Number is even")
	} else {
		fmt.Println("Number is odd")
	}

	// if err != nil {
	// 	return nil, err
	// }

}
