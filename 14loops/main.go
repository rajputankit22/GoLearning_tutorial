package main

import "fmt"

func main() {
	fmt.Println("Welcome to Loops in golang")
	// day := []string{"Sunday", "Tuesday", "Wednesday", "Friday", "Saturday", "Thursday"}

	// for i := 0; i < len(day); i++ {
	// 	fmt.Println(day[i])
	// }

	// for i := range day {
	// 	fmt.Println(day[i])
	// }

	// for _, v := range day {
	// 	fmt.Println("Index of ", v)
	// }

	rougueValue := 1
	for rougueValue <= 10 {

		if rougueValue == 2 {
			goto lco
		}

		if rougueValue == 5 {
			rougueValue++
			break
		}
		fmt.Println("Value is: ", rougueValue)
		rougueValue++
	}

lco:
	fmt.Println("Value is lco")
}
