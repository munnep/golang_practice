package main

import "fmt"

func main() {
	fmt.Println("welcome to methods example. Always uses a struct")

	patrick := User{"patrick", "patrick@test.com", true, 41}
	fmt.Println(patrick)
	fmt.Println("just the email", patrick.Email)
	patrick.GetStatus()
	patrick.NewMail()
	fmt.Println("real email stil is:", patrick.Email)
}

type User struct {
	Name     string
	Email    string
	Verified bool
	Age      int
}

// create a method. Function that goes with the structure
func (u User) GetStatus() {
	fmt.Println("the status of the user is:", u.Verified)
}

// update email but this updates the copy not the real value. Need pointers for this
func (u User) NewMail() {
	u.Email = "test@test.com"
	fmt.Println("email for the copy value is:", u.Email)
}
