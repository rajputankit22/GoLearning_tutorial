package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func main() {
	fmt.Println("Welcome to file in go golang")
	content := "This is the file contents"

	file, err := os.Create("./myfile.txt")
	checkNilError(err)

	length, err := io.WriteString(file, content)
	checkNilError(err)

	fmt.Println("Length is =", length)
	defer file.Close()
	readFile("./myfile.txt")
}

func readFile(filename string) {
	databyte, err := ioutil.ReadFile(filename)
	checkNilError(err)

	fmt.Println("Text data inside the file is\n", databyte)
}

func checkNilError(err error) {
	if err != nil {
		panic(err)
		// return nil, err
	}
}
