package main

import "fmt"

func main() {
	fmt.Println("Welcome to function in golang")
	greeting()
	result := adder(1, 2)
	fmt.Println("Resurlt is : ", result)

	result1 := proAdder(1, 2, 3, 4)
	fmt.Println("Resurlt is : ", result1)

}

func adder(valOne int, valTwo int) int {
	return valOne + valTwo
}

func proAdder(values ...int) int {
	total := 0
	for _, v := range values {
		total += v
	}

	return total
}

func greeting() {
	fmt.Println("Namastey from golang")
}
