package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Welcome to the time study")

	presentTime := time.Now()
	fmt.Println(presentTime)

	fmt.Println(presentTime.Format("01-02-2006 15:04:04 Monday"))

	createdDate := time.Date(2020, time.April, 12, 23, 23, 0, 0, time.UTC)
	fmt.Println(createdDate)
	fmt.Println(createdDate.Format("01-02-2006 15:04:04 Monday"))

}
