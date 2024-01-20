package main

import "fmt"

func main() {
	fmt.Println("Welcome to array in golang")
	var fruitList [4]string
	fruitList[0] = "a"
	fruitList[1] = "b"
	fruitList[2] = "c"
	fmt.Println(" fruitList = ", fruitList)
	fmt.Println(" Length of fruitList = ", len(fruitList))

	var vegList = [3]string{"d", "e", "f"}
	fmt.Println(" vegList = ", vegList)
	fmt.Println(" Length of vegList = ", len(vegList))

}
