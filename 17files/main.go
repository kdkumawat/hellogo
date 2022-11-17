package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func main() {
	fmt.Println("Welcome to files in golang")
	content := "This needs to go in a file - LearnCode"

	file, err := os.Create("./mylcogofile.txt")
	checkNilErr(err)

	length, err := io.WriteString(file, content)
	checkNilErr(err)

	fmt.Println("Length is:", length)
	defer file.Close()
	readFile("./mylcogofile.txt")
}

func readFile(fileName string) {
	databytes, err := ioutil.ReadFile(fileName)
	checkNilErr(err)

	fmt.Println("Text data in bytes \n", databytes)
	fmt.Println("Text data \n", string(databytes))
}

func checkNilErr(err error) {
	if err != nil {
		panic(err)
	}
}
