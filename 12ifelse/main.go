package main

import (
	"fmt"
	"strings"
)

func main() {

	fmt.Println("==========================")
	fmt.Println("example with if else")
	fmt.Println("==========================")
	loginCount := 2
	var result string

	if loginCount < 10 {
		result = "login count smaller then 10"
	} else if loginCount > 10 {
		result = "login bigger then 10"
	} else {
		result = "login exactly 10"
	}

	fmt.Println(strings.ToUpper(result))

	if num := 2; num < 3 {
		fmt.Println("number smaller then 3")
	} else {
		fmt.Println("number 3 or higher")
	}

	// example used a lot to verify is something went wrong
	// if err != nil {
	// 	fmt.Println(error)
	// }
}
