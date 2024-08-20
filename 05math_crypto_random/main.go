package main

import (
	// "crypto/rand"
	"fmt"
	// "log"
	// "math/big"

	"math/rand"
	// "crypto/rand"
)

func main() {
	fmt.Println("Working with math crypto and random")

	// following is using the math/rand package for a random number
	number := rand.Intn(6)
	fmt.Println(number+1)
	

    // using the crypto/rand package
	// randomnumber, err := rand.Int(rand.Reader,big.NewInt(6))
    // if err != nil {
	// 	log.Fatal(err)
	// }

	// fmt.Println("Number is: ", randomnumber)
	// fmt.Printf("the type is %T", randomnumber)
}
