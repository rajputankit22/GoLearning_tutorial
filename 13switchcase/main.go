package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	fmt.Println("Switch and case in golang")
	rand.Seed(time.Now().UnixNano())
	deciNumber := rand.Intn(6) + 1
	fmt.Println("Value of dice is ", deciNumber)

	switch deciNumber {
	case 1:
		fmt.Println("Dice value is 1 and you can open")
	case 2:
		fmt.Println("Dice value is 2 spot")
	case 3:
		fmt.Println("Dice value is 3 spot")
		fallthrough
	case 4:
		fmt.Println("Dice value is 4 spot")
		fallthrough
	case 5:
		fmt.Println("Dice value is 5 spot")
	case 6:
		fmt.Println("Dice value is 6 spot and roll dice again")
	}
}
