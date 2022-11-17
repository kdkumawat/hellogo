package main

import "fmt"

func main() {
	fmt.Println("Welcome to structs")
	// no inheritance in golang

	kuldeep := User{"Kuldeep", "kuldeep@go.dev", true, 20}
	fmt.Println(kuldeep)
	fmt.Printf("Kuldeep details are: %+v\n", kuldeep)
	fmt.Printf("Name is: %v and Email is: %v\n", kuldeep.Name, kuldeep.Email)
	kuldeep.GetStatus()
	kuldeep.NewMail()
	fmt.Printf("Name is: %v and Email is: %v\n", kuldeep.Name, kuldeep.Email)
}

type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
}

func (u User) GetStatus() {
	fmt.Println("Is user active :", u.Status)
}

func (u User) NewMail() {
	u.Email = "test@go.dev"
	fmt.Println("Email of this user is:", u.Email)
}
