package main

import "fmt"

func main() {
	fmt.Println("Welcome to class on pointers")

	var ptr *int
	fmt.Println("Address of pointer is", ptr)
	//fmt.Println("Value of pointer is", *ptr)

	myNumber := 23
	var ptrNumber = &myNumber
	fmt.Println("Address of ptrNumber is", ptrNumber)
	fmt.Println("Value of ptrNumber is", *ptrNumber)

	*ptrNumber = *ptrNumber * 2
	fmt.Println("New Value of ptrNumber is", *ptrNumber)
}
