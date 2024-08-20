package main

import "fmt"

// a constant cannot be altered during the execution of the program
const LoginToken = "sndjfksndfjknskdfj" // Public token allowed to be used.

func main() {
	fmt.Println("Working with variables")

	var username string = "patrick"
	fmt.Println("The username used is:", username)
	fmt.Printf("This variable is of type %T \n", username)

	var isLoggedIn bool = true
	fmt.Println("The variable used is:", isLoggedIn)
	fmt.Printf("This variable is of type %T \n", isLoggedIn)

	var number int = 41
	fmt.Println("The variable used is:", number)
	fmt.Printf("This variable is of type %T \n", number)


	var floating float64 = 124234.33333
	fmt.Println("The variable used is:", floating)
	fmt.Printf("This variable is of type %T \n", floating)

	// implicit use of variables
	var website = "test.com"
	fmt.Println(website)
	fmt.Printf("This variable is of type %T \n", website)

	//novar style
	justanumber := 4
	fmt.Println(justanumber)

	// print the public token
	fmt.Println(LoginToken)
}
