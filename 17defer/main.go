package main

import (
	"fmt"
)

func main() {
	defer fmt.Println("Because of defer this is on the last line")
	defer fmt.Println("Where is this coming?? Not the last line")
	fmt.Println("==========================")
	fmt.Println("example with defer")
	fmt.Println("==========================")
    mydefer()
}

func mydefer() {
	for i := 0; i < 5; i++ {
		defer fmt.Println(i)
	}
}
