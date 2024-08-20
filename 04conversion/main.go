package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	// welcome := "welcome to our pizza place"

	fmt.Println("welcome to our pizza place")
	fmt.Println("Please rate our pizza between 1 and 5")

	reader := bufio.NewReader(os.Stdin)

	// comma or error // the first value is the good input the _ is the error

	input, _ := reader.ReadString('\n') // read until you see the \n character
	fmt.Println("thank you for rating ", input)

	numRating, err := strconv.ParseFloat(strings.TrimSpace(input), 64)
    
	if err != nil {
		fmt.Println("Something wrong converting:", numRating)
	} else {
		fmt.Println("this is a now a float:", numRating)
	}

	fmt.Println("all done")

}
