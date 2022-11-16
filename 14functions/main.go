package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("Welcome to functions in golang")
	greeter()
	// greeter2()

	result := adder(3, 5)
	fmt.Println("Result is :", result)

	proResult, msg := proAdder(2, 3, 4, 5)
	fmt.Println("Pro result is :", proResult)
	fmt.Println("Msg is :", msg)

	fmt.Println("Add your own values")
	addTwoNumber()
}

func adder(valOne int, valTwo int) int {
	return valOne + valTwo
}

func addTwoNumber() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Value 1:")
	strVal1, _ := reader.ReadString('\n')
	val1, _ := strconv.ParseFloat(strings.TrimSpace(strVal1), 32)

	fmt.Print("Value 2:")
	strVal2, _ := reader.ReadString('\n')
	val2, _ := strconv.ParseFloat(strings.TrimSpace(strVal2), 32)

	fmt.Println("Addition is:", val1+val2)
}

func proAdder(values ...int) (int, string) {
	total := 0
	for _, val := range values {
		total += val
	}

	return total, "Hi from proAdder"
}

func greeter2() {
	fmt.Println("Another method")
}

func greeter() {
	fmt.Println("Namstey from golang")
}
