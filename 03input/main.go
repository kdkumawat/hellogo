package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	welcome := "Welcome to inputs"
	fmt.Println(welcome)

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter rating for pizza:")

	//comma ok || error ok

	rating, _ := reader.ReadString('\n')
	fmt.Println("Rating -", rating)
	fmt.Printf("Type of Rating - %T", rating)
}
