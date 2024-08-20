package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	welcome := "welcome to our pizza place"

	fmt.Println(welcome)
	fmt.Println("Please rate our pizza")

	reader := bufio.NewReader(os.Stdin)

	// comma or error // the first value is the good input the _ is the error

	input, _ := reader.ReadString('\n') // read until you see the \n character
	fmt.Println("thank you for giving our pizza a ", input)
	fmt.Printf("our input variable is of type %T", input)
}
