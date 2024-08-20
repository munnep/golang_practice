package main

import "fmt"

func main() {
	fmt.Println("welcome to strucs examples")

	patrick := User{"patrick", "patrick@test.com", true, 41 }
	fmt.Println(patrick)
	fmt.Println("just the email", patrick.Email)
}

type User struct {
	Name     string
	Email    string
	Verified bool
	Age      int
}
