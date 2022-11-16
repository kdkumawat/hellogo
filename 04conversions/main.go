package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("Welcome to conversions")
	fmt.Println("Enter pizza rating between 1 and 5")

	reader := bufio.NewReader(os.Stdin)

	rating, _ := reader.ReadString('\n')

	fmt.Println("Thanks for rating - ", rating)

	numRating, err := strconv.ParseFloat(strings.TrimSpace(rating), 64)

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Added 1 to rating - ", numRating+1)
	}
}
