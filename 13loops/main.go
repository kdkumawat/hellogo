package main

import "fmt"

func main() {
	fmt.Println("Welcome to loops in golang")

	days := []string{"Sunday", "Tuesday", "Wednesday", "Friday", "Saturday"}

	fmt.Println(days)

	for d := 0; d < len(days); d++ {
		fmt.Println(days[d])
	}

	fmt.Println("\nLoop using range ---")

	for i := range days {
		fmt.Println(days[i])
	}

	fmt.Println("\nLoop using range 2 ---")

	for index, day := range days {
		fmt.Printf("Index is %v and value is %v \n", index, day)
	}

	rougueValue := 1

	for rougueValue < 10 {
		if rougueValue == 2 {
			goto lco
		}

		if rougueValue == 3 {
			rougueValue++
			continue
		}

		if rougueValue == 8 {
			break
		}
		fmt.Println("Value is:", rougueValue)
		rougueValue++
	}

lco:
	fmt.Println("Goto lco")
}
