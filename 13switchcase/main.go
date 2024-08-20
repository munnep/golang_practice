package main

import (
	"fmt"
	"math/rand"
)

func main() {
	fmt.Println("example switch case")

	min := 1
    max := 6
    // fmt.Println(rand.Intn(max - min) + min)

	random_number := (rand.Intn(max - min) + min)

	fmt.Println(random_number)

	switch random_number {
	case 1:
		fmt.Println("Thrown the number:", random_number)
	case 2:
		fmt.Println("Thrown the number:", random_number)
	case 3:
		fmt.Println("Thrown the number:", random_number)
	case 4:
		fmt.Println("Thrown the number:", random_number)
	case 5:
		fmt.Println("Thrown the number:", random_number)
	case 6:
		fmt.Println("Thrown the number:", random_number)
	default:
		fmt.Println("this is unexpected")
	}

}
