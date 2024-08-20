package main

import "fmt"

func main() {
	mynumber1 := 2
	mynumber2 := 5

	fmt.Println("value mynumber1: ", mynumber1)
	fmt.Println("value mynumber2: ", mynumber2)

	ptrnumber1 := &mynumber1

	mynumber2 = mynumber2 + 33
	fmt.Println("value mynumber1: ", mynumber1)
	fmt.Println("value ptrnumber1: ", *ptrnumber1)
	fmt.Println("value mynumber2: ", mynumber2)

	mynumber1 = 44
	fmt.Println("value ptrnumber1 that is 44: ", *ptrnumber1)

	*ptrnumber1 = 88

	fmt.Println("value mynumber1: value 88? ", mynumber1)

}
