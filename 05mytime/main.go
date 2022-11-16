package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Welcome to time study of golang")

	presentTime := time.Now()
	fmt.Println("Present time is -", presentTime)
	fmt.Println("Present time with format is -", presentTime.Format("01-02-2006 15:04:05 Monday"))

	createdDate := time.Date(2022, time.November, 16, 02, 39, 10, 0, time.UTC)
	fmt.Println(createdDate)
	fmt.Println(createdDate.Format("01/02/2006 15:04:05 Monday"))
}
